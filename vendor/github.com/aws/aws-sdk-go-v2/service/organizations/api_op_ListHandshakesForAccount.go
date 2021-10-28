// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListHandshakesForAccountInput struct {
	_ struct{} `type:"structure"`

	// Filters the handshakes that you want included in the response. The default
	// is all types. Use the ActionType element to limit the output to only a specified
	// type, such as INVITE, ENABLE_ALL_FEATURES, or APPROVE_ALL_FEATURES. Alternatively,
	// for the ENABLE_ALL_FEATURES handshake that generates a separate child handshake
	// for each member account, you can specify ParentHandshakeId to see only the
	// handshakes that were generated by that parent request.
	Filter *HandshakeFilter `type:"structure"`

	// The total number of results that you want included on each page of the response.
	// If you do not include this parameter, it defaults to a value that is specific
	// to the operation. If additional items exist beyond the maximum you specify,
	// the NextToken response element is present and has a value (is not null).
	// Include that value as the NextToken request parameter in the next call to
	// the operation to get the next part of the results. Note that Organizations
	// might return fewer results than the maximum even when there are more results
	// available. You should check NextToken after every operation to ensure that
	// you receive all of the results.
	MaxResults *int64 `min:"1" type:"integer"`

	// The parameter for receiving additional results if you receive a NextToken
	// response in a previous request. A NextToken response indicates that more
	// output is available. Set this parameter to the value of the previous call's
	// NextToken response to indicate where the output should continue from.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListHandshakesForAccountInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListHandshakesForAccountInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListHandshakesForAccountInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListHandshakesForAccountOutput struct {
	_ struct{} `type:"structure"`

	// A list of Handshake objects with details about each of the handshakes that
	// is associated with the specified account.
	Handshakes []Handshake `type:"list"`

	// If present, indicates that more output is available than is included in the
	// current response. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to get the next part of the output. You
	// should repeat this until the NextToken response element comes back as null.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListHandshakesForAccountOutput) String() string {
	return awsutil.Prettify(s)
}

const opListHandshakesForAccount = "ListHandshakesForAccount"

// ListHandshakesForAccountRequest returns a request value for making API operation for
// AWS Organizations.
//
// Lists the current handshakes that are associated with the account of the
// requesting user.
//
// Handshakes that are ACCEPTED, DECLINED, or CANCELED appear in the results
// of this API for only 30 days after changing to that state. After that, they're
// deleted and no longer accessible.
//
// Always check the NextToken response parameter for a null value when calling
// a List* operation. These operations can occasionally return an empty set
// of results even when there are more results available. The NextToken response
// parameter value is null only when there are no more results to display.
//
// This operation can be called only from the organization's master account
// or by a member account that is a delegated administrator for an AWS service.
//
//    // Example sending a request using ListHandshakesForAccountRequest.
//    req := client.ListHandshakesForAccountRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/ListHandshakesForAccount
func (c *Client) ListHandshakesForAccountRequest(input *ListHandshakesForAccountInput) ListHandshakesForAccountRequest {
	op := &aws.Operation{
		Name:       opListHandshakesForAccount,
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
		input = &ListHandshakesForAccountInput{}
	}

	req := c.newRequest(op, input, &ListHandshakesForAccountOutput{})

	return ListHandshakesForAccountRequest{Request: req, Input: input, Copy: c.ListHandshakesForAccountRequest}
}

// ListHandshakesForAccountRequest is the request type for the
// ListHandshakesForAccount API operation.
type ListHandshakesForAccountRequest struct {
	*aws.Request
	Input *ListHandshakesForAccountInput
	Copy  func(*ListHandshakesForAccountInput) ListHandshakesForAccountRequest
}

// Send marshals and sends the ListHandshakesForAccount API request.
func (r ListHandshakesForAccountRequest) Send(ctx context.Context) (*ListHandshakesForAccountResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListHandshakesForAccountResponse{
		ListHandshakesForAccountOutput: r.Request.Data.(*ListHandshakesForAccountOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListHandshakesForAccountRequestPaginator returns a paginator for ListHandshakesForAccount.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListHandshakesForAccountRequest(input)
//   p := organizations.NewListHandshakesForAccountRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListHandshakesForAccountPaginator(req ListHandshakesForAccountRequest) ListHandshakesForAccountPaginator {
	return ListHandshakesForAccountPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListHandshakesForAccountInput
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

// ListHandshakesForAccountPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListHandshakesForAccountPaginator struct {
	aws.Pager
}

func (p *ListHandshakesForAccountPaginator) CurrentPage() *ListHandshakesForAccountOutput {
	return p.Pager.CurrentPage().(*ListHandshakesForAccountOutput)
}

// ListHandshakesForAccountResponse is the response type for the
// ListHandshakesForAccount API operation.
type ListHandshakesForAccountResponse struct {
	*ListHandshakesForAccountOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListHandshakesForAccount request.
func (r *ListHandshakesForAccountResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
