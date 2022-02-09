package resources

import (
	"context"
	"fmt"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/cfgfile"
	"github.com/elastic/beats/v7/libbeat/common"
)

type Ff struct {
}

func (f *Ff) Create(p beat.PipelineConnector, config *common.Config) (cfgfile.Runner, error) {
	f1 := ParseConfigFetcher(config)
	if f1.err != nil {
		return nil, f1.err
	}

	return &FetcherRunner{
		f: f1.f,
	}, nil
}

func (f *Ff) CheckConfig(config *common.Config) error {
	return nil
}

type FetcherRunner struct {
	fmt.Stringer
	f Fetcher
}

func (r *FetcherRunner) Start() {
	r.f.Fetch(context.Background())
}

func (r *FetcherRunner) Stop() {
	r.f.Stop()
}
