// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeIpv6PoolsInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// One or more filters.
	//
	//    * tag:<key> - The key/value combination of a tag assigned to the resource.
	//    Use the tag key in the filter name and the tag value as the filter value.
	//    For example, to find all resources that have a tag with the key Owner
	//    and the value TeamA, specify tag:Owner for the filter name and TeamA for
	//    the filter value.
	//
	//    * tag-key - The key of a tag assigned to the resource. Use this filter
	//    to find all resources assigned a tag with a specific key, regardless of
	//    the tag value.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int64 `min:"1" type:"integer"`

	// The token for the next page of results.
	NextToken *string `type:"string"`

	// The IDs of the IPv6 address pools.
	PoolIds []string `locationName:"PoolId" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeIpv6PoolsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeIpv6PoolsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeIpv6PoolsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeIpv6PoolsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the IPv6 address pools.
	Ipv6Pools []Ipv6Pool `locationName:"ipv6PoolSet" locationNameList:"item" type:"list"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeIpv6PoolsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeIpv6Pools = "DescribeIpv6Pools"

// DescribeIpv6PoolsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes your IPv6 address pools.
//
//    // Example sending a request using DescribeIpv6PoolsRequest.
//    req := client.DescribeIpv6PoolsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeIpv6Pools
func (c *Client) DescribeIpv6PoolsRequest(input *DescribeIpv6PoolsInput) DescribeIpv6PoolsRequest {
	op := &aws.Operation{
		Name:       opDescribeIpv6Pools,
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
		input = &DescribeIpv6PoolsInput{}
	}

	req := c.newRequest(op, input, &DescribeIpv6PoolsOutput{})

	return DescribeIpv6PoolsRequest{Request: req, Input: input, Copy: c.DescribeIpv6PoolsRequest}
}

// DescribeIpv6PoolsRequest is the request type for the
// DescribeIpv6Pools API operation.
type DescribeIpv6PoolsRequest struct {
	*aws.Request
	Input *DescribeIpv6PoolsInput
	Copy  func(*DescribeIpv6PoolsInput) DescribeIpv6PoolsRequest
}

// Send marshals and sends the DescribeIpv6Pools API request.
func (r DescribeIpv6PoolsRequest) Send(ctx context.Context) (*DescribeIpv6PoolsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeIpv6PoolsResponse{
		DescribeIpv6PoolsOutput: r.Request.Data.(*DescribeIpv6PoolsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeIpv6PoolsRequestPaginator returns a paginator for DescribeIpv6Pools.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeIpv6PoolsRequest(input)
//   p := ec2.NewDescribeIpv6PoolsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeIpv6PoolsPaginator(req DescribeIpv6PoolsRequest) DescribeIpv6PoolsPaginator {
	return DescribeIpv6PoolsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeIpv6PoolsInput
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

// DescribeIpv6PoolsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeIpv6PoolsPaginator struct {
	aws.Pager
}

func (p *DescribeIpv6PoolsPaginator) CurrentPage() *DescribeIpv6PoolsOutput {
	return p.Pager.CurrentPage().(*DescribeIpv6PoolsOutput)
}

// DescribeIpv6PoolsResponse is the response type for the
// DescribeIpv6Pools API operation.
type DescribeIpv6PoolsResponse struct {
	*DescribeIpv6PoolsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeIpv6Pools request.
func (r *DescribeIpv6PoolsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
