// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListGroupsForUserInput struct {
	_ struct{} `type:"structure"`

	// Use this parameter only when paginating results and only after you receive
	// a response indicating that the results are truncated. Set it to the value
	// of the Marker element in the response that you received to indicate where
	// the next call should start.
	Marker *string `min:"1" type:"string"`

	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true.
	//
	// If you do not include this parameter, the number of items defaults to 100.
	// Note that IAM might return fewer results, even when there are more results
	// available. In that case, the IsTruncated response element returns true, and
	// Marker contains a value to include in the subsequent call that tells the
	// service where to continue from.
	MaxItems *int64 `min:"1" type:"integer"`

	// The name of the user to list groups for.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// UserName is a required field
	UserName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ListGroupsForUserInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListGroupsForUserInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListGroupsForUserInput"}
	if s.Marker != nil && len(*s.Marker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Marker", 1))
	}
	if s.MaxItems != nil && *s.MaxItems < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxItems", 1))
	}

	if s.UserName == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserName"))
	}
	if s.UserName != nil && len(*s.UserName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a successful ListGroupsForUser request.
type ListGroupsForUserOutput struct {
	_ struct{} `type:"structure"`

	// A list of groups.
	//
	// Groups is a required field
	Groups []Group `type:"list" required:"true"`

	// A flag that indicates whether there are more items to return. If your results
	// were truncated, you can make a subsequent pagination request using the Marker
	// request parameter to retrieve more items. Note that IAM might return fewer
	// than the MaxItems number of results even when there are more results available.
	// We recommend that you check IsTruncated after every call to ensure that you
	// receive all your results.
	IsTruncated *bool `type:"boolean"`

	// When IsTruncated is true, this element is present and contains the value
	// to use for the Marker parameter in a subsequent pagination request.
	Marker *string `type:"string"`
}

// String returns the string representation
func (s ListGroupsForUserOutput) String() string {
	return awsutil.Prettify(s)
}

const opListGroupsForUser = "ListGroupsForUser"

// ListGroupsForUserRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Lists the IAM groups that the specified IAM user belongs to.
//
// You can paginate the results using the MaxItems and Marker parameters.
//
//    // Example sending a request using ListGroupsForUserRequest.
//    req := client.ListGroupsForUserRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/ListGroupsForUser
func (c *Client) ListGroupsForUserRequest(input *ListGroupsForUserInput) ListGroupsForUserRequest {
	op := &aws.Operation{
		Name:       opListGroupsForUser,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxItems",
			TruncationToken: "IsTruncated",
		},
	}

	if input == nil {
		input = &ListGroupsForUserInput{}
	}

	req := c.newRequest(op, input, &ListGroupsForUserOutput{})

	return ListGroupsForUserRequest{Request: req, Input: input, Copy: c.ListGroupsForUserRequest}
}

// ListGroupsForUserRequest is the request type for the
// ListGroupsForUser API operation.
type ListGroupsForUserRequest struct {
	*aws.Request
	Input *ListGroupsForUserInput
	Copy  func(*ListGroupsForUserInput) ListGroupsForUserRequest
}

// Send marshals and sends the ListGroupsForUser API request.
func (r ListGroupsForUserRequest) Send(ctx context.Context) (*ListGroupsForUserResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListGroupsForUserResponse{
		ListGroupsForUserOutput: r.Request.Data.(*ListGroupsForUserOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListGroupsForUserRequestPaginator returns a paginator for ListGroupsForUser.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListGroupsForUserRequest(input)
//   p := iam.NewListGroupsForUserRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListGroupsForUserPaginator(req ListGroupsForUserRequest) ListGroupsForUserPaginator {
	return ListGroupsForUserPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListGroupsForUserInput
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

// ListGroupsForUserPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListGroupsForUserPaginator struct {
	aws.Pager
}

func (p *ListGroupsForUserPaginator) CurrentPage() *ListGroupsForUserOutput {
	return p.Pager.CurrentPage().(*ListGroupsForUserOutput)
}

// ListGroupsForUserResponse is the response type for the
// ListGroupsForUser API operation.
type ListGroupsForUserResponse struct {
	*ListGroupsForUserOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListGroupsForUser request.
func (r *ListGroupsForUserResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
