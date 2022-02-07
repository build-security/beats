package fetchers

import (
	"context"
	"github.com/elastic/beats/v7/x-pack/osquerybeat/ext/osquery-extension/pkg/proc"
)

const (
	ProcessType = "process"
)

type ProcessesFetcher struct {
	directory string // parent directory of target procfs
}

func NewProcessesFetcher(dir string) Fetcher {
	return &ProcessesFetcher{
		directory: dir,
	}
}

func (f *ProcessesFetcher) Fetch(ctx context.Context) ([]FetcherResult, error) {
	pids, err := proc.List(f.directory)
	if err != nil {
		return nil, err
	}

	ret := make([]FetcherResult, 0)

	// If errors occur during read, then return what we have till now
	// without reporting errors.
	for _, p := range pids {
		cmd, err := proc.ReadCmdLine(f.directory, p)
		if err != nil {
			return ret, nil
		}

		stat, err := proc.ReadStat(f.directory, p)
		if err != nil {
			return ret, nil
		}

		resourceObj := ProcessResource{PID: p, Cmd: cmd, Stat: stat}
		ret = append(ret, resourceObj)
	}

	return ret, nil
}

func (f *ProcessesFetcher) Stop() {
}

func (res ProcessResource) GetID() string {
	return res.PID
}
