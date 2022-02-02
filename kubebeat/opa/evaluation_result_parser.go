package opa

import (
	"time"

	"github.com/elastic/beats/v7/kubebeat/resources"
	"github.com/elastic/beats/v7/libbeat/beat"
	libevents "github.com/elastic/beats/v7/libbeat/beat/events"
	"github.com/elastic/beats/v7/libbeat/common"

	"github.com/gofrs/uuid"
	"github.com/mitchellh/mapstructure"
)

type EvaluationResultParser struct {
	index string
}

func NewEvaluationResultParser(index string) (*EvaluationResultParser, error) {
	return &EvaluationResultParser{index: index}, nil
}

func (parser *EvaluationResultParser) ParseResult(result interface{}, cycleId uuid.UUID) ([]beat.Event, error) {
	events := make([]beat.Event, 0)
	var opaResult RuleResult
	if err := decodeResults(result, &opaResult); err != nil {
		return nil, err
	}

	timestamp := time.Now()
	for _, finding := range opaResult.Findings {
		event := beat.Event{
			Timestamp: timestamp,
			Fields: common.MapStr{
				"id":       opaResult.OpaInput.ID,
				"type":     opaResult.OpaInput.Type,
				"cycle_id": cycleId,
				"result":   finding.Result,
				"resource": opaResult.OpaInput.Resource,
				"rule":     finding.Rule,
			},
		}

		// Insert datastream as index to event struct
		event.Meta = common.MapStr{libevents.FieldMetaIndex: parser.index}
		events = append(events, event)
	}

	return events, nil
}

func decodeResults(result interface{}, output *RuleResult) error {
	var opaResultMap = result.(map[string]interface{})
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{TagName: "json", Result: output})
	if err != nil {
		return err
	}

	return decoder.Decode(opaResultMap)
}

type RuleResult struct {
	Findings []Finding               `json:"findings"`
	OpaInput resources.FetcherResult `json:"opa_input"`
}

type Finding struct {
	Result interface{} `json:"result"`
	Rule   interface{} `json:"rule"`
}
