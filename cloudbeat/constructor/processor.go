package constructor

import (
	"context"
	"fmt"
	"github.com/elastic/beats/v7/kubebeat/resources"
	"github.com/elastic/beats/v7/kubebeat/resources/fetchers"
	"github.com/elastic/beats/v7/libbeat/beat"
	libevents "github.com/elastic/beats/v7/libbeat/beat/events"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/logp"
	"github.com/mitchellh/mapstructure"
	"time"
)

type cb func(ctx context.Context, input interface{}) (interface{}, error)

type Constructor struct {
	callback      cb
	eventMetadata common.MapStr
}

func NewConstructor(cb cb, index string) Constructor {
	eventMetadata := common.MapStr{libevents.FieldMetaIndex: index}

	return Constructor{callback: cb, eventMetadata: eventMetadata}
}

func (c *Constructor) ProcessOutput(ctx context.Context, client beat.Client, o resources.Map, metadata CycleMetadata) {
	for fetcherType, fetcherResults := range o {
		c.processResources(ctx, client, fetcherResults, ResourceTypeMetadata{CycleMetadata: metadata, Type: fetcherType})
	}
}

func (c *Constructor) processResources(ctx context.Context, client beat.Client, results []fetchers.FetcherResult, metadata ResourceTypeMetadata) {
	for _, result := range results {
		events, err := c.createBeatEvents(ctx, result, ResourceMetadata{ResourceTypeMetadata: metadata, ResourceId: result.GetID()})
		if err != nil {
			fmt.Errorf("failed to create beat events for, %v, Error: %v", metadata, err)
		}
		client.PublishAll(events)
	}
}

func (c *Constructor) createBeatEvents(ctx context.Context, resource interface{}, metadata ResourceMetadata) ([]beat.Event, error) {
	events := make([]beat.Event, 0)
	result, err := c.callback(ctx, resource)
	if err != nil {
		logp.Error(fmt.Errorf("error running the policy: %w", err))
		return nil, err
	}

	findings, err := ParseResult(result)
	timestamp := time.Now()
	for _, finding := range findings {
		event := beat.Event{
			Meta:      c.eventMetadata,
			Timestamp: timestamp,
			Fields: common.MapStr{
				"resource_id": metadata.ResourceId,
				"type":        metadata.Type,
				"cycle_id":    metadata.CycleId,
				"result":      finding.Result,
				"resource":    resource,
				"rule":        finding.Rule,
			},
		}

		events = append(events, event)
	}

	return events, nil
}

func ParseResult(result interface{}) ([]Finding, error) {
	var opaResult RuleResult
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{Result: &opaResult})
	if err != nil {
		return nil, err
	}

	err = decoder.Decode(result)
	return opaResult.Findings, err
}
