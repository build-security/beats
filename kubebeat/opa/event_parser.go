package opa

import (
	"time"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/gofrs/uuid"
)

type EventParser struct {
}

func NewEventParser() (*EventParser, error) {
	return &EventParser{}, nil
}

func (parser *EventParser) ParseResult(result interface{}, uuid uuid.UUID, timestamp time.Time) ([]beat.Event, error) {

	events := make([]beat.Event, 0)
	var opaResult = result.(map[string]interface{})

	if findings, ok := opaResult["findings"].([]interface{}); ok {
		for _, findingRaw := range findings {
			if finding, ok := findingRaw.(map[string]interface{}); ok {
				beatEvent := generateBeatEvent(uuid, opaResult, finding, timestamp)
				events = append(events, beatEvent)
			}
		}
	}

	return events, nil
}

func generateBeatEvent(uuid uuid.UUID, opaResult map[string]interface{}, finding map[string]interface{},
	timestamp time.Time) beat.Event {
	return beat.Event{
		Timestamp: timestamp,
		Fields: common.MapStr{
			"run_id":   uuid,
			"result":   finding["result"],
			"resource": opaResult["resource"],
			"id":       opaResult["id"],
			"type":     opaResult["type"],
			"rule":     finding["rule"],
		},
	}
}
