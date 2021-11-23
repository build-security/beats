package beater

import (
	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/cfgfile"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/common/reload"
	"github.com/elastic/beats/v7/libbeat/logp"
	"github.com/elastic/beats/v7/libbeat/management"
)

// kubebeat configuration.
type kubebeat struct {
	config  *common.Config
	factory *factory
	done    chan struct{}

}

// New creates an instance of kubebeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	factory := newFactory(b, make(chan error))
	if err := factory.CheckConfig(cfg); err != nil {
		return nil, err
	}

	return &kubebeat{
		config:  cfg,
		factory: factory,
		done:    make(chan struct{}),
	}, nil
}


func (bt *kubebeat) Run(b *beat.Beat) error {
	if !b.Manager.Enabled() {
		return bt.runStatic(b, bt.factory)
	}
	return bt.runManaged(b, bt.factory)
}

func (bt *kubebeat) runStatic(b *beat.Beat, factory *factory) error {
	runner, err := factory.Create(b.Publisher, bt.config)
	if err != nil {
		return err
	}
	runner.Start()
	defer runner.Stop()

	logp.Debug("main", "Waiting for the runner to finish")

	select {
	case <-bt.done:
	case err := <-factory.err:
		close(bt.done)
		return err
	}
	return nil
}

func (bt *kubebeat) runManaged(b *beat.Beat, factory *factory) error {
	runner := cfgfile.NewRunnerList(management.DebugK, factory, b.Publisher)
	reload.Register.MustRegisterList("inputs", runner)
	defer runner.Stop()

	logp.Debug("main", "Waiting for the runner to finish")

	for {
		select {
		case <-bt.done:
			return nil
		case err := <-factory.err:
			// when we're managed we don't want
			// to stop if the sniffer(s) exited without an error
			// this would happen during a configuration reload
			if err != nil {
				close(bt.done)
				return err
			}
		}
	}
}

func (bt *kubebeat) Stop() {
	logp.Info("Kubebeat send stop signal")
	close(bt.done)
}
