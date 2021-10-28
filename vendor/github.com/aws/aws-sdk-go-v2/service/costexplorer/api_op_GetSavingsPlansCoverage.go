// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package costexplorer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetSavingsPlansCoverageInput struct {
	_ struct{} `type:"structure"`

	// Filters Savings Plans coverage data by dimensions. You can filter data for
	// Savings Plans usage with the following dimensions:
	//
	//    * LINKED_ACCOUNT
	//
	//    * REGION
	//
	//    * SERVICE
	//
	//    * INSTANCE_FAMILY
	//
	// GetSavingsPlansCoverage uses the same Expression (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_Expression.html)
	// object as the other operations, but only AND is supported among each dimension.
	// If there are multiple values for a dimension, they are OR'd together.
	//
	// Cost category is also supported.
	Filter *Expression `type:"structure"`

	// The granularity of the Amazon Web Services cost data for your Savings Plans.
	// Granularity can't be set if GroupBy is set.
	//
	// The GetSavingsPlansCoverage operation supports only DAILY and MONTHLY granularities.
	Granularity Granularity `type:"string" enum:"true"`

	// You can group the data using the attributes INSTANCE_FAMILY, REGION, or SERVICE.
	GroupBy []GroupDefinition `type:"list"`

	// The number of items to be returned in a response. The default is 20, with
	// a minimum value of 1.
	MaxResults *int64 `min:"1" type:"integer"`

	// The measurement that you want your Savings Plans coverage reported in. The
	// only valid value is SpendCoveredBySavingsPlans.
	Metrics []string `type:"list"`

	// The token to retrieve the next set of results. Amazon Web Services provides
	// the token when the response from a previous call has more results than the
	// maximum page size.
	NextToken *string `type:"string"`

	// The time period that you want the usage and costs for. The Start date must
	// be within 13 months. The End date must be after the Start date, and before
	// the current date. Future dates can't be used as an End date.
	//
	// TimePeriod is a required field
	TimePeriod *DateInterval `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetSavingsPlansCoverageInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetSavingsPlansCoverageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetSavingsPlansCoverageInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.TimePeriod == nil {
		invalidParams.Add(aws.NewErrParamRequired("TimePeriod"))
	}
	if s.Filter != nil {
		if err := s.Filter.Validate(); err != nil {
			invalidParams.AddNested("Filter", err.(aws.ErrInvalidParams))
		}
	}
	if s.TimePeriod != nil {
		if err := s.TimePeriod.Validate(); err != nil {
			invalidParams.AddNested("TimePeriod", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetSavingsPlansCoverageOutput struct {
	_ struct{} `type:"structure"`

	// The token to retrieve the next set of results. Amazon Web Services provides
	// the token when the response from a previous call has more results than the
	// maximum page size.
	NextToken *string `type:"string"`

	// The amount of spend that your Savings Plans covered.
	//
	// SavingsPlansCoverages is a required field
	SavingsPlansCoverages []SavingsPlansCoverage `type:"list" required:"true"`
}

// String returns the string representation
func (s GetSavingsPlansCoverageOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetSavingsPlansCoverage = "GetSavingsPlansCoverage"

// GetSavingsPlansCoverageRequest returns a request value for making API operation for
// AWS Cost Explorer Service.
//
// Retrieves the Savings Plans covered for your account. This enables you to
// see how much of your cost is covered by a Savings Plan. An organization’s
// master account can see the coverage of the associated member accounts. This
// supports dimensions, Cost Categories, and nested expressions. For any time
// period, you can filter data for Savings Plans usage with the following dimensions:
//
//    * LINKED_ACCOUNT
//
//    * REGION
//
//    * SERVICE
//
//    * INSTANCE_FAMILY
//
// To determine valid values for a dimension, use the GetDimensionValues operation.
//
//    // Example sending a request using GetSavingsPlansCoverageRequest.
//    req := client.GetSavingsPlansCoverageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ce-2017-10-25/GetSavingsPlansCoverage
func (c *Client) GetSavingsPlansCoverageRequest(input *GetSavingsPlansCoverageInput) GetSavingsPlansCoverageRequest {
	op := &aws.Operation{
		Name:       opGetSavingsPlansCoverage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetSavingsPlansCoverageInput{}
	}

	req := c.newRequest(op, input, &GetSavingsPlansCoverageOutput{})

	return GetSavingsPlansCoverageRequest{Request: req, Input: input, Copy: c.GetSavingsPlansCoverageRequest}
}

// GetSavingsPlansCoverageRequest is the request type for the
// GetSavingsPlansCoverage API operation.
type GetSavingsPlansCoverageRequest struct {
	*aws.Request
	Input *GetSavingsPlansCoverageInput
	Copy  func(*GetSavingsPlansCoverageInput) GetSavingsPlansCoverageRequest
}

// Send marshals and sends the GetSavingsPlansCoverage API request.
func (r GetSavingsPlansCoverageRequest) Send(ctx context.Context) (*GetSavingsPlansCoverageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSavingsPlansCoverageResponse{
		GetSavingsPlansCoverageOutput: r.Request.Data.(*GetSavingsPlansCoverageOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetSavingsPlansCoverageRequestPaginator returns a paginator for GetSavingsPlansCoverage.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetSavingsPlansCoverageRequest(input)
//   p := costexplorer.NewGetSavingsPlansCoverageRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetSavingsPlansCoveragePaginator(req GetSavingsPlansCoverageRequest) GetSavingsPlansCoveragePaginator {
	return GetSavingsPlansCoveragePaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetSavingsPlansCoverageInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// GetSavingsPlansCoveragePaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetSavingsPlansCoveragePaginator struct {
	aws.Pager
}

func (p *GetSavingsPlansCoveragePaginator) CurrentPage() *GetSavingsPlansCoverageOutput {
	return p.Pager.CurrentPage().(*GetSavingsPlansCoverageOutput)
}

// GetSavingsPlansCoverageResponse is the response type for the
// GetSavingsPlansCoverage API operation.
type GetSavingsPlansCoverageResponse struct {
	*GetSavingsPlansCoverageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSavingsPlansCoverage request.
func (r *GetSavingsPlansCoverageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
