// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeVpcEndpointConnectionsInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// One or more filters.
	//
	//    * service-id - The ID of the service.
	//
	//    * vpc-endpoint-owner - The AWS account number of the owner of the endpoint.
	//
	//    * vpc-endpoint-state - The state of the endpoint (pendingAcceptance |
	//    pending | available | deleting | deleted | rejected | failed).
	//
	//    * vpc-endpoint-id - The ID of the endpoint.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The maximum number of results to return for the request in a single page.
	// The remaining results of the initial request can be seen by sending another
	// request with the returned NextToken value. This value can be between 5 and
	// 1,000; if MaxResults is given a value larger than 1,000, only 1,000 results
	// are returned.
	MaxResults *int64 `type:"integer"`

	// The token to retrieve the next page of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeVpcEndpointConnectionsInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeVpcEndpointConnectionsOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Information about one or more VPC endpoint connections.
	VpcEndpointConnections []VpcEndpointConnection `locationName:"vpcEndpointConnectionSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeVpcEndpointConnectionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeVpcEndpointConnections = "DescribeVpcEndpointConnections"

// DescribeVpcEndpointConnectionsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the VPC endpoint connections to your VPC endpoint services, including
// any endpoints that are pending your acceptance.
//
//    // Example sending a request using DescribeVpcEndpointConnectionsRequest.
//    req := client.DescribeVpcEndpointConnectionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeVpcEndpointConnections
func (c *Client) DescribeVpcEndpointConnectionsRequest(input *DescribeVpcEndpointConnectionsInput) DescribeVpcEndpointConnectionsRequest {
	op := &aws.Operation{
		Name:       opDescribeVpcEndpointConnections,
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
		input = &DescribeVpcEndpointConnectionsInput{}
	}

	req := c.newRequest(op, input, &DescribeVpcEndpointConnectionsOutput{})

	return DescribeVpcEndpointConnectionsRequest{Request: req, Input: input, Copy: c.DescribeVpcEndpointConnectionsRequest}
}

// DescribeVpcEndpointConnectionsRequest is the request type for the
// DescribeVpcEndpointConnections API operation.
type DescribeVpcEndpointConnectionsRequest struct {
	*aws.Request
	Input *DescribeVpcEndpointConnectionsInput
	Copy  func(*DescribeVpcEndpointConnectionsInput) DescribeVpcEndpointConnectionsRequest
}

// Send marshals and sends the DescribeVpcEndpointConnections API request.
func (r DescribeVpcEndpointConnectionsRequest) Send(ctx context.Context) (*DescribeVpcEndpointConnectionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeVpcEndpointConnectionsResponse{
		DescribeVpcEndpointConnectionsOutput: r.Request.Data.(*DescribeVpcEndpointConnectionsOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeVpcEndpointConnectionsRequestPaginator returns a paginator for DescribeVpcEndpointConnections.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeVpcEndpointConnectionsRequest(input)
//   p := ec2.NewDescribeVpcEndpointConnectionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeVpcEndpointConnectionsPaginator(req DescribeVpcEndpointConnectionsRequest) DescribeVpcEndpointConnectionsPaginator {
	return DescribeVpcEndpointConnectionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeVpcEndpointConnectionsInput
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

// DescribeVpcEndpointConnectionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeVpcEndpointConnectionsPaginator struct {
	aws.Pager
}

func (p *DescribeVpcEndpointConnectionsPaginator) CurrentPage() *DescribeVpcEndpointConnectionsOutput {
	return p.Pager.CurrentPage().(*DescribeVpcEndpointConnectionsOutput)
}

// DescribeVpcEndpointConnectionsResponse is the response type for the
// DescribeVpcEndpointConnections API operation.
type DescribeVpcEndpointConnectionsResponse struct {
	*DescribeVpcEndpointConnectionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeVpcEndpointConnections request.
func (r *DescribeVpcEndpointConnectionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
