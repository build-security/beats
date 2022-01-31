package fetchers

import (
	"github.com/elastic/beats/v7/kubebeat/resources"
	"github.com/elastic/beats/v7/x-pack/osquerybeat/ext/osquery-extension/pkg/proc"
)

const (
	ProcessType = "process"
)

type ProcessResource struct {
	PID  string        `json:"pid"`
	Cmd  string        `json:"command"`
	Stat proc.ProcStat `json:"stat"`
}

type ProcessesFetcher struct {
	directory string // parent directory of target procfs
}

func NewProcessesFetcher(dir string) resources.Fetcher {
	return &ProcessesFetcher{
		directory: dir,
	}
}

func (f *ProcessesFetcher) Fetch() ([]resources.FetcherResult, error) {
	pids, err := proc.List(f.directory)
	if err != nil {
		return nil, err
	}

	ret := make([]resources.FetcherResult, 0)

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

		resourceObj := resources.ResourceInfo{ID: p, Data: ProcessResource{p, cmd, stat}}
		ret = append(ret, resources.FetcherResult{Type: ProcessType, Resource: resourceObj})

	}

	return ret, nil
}

func (f *ProcessesFetcher) Stop() {
}
