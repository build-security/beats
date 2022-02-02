package resources

import "github.com/elastic/beats/v7/x-pack/osquerybeat/ext/osquery-extension/pkg/proc"

// Fetcher represents a data fetcher.
type Fetcher interface {
	Fetch() ([]FetcherResult, error)
	Stop()
	GetResourceID(resource interface{}) string
}

type FetcherResult struct {
	ID       string      `json:"id"`
	Type     string      `json:"type"`
	Resource interface{} `json:"resource"`
}

type FileSystemResource struct {
	FileName string `json:"filename"`
	FileMode string `json:"mode"`
	Gid      string `json:"gid"`
	Uid      string `json:"uid"`
	Path     string `json:"path"`
	Inode    string `json:"inode"`
}

type ProcessResource struct {
	PID  string        `json:"pid"`
	Cmd  string        `json:"command"`
	Stat proc.ProcStat `json:"stat"`
}
