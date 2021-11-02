// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteTrafficMirrorFilterRuleInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The ID of the Traffic Mirror rule.
	//
	// TrafficMirrorFilterRuleId is a required field
	TrafficMirrorFilterRuleId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteTrafficMirrorFilterRuleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteTrafficMirrorFilterRuleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteTrafficMirrorFilterRuleInput"}

	if s.TrafficMirrorFilterRuleId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TrafficMirrorFilterRuleId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteTrafficMirrorFilterRuleOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the deleted Traffic Mirror rule.
	TrafficMirrorFilterRuleId *string `locationName:"trafficMirrorFilterRuleId" type:"string"`
}

// String returns the string representation
func (s DeleteTrafficMirrorFilterRuleOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteTrafficMirrorFilterRule = "DeleteTrafficMirrorFilterRule"

// DeleteTrafficMirrorFilterRuleRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Deletes the specified Traffic Mirror rule.
//
//    // Example sending a request using DeleteTrafficMirrorFilterRuleRequest.
//    req := client.DeleteTrafficMirrorFilterRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteTrafficMirrorFilterRule
func (c *Client) DeleteTrafficMirrorFilterRuleRequest(input *DeleteTrafficMirrorFilterRuleInput) DeleteTrafficMirrorFilterRuleRequest {
	op := &aws.Operation{
		Name:       opDeleteTrafficMirrorFilterRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteTrafficMirrorFilterRuleInput{}
	}

	req := c.newRequest(op, input, &DeleteTrafficMirrorFilterRuleOutput{})

	return DeleteTrafficMirrorFilterRuleRequest{Request: req, Input: input, Copy: c.DeleteTrafficMirrorFilterRuleRequest}
}

// DeleteTrafficMirrorFilterRuleRequest is the request type for the
// DeleteTrafficMirrorFilterRule API operation.
type DeleteTrafficMirrorFilterRuleRequest struct {
	*aws.Request
	Input *DeleteTrafficMirrorFilterRuleInput
	Copy  func(*DeleteTrafficMirrorFilterRuleInput) DeleteTrafficMirrorFilterRuleRequest
}

// Send marshals and sends the DeleteTrafficMirrorFilterRule API request.
func (r DeleteTrafficMirrorFilterRuleRequest) Send(ctx context.Context) (*DeleteTrafficMirrorFilterRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteTrafficMirrorFilterRuleResponse{
		DeleteTrafficMirrorFilterRuleOutput: r.Request.Data.(*DeleteTrafficMirrorFilterRuleOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteTrafficMirrorFilterRuleResponse is the response type for the
// DeleteTrafficMirrorFilterRule API operation.
type DeleteTrafficMirrorFilterRuleResponse struct {
	*DeleteTrafficMirrorFilterRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteTrafficMirrorFilterRule request.
func (r *DeleteTrafficMirrorFilterRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}