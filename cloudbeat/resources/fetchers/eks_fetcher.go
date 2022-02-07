package fetchers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/elastic/beats/v7/cloudbeat/resources"
)

const EKSType = "aws-eks"

type EKSFetcher struct {
	cfg         EKSFetcherConfig
	eksProvider *EKSProvider
}

type EKSFetcherConfig struct {
	resources.BaseFetcherConfig
	ClusterName string `config:"clusterName"`
}

func NewEKSFetcher(awsCfg aws.Config, cfg EKSFetcherConfig) (Fetcher, error) {
	eks := NewEksProvider(awsCfg)

	return &EKSFetcher{
		cfg:         cfg,
		eksProvider: eks,
	}, nil
}

func (f EKSFetcher) Fetch(ctx context.Context) ([]FetcherResult, error) {
	results := make([]FetcherResult, 0)

	result, err := f.eksProvider.DescribeCluster(ctx, f.cfg.ClusterName)
	results = append(results, FetcherResult{
		Type:     EKSType,
		Resource: result,
	})

	return results, err
}

func (f EKSFetcher) Stop() {
}

func (res EKSFetchResult) GetID() string {
	return ""
}
