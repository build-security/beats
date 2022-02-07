package fetchers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
)

const ECRType = "aws-ecr"

type ECRFetcher struct {
	ecrProvider *ECRProvider
}

type ECRFetchResult struct {
	obj interface{}
}

func NewECRFetcher(cfg aws.Config) (Fetcher, error) {
	ecr := NewEcrProvider(cfg)

	return &ECRFetcher{
		ecrProvider: ecr,
	}, nil
}

func (f ECRFetcher) Fetch(ctx context.Context) ([]FetcherResult, error) {
	results := make([]FetcherResult, 0)

	// TODO - The provider should get a list of the repositories it needs to check, and not check the entire ECR account`
	repositories, err := f.ecrProvider.DescribeAllECRRepositories(ctx)
	results = append(results, &ECRFetchResult{repositories})

	return results, err
}

func (f ECRFetcher) Stop() {
}

func (res ECRFetchResult) GetID() string {
	return ""
}
