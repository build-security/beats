package beater

// Fetcher represents a data fetcher.
type Fetcher interface {
	Fetch() ([]FetcherResult, error)
	Stop()
}

type FetcherResult struct {
	Type     string       `json:"type"`
	Resource ResourceInfo `json:"resource"`
}

type ResourceInfo struct {
	ID   string      `json:"id"`
	Data interface{} `json:"data"`
}
