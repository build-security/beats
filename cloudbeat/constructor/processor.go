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

type cb func(ctx context.Context, input fetchers.FetcherResult) (interface{}, error)

type Constructor struct {
	context       context.Context
	callback      cb
	eventMetadata common.MapStr
	eventsCh      chan beat.Event
}

func NewConstructor(ctx context.Context, cb cb, index string, ch chan beat.Event) Constructor {
	eventMetadata := common.MapStr{libevents.FieldMetaIndex: index}

	return Constructor{
		context:       ctx,
		callback:      cb,
		eventMetadata: eventMetadata,
		eventsCh:      ch,
	}
}

func (c *Constructor) ProcessAggregatedResources(resources resources.Map, metadata CycleMetadata) {
	for fetcherType, fetcherResults := range resources {
		c.processEachResource(fetcherResults, ResourceTypeMetadata{CycleMetadata: metadata, Type: fetcherType})
	}
	close(c.eventsCh)
}

func (c *Constructor) processEachResource(results []fetchers.FetcherResult, metadata ResourceTypeMetadata) {
	for _, result := range results {
		resMetadata := ResourceMetadata{ResourceTypeMetadata: metadata, ResourceId: result.Resource.GetID()}
		if err := c.createBeatEvents(result, resMetadata); err != nil {
			fmt.Errorf("failed to create beat events for, %v, Error: %v", metadata, err)
		}
	}
}

func (c *Constructor) createBeatEvents(result fetchers.FetcherResult, metadata ResourceMetadata) error {
	opaResult, err := c.callback(c.context, result)
	if err != nil {
		logp.Error(fmt.Errorf("error running the policy: %w", err))
		return err
	}

	findings, err := ParseResult(opaResult)
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
				"resource":    result.Resource,
				"rule":        finding.Rule,
			},
		}

		c.eventsCh <- event
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
