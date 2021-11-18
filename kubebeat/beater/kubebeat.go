package beater

import (
	"context"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"io/fs"
	"log"
	"os"
	"strings"
	"time"

	"github.com/elastic/beats/v7/kubebeat/config"
	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/logp"
)

// kubebeat configuration.
type kubebeat struct {
	done      chan struct{}
	config    config.Config
	client    beat.Client
	eval      *evaluator
	data      *Data
	scheduler ResourceScheduler
}

// New creates an instance of kubebeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	ctx := context.Background()

	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	logp.Info("Config initiated.")

	data := NewData(ctx, c.Period)
	scheduler := NewSynchronousScheduler()
	evaluator, err := NewEvaluator()
	if err != nil {
		return nil, err
	}

	kubef, err := NewKubeFetcher(c.KubeConfig, c.Period)
	if err != nil {
		return nil, err
	}

	data.RegisterFetcher("kube_api", kubef)
	data.RegisterFetcher("processes", NewProcessesFetcher(procfsdir))
	data.RegisterFetcher("file_system", NewFileFetcher(c.Files))

	bt := &kubebeat{
		done:      make(chan struct{}),
		config:    c,
		eval:      evaluator,
		data:      data,
		scheduler: scheduler,
	}
	return bt, nil
}

type PolicyResult map[string]RuleResult

type RuleResult struct {
	Findings []Finding   `json:"findings"`
	Resource interface{} `json:"resource"`
}

type Finding struct {
	Result interface{} `json:"result"`
	Rule   interface{} `json:"rule"`
}

// Run starts kubebeat.
func (bt *kubebeat) Run(b *beat.Beat) error {
	logp.Info("kubebeat is running! Hit CTRL-C to stop it.")

	err := bt.watcher.Start()
	if err != nil {
		return err
	}
	defer bt.data.Stop()

	if bt.client, err = b.Publisher.Connect(); err != nil {
		return err
	}

	output := bt.data.Output()
	//config := &mapstructure.DecoderConfig{
	//	TagName: "json",
	//}
	for {
		select {
		case <-bt.done:
			return nil
		case o := <-output:
			runId, _ := uuid.NewV4()
			events := make([]beat.Event, 0)
			timestamp := time.Now()

			result, err := bt.Decision(o)
			if err != nil {
				logp.Error(err)
			} else {
				var opaResult = result.(map[string]interface{})

				if findings, ok := opaResult["findings"].([]interface{}); ok {
					for _, findingRaw := range findings {
						if finding, ok := findingRaw.(map[string]interface{}); ok {
							event := beat.Event{
								Timestamp: timestamp,
								Fields: common.MapStr{
									"run_id":   runId,
									"result":   finding["result"],
									"resource": opaResult["resource"],
									"rule":     finding["rule"],
								},
							}
							events = append(events, event)

						}
					}
				}

			}

			bt.scheduler.RunResource(omap, func1)
		}

		bt.client.PublishAll(events)
		logp.Info("%v events sent", len(events))
	}
}

func (bt *kubebeat) Decision(input interface{}) (interface{}, error) {
	// get the named policy decision for the specified input
	allFile, canParse := input.(map[string]interface{})
	if canParse == true {
		if _, ok := allFile["file_system"]; !ok {
			return nil, nil
		}
		opaInputArray := allFile["file_system"].([]interface{})

		result, err := bt.opa.Decision(context.Background(), sdk.DecisionOptions{
			Path:  "main",
			Input: opaInputArray[0].(FileSystemResourceData),
		})
		if err != nil {
			return nil, err
		}

		return result.Result, nil
	}
	return nil, nil
}

// Stop stops kubebeat.
func (bt *kubebeat) Stop() {
	bt.client.Close()
	bt.eval.Stop()

	close(bt.done)
}
