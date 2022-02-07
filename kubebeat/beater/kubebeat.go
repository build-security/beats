package beater

import (
	"context"
	"fmt"
	"github.com/elastic/beats/v7/kubebeat/constructor"
	"time"

	"github.com/elastic/beats/v7/kubebeat/config"
	"github.com/elastic/beats/v7/kubebeat/opa"
	_ "github.com/elastic/beats/v7/kubebeat/processor" // Add kubebeat default processors.
	"github.com/elastic/beats/v7/kubebeat/resources"
	"github.com/elastic/beats/v7/kubebeat/resources/conditions"
	"github.com/elastic/beats/v7/kubebeat/resources/fetchers"
	"github.com/elastic/beats/v7/libbeat/beat"
	libevents "github.com/elastic/beats/v7/libbeat/beat/events"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/common/kubernetes"
	"github.com/elastic/beats/v7/libbeat/logp"
	"github.com/elastic/beats/v7/libbeat/processors"

	"github.com/gofrs/uuid"
)

// kubebeat configuration.
type kubebeat struct {
	ctx    context.Context
	cancel context.CancelFunc

	config       config.Config
	client       beat.Client
	data         *resources.Data
	eval         *opa.Evaluator
	resultsIndex string
	constructor  constructor.Constructor
}

const (
	cycleStatusStart = "start"
	cycleStatusEnd   = "end"
	processesDir     = "/hostfs"
)

// New creates an instance of kubebeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	ctx, cancel := context.WithCancel(context.Background())

	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	logp.Info("Config initiated.")

	fetchersRegistry, err := InitRegistry(ctx, c)
	if err != nil {
		return nil, err
	}

	data, err := resources.NewData(c.Period, fetchersRegistry)
	if err != nil {
		return nil, err
	}

	evaluator, err := opa.NewEvaluator(ctx)
	if err != nil {
		return nil, err
	}

	// namespace will be passed as param from fleet on https://github.com/elastic/security-team/issues/2383 and it's user configurable
	resultsIndex := config.Datastream("", config.ResultsDatastreamIndexPrefix)
	if err != nil {
		return nil, err
	}

	constructor := constructor.NewConstructor(evaluator.Decision, resultsIndex)

	bt := &kubebeat{
		ctx:          ctx,
		cancel:       cancel,
		config:       c,
		eval:         evaluator,
		data:         data,
		resultsIndex: resultsIndex,
		constructor:  constructor,
	}
	return bt, nil
}

// Run starts kubebeat.
func (bt *kubebeat) Run(b *beat.Beat) error {
	logp.Info("kubebeat is running! Hit CTRL-C to stop it.")
	if err := bt.data.Run(bt.ctx); err != nil {
		return err
	}

	procs, err := bt.configureProcessors(bt.config.Processors)
	if err != nil {
		return err
	}

	// Connect publisher (with beat's processors)
	if bt.client, err = b.Publisher.ConnectWith(beat.ClientConfig{
		Processing: beat.ProcessingConfig{
			Processor: procs,
		},
	}); err != nil {
		return err
	}

	output := bt.data.Output()

	for {
		select {
		case <-bt.ctx.Done():
			return nil
		case o := <-output:
			cycleId, _ := uuid.NewV4()
			// update hidden-index that the beat's cycle has started
			bt.updateCycleStatus(cycleId, cycleStatusStart)
			bt.constructor.ProcessOutput(bt.ctx, bt.client, o, constructor.CycleMetadata{CycleId: cycleId})
			// update hidden-index that the beat's cycle has ended
			bt.updateCycleStatus(cycleId, cycleStatusEnd)
		}
	}
}

func InitRegistry(ctx context.Context, c config.Config) (resources.FetchersRegistry, error) {
	registry := resources.NewFetcherRegistry()
	kubef, err := fetchers.NewKubeFetcher(c.KubeConfig, c.Period)
	if err != nil {
		return nil, err
	}

	client, err := kubernetes.GetKubernetesClient("", kubernetes.KubeClientOptions{})
	if err != nil {
		return nil, err
	}

	leaseProvider := conditions.NewLeaderLeaseProvider(ctx, client)
	condition := conditions.NewLeaseFetcherCondition(leaseProvider)

	if err = registry.Register(fetchers.KubeAPIType, kubef, condition); err != nil {
		return nil, err
	}
	if err = registry.Register(fetchers.ProcessType, fetchers.NewProcessesFetcher(processesDir)); err != nil {
		return nil, err
	}
	if err = registry.Register(fetchers.FileSystemType, fetchers.NewFileFetcher(c.Files)); err != nil {
		return nil, err
	}

	return registry, nil
}

// Stop stops kubebeat.
func (bt *kubebeat) Stop() {
	bt.data.Stop(bt.ctx, bt.cancel)
	bt.eval.Stop(bt.ctx)

	bt.client.Close()
}

// updateCycleStatus updates beat status in metadata ES index.
func (bt *kubebeat) updateCycleStatus(cycleId uuid.UUID, status string) {
	metadataIndex := config.Datastream("", config.MetadataDatastreamIndexPrefix)
	cycleEndedEvent := beat.Event{
		Timestamp: time.Now(),
		Meta:      common.MapStr{libevents.FieldMetaIndex: metadataIndex},
		Fields: common.MapStr{
			"cycle_id": cycleId,
			"status":   status,
		},
	}
	bt.client.Publish(cycleEndedEvent)
}

// configureProcessors configure processors to be used by the beat
func (bt *kubebeat) configureProcessors(processorsList processors.PluginConfig) (procs *processors.Processors, err error) {
	return processors.New(processorsList)
}
