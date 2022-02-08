package fetchers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eks"
)

const EKSType = "aws-eks"

type EKSFetcher struct {
	cfg         EKSFetcherConfig
	eksProvider *EKSProvider
}

type EKSFetcherConfig struct {
	BaseFetcherConfig
	ClusterName string `config:"clusterName"`
}

type EKSResource struct {
	*eks.DescribeClusterResponse
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
		Resource: EKSResource{result},
	})

	return results, err
}

func (f EKSFetcher) Stop() {
}

func (res EKSResource) GetID() string {
	return ""
}
