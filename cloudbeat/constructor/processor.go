package constructor

import (
	"context"
	"fmt"
	"github.com/elastic/beats/v7/cloudbeat/resources"
	"github.com/elastic/beats/v7/cloudbeat/resources/fetchers"
	"github.com/elastic/beats/v7/libbeat/beat"
	libevents "github.com/elastic/beats/v7/libbeat/beat/events"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/logp"
	"github.com/mitchellh/mapstructure"
	"time"
)

type cb func(ctx context.Context, input interface{}) (interface{}, error)

type Constructor struct {
	context       context.Context
	callback      cb
	eventMetadata common.MapStr
	events        []beat.Event
}

func NewConstructor(ctx context.Context, cb cb, index string) Constructor {
	eventMetadata := common.MapStr{libevents.FieldMetaIndex: index}
	events := make([]beat.Event, 0)

	return Constructor{
		context:       ctx,
		callback:      cb,
		eventMetadata: eventMetadata,
		events:        events,
	}
}

func (c *Constructor) ProcessAggregatedResources(resources resources.Map, metadata CycleMetadata) []beat.Event {
	c.events = make([]beat.Event, 0)
	for fetcherType, fetcherResults := range resources {
		c.processEachResource(fetcherResults, ResourceTypeMetadata{CycleMetadata: metadata, Type: fetcherType})
	}

	return c.events
}

func (c *Constructor) processEachResource(results []fetchers.FetcherResult, metadata ResourceTypeMetadata) {
	for _, result := range results {
		resMetadata := ResourceMetadata{ResourceTypeMetadata: metadata, ResourceId: result.Resource.GetID()}
		if err := c.createBeatEvents(result, resMetadata); err != nil {
			fmt.Errorf("failed to create beat events for, %v, Error: %v", metadata, err)
		}
	}
}

func (c *Constructor) createBeatEvents(resource interface{}, metadata ResourceMetadata) error {
	result, err := c.callback(c.context, resource)
	if err != nil {
		logp.Error(fmt.Errorf("error running the policy: %w", err))
		return err
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

		c.events = append(c.events, event)
	}
	return nil
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
