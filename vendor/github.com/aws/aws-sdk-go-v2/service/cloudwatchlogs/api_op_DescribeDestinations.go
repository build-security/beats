// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatchlogs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeDestinationsInput struct {
	_ struct{} `type:"structure"`

	// The prefix to match. If you don't specify a value, no prefix filter is applied.
	DestinationNamePrefix *string `min:"1" type:"string"`

	// The maximum number of items returned. If you don't specify a value, the default
	// is up to 50 items.
	Limit *int64 `locationName:"limit" min:"1" type:"integer"`

	// The token for the next set of items to return. (You received this token from
	// a previous call.)
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeDestinationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDestinationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDestinationsInput"}
	if s.DestinationNamePrefix != nil && len(*s.DestinationNamePrefix) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DestinationNamePrefix", 1))
	}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeDestinationsOutput struct {
	_ struct{} `type:"structure"`

	// The destinations.
	Destinations []Destination `locationName:"destinations" type:"list"`

	// The token for the next set of items to return. The token expires after 24
	// hours.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeDestinationsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeDestinations = "DescribeDestinations"

// DescribeDestinationsRequest returns a request value for making API operation for
// Amazon CloudWatch Logs.
//
// Lists all your destinations. The results are ASCII-sorted by destination
// name.
//
//    // Example sending a request using DescribeDestinationsRequest.
//    req := client.DescribeDestinationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/DescribeDestinations
func (c *Client) DescribeDestinationsRequest(input *DescribeDestinationsInput) DescribeDestinationsRequest {
	op := &aws.Operation{
		Name:       opDescribeDestinations,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeDestinationsInput{}
	}

	req := c.newRequest(op, input, &DescribeDestinationsOutput{})

	return DescribeDestinationsRequest{Request: req, Input: input, Copy: c.DescribeDestinationsRequest}
}

// DescribeDestinationsRequest is the request type for the
// DescribeDestinations API operation.
type DescribeDestinationsRequest struct {
	*aws.Request
	Input *DescribeDestinationsInput
	Copy  func(*DescribeDestinationsInput) DescribeDestinationsRequest
}

// Send marshals and sends the DescribeDestinations API request.
func (r DescribeDestinationsRequest) Send(ctx context.Context) (*DescribeDestinationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDestinationsResponse{
		DescribeDestinationsOutput: r.Request.Data.(*DescribeDestinationsOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeDestinationsRequestPaginator returns a paginator for DescribeDestinations.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeDestinationsRequest(input)
//   p := cloudwatchlogs.NewDescribeDestinationsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeDestinationsPaginator(req DescribeDestinationsRequest) DescribeDestinationsPaginator {
	return DescribeDestinationsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeDestinationsInput
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

// DescribeDestinationsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeDestinationsPaginator struct {
	aws.Pager
}

func (p *DescribeDestinationsPaginator) CurrentPage() *DescribeDestinationsOutput {
	return p.Pager.CurrentPage().(*DescribeDestinationsOutput)
}

// DescribeDestinationsResponse is the response type for the
// DescribeDestinations API operation.
type DescribeDestinationsResponse struct {
	*DescribeDestinationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDestinations request.
func (r *DescribeDestinationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
