package fetchers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
)

const ELBType = "aws-elb"

type ELBFetcher struct {
	elbProvider *ELBProvider
	lbNames     []string
}

type ELBFetchResult struct {
	obj interface{}
}

func NewELBFetcher(cfg aws.Config, loadBalancersNames []string) (Fetcher, error) {
	elb := NewELBProvider(cfg)

	return &ELBFetcher{
		elbProvider: elb,
		lbNames:     loadBalancersNames,
	}, nil
}

func (f ELBFetcher) Fetch(ctx context.Context) ([]FetcherResult, error) {
	results := make([]FetcherResult, 0)

	result, err := f.elbProvider.DescribeLoadBalancer(ctx, f.lbNames)
	results = append(results, ELBFetchResult{result})

	return results, err
}

func (f ELBFetcher) Stop() {
}

func (res ELBFetchResult) GetID() string {
	return ""
}
