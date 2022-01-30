package beater

import (
	"github.com/elastic/beats/v7/kubebeat/proc"
)

const (
	procfsdir   = "/hostfs"
	ProcessType = "process"
)

type ProcessResource struct {
	ID   string        `json:"id"`
	PID  string        `json:"pid"`
	Cmd  string        `json:"command"`
	Stat proc.ProcStat `json:"stat"`
}

type ProcessesFetcher struct {
	directory string // parent directory of target procfs
}

func NewProcessesFetcher(dir string) Fetcher {
	return &ProcessesFetcher{
		directory: dir,
	}
}

func (f *ProcessesFetcher) Fetch() ([]FetcherResult, error) {
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

		ret = append(ret, FetcherResult{ProcessType, ProcessResource{p, p, cmd, stat}})
	}

	return ret, nil
}

func (f *ProcessesFetcher) Stop() {
}
