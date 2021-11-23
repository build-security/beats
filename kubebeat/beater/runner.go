package beater

import (
	"fmt"
	"github.com/elastic/beats/v7/kubebeat/config"
	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/logp"
	"github.com/gofrs/uuid"
	"time"
)

type runner struct {
	done           chan struct{}
	config         config.Config
	client         beat.Client
	eval           *evaluator
	data           *Data
	opaEventParser *opaEventParser
	scheduler      ResourceScheduler
	pipe           beat.PipelineConnector
	err            chan error
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

func (r *runner) String() string {
	return "kubebeat.runner"
}

func (r *runner) Start() {
	logp.Info("kubebeat is running! Hit CTRL-C to stop it.")

	err := r.data.Run()
	if err != nil {
		r.err <- err
		return
	}
	defer r.data.Stop()

	if r.client, err = r.pipe.Connect(); err != nil {
		r.err <- err
		return
	}

	// ticker := time.NewTicker(r.config.Period)
	output := r.data.Output()

	for {
		select {
		case <-r.done:
			return
		case o := <-output:
			runId, _ := uuid.NewV4()
			omap := o.(map[string][]interface{})

			resourceCallback := func(resource interface{}) {
				r.resourceIteration(resource, runId)
			}

			r.scheduler.ScheduleResources(omap, resourceCallback)
		}
	}
}

func (r *runner) resourceIteration(resource interface{}, runId uuid.UUID) {
	timestamp := time.Now()

	result, err := r.eval.Decision(resource)
	if err != nil {
		logp.Error(fmt.Errorf("error running the policy: %w", err))
		return
	}

	events, err := r.opaEventParser.ParseResult(result, runId, timestamp)

	if err != nil {
		logp.Error(fmt.Errorf("error running the policy: %w", err))
		return
	}

	r.client.PublishAll(events)
}

func (r *runner) Stop() {
	r.client.Close()
	r.eval.Stop()
	close(r.done)
}

