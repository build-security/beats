// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeVpcEndpointConnectionNotificationsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the notification.
	ConnectionNotificationId *string `type:"string"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// One or more filters.
	//
	//    * connection-notification-arn - The ARN of the SNS topic for the notification.
	//
	//    * connection-notification-id - The ID of the notification.
	//
	//    * connection-notification-state - The state of the notification (Enabled
	//    | Disabled).
	//
	//    * connection-notification-type - The type of notification (Topic).
	//
	//    * service-id - The ID of the endpoint service.
	//
	//    * vpc-endpoint-id - The ID of the VPC endpoint.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another request with the returned NextToken value.
	MaxResults *int64 `type:"integer"`

	// The token to request the next page of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeVpcEndpointConnectionNotificationsInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeVpcEndpointConnectionNotificationsOutput struct {
	_ struct{} `type:"structure"`

	// One or more notifications.
	ConnectionNotificationSet []ConnectionNotification `locationName:"connectionNotificationSet" locationNameList:"item" type:"list"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeVpcEndpointConnectionNotificationsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeVpcEndpointConnectionNotifications = "DescribeVpcEndpointConnectionNotifications"

// DescribeVpcEndpointConnectionNotificationsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the connection notifications for VPC endpoints and VPC endpoint
// services.
//
//    // Example sending a request using DescribeVpcEndpointConnectionNotificationsRequest.
//    req := client.DescribeVpcEndpointConnectionNotificationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeVpcEndpointConnectionNotifications
func (c *Client) DescribeVpcEndpointConnectionNotificationsRequest(input *DescribeVpcEndpointConnectionNotificationsInput) DescribeVpcEndpointConnectionNotificationsRequest {
	op := &aws.Operation{
		Name:       opDescribeVpcEndpointConnectionNotifications,
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
		input = &DescribeVpcEndpointConnectionNotificationsInput{}
	}

	req := c.newRequest(op, input, &DescribeVpcEndpointConnectionNotificationsOutput{})

	return DescribeVpcEndpointConnectionNotificationsRequest{Request: req, Input: input, Copy: c.DescribeVpcEndpointConnectionNotificationsRequest}
}

// DescribeVpcEndpointConnectionNotificationsRequest is the request type for the
// DescribeVpcEndpointConnectionNotifications API operation.
type DescribeVpcEndpointConnectionNotificationsRequest struct {
	*aws.Request
	Input *DescribeVpcEndpointConnectionNotificationsInput
	Copy  func(*DescribeVpcEndpointConnectionNotificationsInput) DescribeVpcEndpointConnectionNotificationsRequest
}

// Send marshals and sends the DescribeVpcEndpointConnectionNotifications API request.
func (r DescribeVpcEndpointConnectionNotificationsRequest) Send(ctx context.Context) (*DescribeVpcEndpointConnectionNotificationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeVpcEndpointConnectionNotificationsResponse{
		DescribeVpcEndpointConnectionNotificationsOutput: r.Request.Data.(*DescribeVpcEndpointConnectionNotificationsOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeVpcEndpointConnectionNotificationsRequestPaginator returns a paginator for DescribeVpcEndpointConnectionNotifications.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeVpcEndpointConnectionNotificationsRequest(input)
//   p := ec2.NewDescribeVpcEndpointConnectionNotificationsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeVpcEndpointConnectionNotificationsPaginator(req DescribeVpcEndpointConnectionNotificationsRequest) DescribeVpcEndpointConnectionNotificationsPaginator {
	return DescribeVpcEndpointConnectionNotificationsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeVpcEndpointConnectionNotificationsInput
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

// DescribeVpcEndpointConnectionNotificationsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeVpcEndpointConnectionNotificationsPaginator struct {
	aws.Pager
}

func (p *DescribeVpcEndpointConnectionNotificationsPaginator) CurrentPage() *DescribeVpcEndpointConnectionNotificationsOutput {
	return p.Pager.CurrentPage().(*DescribeVpcEndpointConnectionNotificationsOutput)
}

// DescribeVpcEndpointConnectionNotificationsResponse is the response type for the
// DescribeVpcEndpointConnectionNotifications API operation.
type DescribeVpcEndpointConnectionNotificationsResponse struct {
	*DescribeVpcEndpointConnectionNotificationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeVpcEndpointConnectionNotifications request.
func (r *DescribeVpcEndpointConnectionNotificationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
