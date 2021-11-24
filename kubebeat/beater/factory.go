package beater

import (
	"context"
	"fmt"
	"github.com/elastic/beats/v7/kubebeat/config"
	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/cfgfile"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/logp"
)

type factory struct {
	err  chan error
	beat *beat.Beat
}

func newFactory(b *beat.Beat, err chan error) *factory {
	return &factory{
		beat: b,
		err:  err,
	}
}

func (k factory) Create(p beat.PipelineConnector, cfg *common.Config) (cfgfile.Runner, error) {
	ctx, cancelFunc := context.WithCancel(context.Background())

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

	eventParser, err := NewEvaluationResultParser()
	if err != nil {
		return nil, err
	}

	kubef, err := NewKubeFetcher(c.KubeConfig, c.Period)
	if err != nil {
		return nil, err
	}

	// Register the fetchers to collect different types of data
	data.RegisterFetcher("kube_api", kubef)
	data.RegisterFetcher("processes", NewProcessesFetcher(procfsdir))
	data.RegisterFetcher("file_system", NewFileFetcher(c.Files))
	// Connect the publisher to later send events using the returned client
	client, err := p.Connect()
	if err != nil {
		return nil, err
	}
	r := &runner{
		done:         ctx.Done(),
		config:       c,
		eval:         evaluator,
		data:         data,
		resultParser: eventParser,
		scheduler:    scheduler,
		pipe:         p,
		err:          k.err,
		client:       client,
		cancelFunc:   cancelFunc,
	}
	return r, nil
}

func (p *factory) CheckConfig(cfg *common.Config) error {
	c := config.Config{}
	return cfg.Unpack(&c)
}
