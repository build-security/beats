// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeHandshakeInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier (ID) of the handshake that you want information about.
	// You can get the ID from the original call to InviteAccountToOrganization,
	// or from a call to ListHandshakesForAccount or ListHandshakesForOrganization.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) for handshake ID string
	// requires "h-" followed by from 8 to 32 lowercase letters or digits.
	//
	// HandshakeId is a required field
	HandshakeId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeHandshakeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeHandshakeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeHandshakeInput"}

	if s.HandshakeId == nil {
		invalidParams.Add(aws.NewErrParamRequired("HandshakeId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeHandshakeOutput struct {
	_ struct{} `type:"structure"`

	// A structure that contains information about the specified handshake.
	Handshake *Handshake `type:"structure"`
}

// String returns the string representation
func (s DescribeHandshakeOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeHandshake = "DescribeHandshake"

// DescribeHandshakeRequest returns a request value for making API operation for
// AWS Organizations.
//
// Retrieves information about a previously requested handshake. The handshake
// ID comes from the response to the original InviteAccountToOrganization operation
// that generated the handshake.
//
// You can access handshakes that are ACCEPTED, DECLINED, or CANCELED for only
// 30 days after they change to that state. They're then deleted and no longer
// accessible.
//
// This operation can be called from any account in the organization.
//
//    // Example sending a request using DescribeHandshakeRequest.
//    req := client.DescribeHandshakeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/DescribeHandshake
func (c *Client) DescribeHandshakeRequest(input *DescribeHandshakeInput) DescribeHandshakeRequest {
	op := &aws.Operation{
		Name:       opDescribeHandshake,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeHandshakeInput{}
	}

	req := c.newRequest(op, input, &DescribeHandshakeOutput{})

	return DescribeHandshakeRequest{Request: req, Input: input, Copy: c.DescribeHandshakeRequest}
}

// DescribeHandshakeRequest is the request type for the
// DescribeHandshake API operation.
type DescribeHandshakeRequest struct {
	*aws.Request
	Input *DescribeHandshakeInput
	Copy  func(*DescribeHandshakeInput) DescribeHandshakeRequest
}

// Send marshals and sends the DescribeHandshake API request.
func (r DescribeHandshakeRequest) Send(ctx context.Context) (*DescribeHandshakeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeHandshakeResponse{
		DescribeHandshakeOutput: r.Request.Data.(*DescribeHandshakeOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeHandshakeResponse is the response type for the
// DescribeHandshake API operation.
type DescribeHandshakeResponse struct {
	*DescribeHandshakeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeHandshake request.
func (r *DescribeHandshakeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
