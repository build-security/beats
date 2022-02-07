package fetchers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
)

const EKSType = "aws-eks"

type EKSFetcher struct {
	eksProvider *EKSProvider
	clusterName string
}

type EKSFetchResult struct {
	obj interface{}
}

func NewEKSFetcher(cfg aws.Config, clusterName string) (Fetcher, error) {
	eks := NewEksProvider(cfg)

	return &EKSFetcher{
		eksProvider: eks,
		clusterName: clusterName,
	}, nil
}

func (f EKSFetcher) Fetch(ctx context.Context) ([]FetcherResult, error) {
	results := make([]FetcherResult, 0)

	result, err := f.eksProvider.DescribeCluster(ctx, f.clusterName)
	results = append(results, &EKSFetchResult{result})

	return results, err
}

func (f EKSFetcher) Stop() {
}

func (res EKSFetchResult) GetID() string {
	return ""
}
