// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeregisterDBProxyTargetsInput struct {
	_ struct{} `type:"structure"`

	// One or more DB cluster identifiers.
	DBClusterIdentifiers []string `type:"list"`

	// One or more DB instance identifiers.
	DBInstanceIdentifiers []string `type:"list"`

	// The identifier of the DBProxy that is associated with the DBProxyTargetGroup.
	//
	// DBProxyName is a required field
	DBProxyName *string `type:"string" required:"true"`

	// The identifier of the DBProxyTargetGroup.
	TargetGroupName *string `type:"string"`
}

// String returns the string representation
func (s DeregisterDBProxyTargetsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeregisterDBProxyTargetsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeregisterDBProxyTargetsInput"}

	if s.DBProxyName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBProxyName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeregisterDBProxyTargetsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeregisterDBProxyTargetsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeregisterDBProxyTargets = "DeregisterDBProxyTargets"

// DeregisterDBProxyTargetsRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Remove the association between one or more DBProxyTarget data structures
// and a DBProxyTargetGroup.
//
//    // Example sending a request using DeregisterDBProxyTargetsRequest.
//    req := client.DeregisterDBProxyTargetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DeregisterDBProxyTargets
func (c *Client) DeregisterDBProxyTargetsRequest(input *DeregisterDBProxyTargetsInput) DeregisterDBProxyTargetsRequest {
	op := &aws.Operation{
		Name:       opDeregisterDBProxyTargets,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeregisterDBProxyTargetsInput{}
	}

	req := c.newRequest(op, input, &DeregisterDBProxyTargetsOutput{})

	return DeregisterDBProxyTargetsRequest{Request: req, Input: input, Copy: c.DeregisterDBProxyTargetsRequest}
}

// DeregisterDBProxyTargetsRequest is the request type for the
// DeregisterDBProxyTargets API operation.
type DeregisterDBProxyTargetsRequest struct {
	*aws.Request
	Input *DeregisterDBProxyTargetsInput
	Copy  func(*DeregisterDBProxyTargetsInput) DeregisterDBProxyTargetsRequest
}

// Send marshals and sends the DeregisterDBProxyTargets API request.
func (r DeregisterDBProxyTargetsRequest) Send(ctx context.Context) (*DeregisterDBProxyTargetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeregisterDBProxyTargetsResponse{
		DeregisterDBProxyTargetsOutput: r.Request.Data.(*DeregisterDBProxyTargetsOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeregisterDBProxyTargetsResponse is the response type for the
// DeregisterDBProxyTargets API operation.
type DeregisterDBProxyTargetsResponse struct {
	*DeregisterDBProxyTargetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeregisterDBProxyTargets request.
func (r *DeregisterDBProxyTargetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
