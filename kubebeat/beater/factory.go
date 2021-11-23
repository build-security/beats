package beater

import (
	"context"
	"fmt"
	"github.com/elastic/beats/v7/kubebeat/config"
	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/cfgfile"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/logp"
	"github.com/elastic/beats/v7/libbeat/publisher/pipeline"
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

	eventParser, err := NewOpaEventParser()
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
	client, err := p.Connect()
	if err != nil {
		return nil, err
	}
	r := &runner{
		done:           make(chan struct{}),
		config:         c,
		eval:           evaluator,
		data:           data,
		opaEventParser: eventParser,
		scheduler:      scheduler,
		pipe:           p,
		err:            k.err,
		client:         client,
	}
	return r, nil
}

func (p *factory) CheckConfig(config *common.Config) error {
	runner, err := p.Create(pipeline.NewNilPipeline(), config)
	if err != nil {
		return err
	}
	runner.Stop()
	return nil
}
