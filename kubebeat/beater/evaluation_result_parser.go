package beater

import (
	"encoding/json"
	"strings"
	"time"

	libevents "github.com/elastic/beats/v7/libbeat/beat/events"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"

	"github.com/gofrs/uuid"
	"github.com/mitchellh/mapstructure"
)

type evaluationResultParser struct {
}

func NewEvaluationResultParser() (*evaluationResultParser, error) {

	return &evaluationResultParser{}, nil
}

func (parser *evaluationResultParser) ParseResult(index, result interface{}, uuid uuid.UUID, timestamp time.Time) ([]beat.Event, error) {

	events := make([]beat.Event, 0)
	var opaResultMap = result.(map[string]interface{})
	var opaResult RuleResult
	err := mapstructure.Decode(opaResultMap, &opaResult)

	if err != nil {
		return nil, err
	}

	for _, finding := range opaResult.Findings {
		event := beat.Event{
			Timestamp: timestamp,
			Fields: common.MapStr{
				"run_id":   uuid,
				"result":   finding.Result,
				"resource": opaResult.Resource,
				"rule":     finding.Rule,
			},
		}

		// This event can't be indexed by Elasticsearch until
		// fields containing only a dot are removed or modified.
		j, err := json.Marshal(event)
		if err != nil {
			return nil, err
		}
		js := string(j)
		js = strings.ReplaceAll(js, "\".\"", "\"DOT\"")

		var e beat.Event
		json.Unmarshal([]byte(js), &e)

		// Insert datastream as index to event struct
		if index != "" {

			event.Meta = common.MapStr{libevents.FieldMetaIndex: index}
			e.Meta = common.MapStr{libevents.FieldMetaIndex: index}
		}

		events = append(events, e)
	}

	return events, err
}

type RuleResult struct {
	Findings []Finding   `json:"findings"`
	Resource interface{} `json:"resource"`
}

type Finding struct {
	Result interface{} `json:"result"`
	Rule   interface{} `json:"rule"`
}
