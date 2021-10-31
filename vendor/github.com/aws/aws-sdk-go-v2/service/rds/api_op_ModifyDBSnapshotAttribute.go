// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ModifyDBSnapshotAttributeInput struct {
	_ struct{} `type:"structure"`

	// The name of the DB snapshot attribute to modify.
	//
	// To manage authorization for other AWS accounts to copy or restore a manual
	// DB snapshot, set this value to restore.
	//
	// To view the list of attributes available to modify, use the DescribeDBSnapshotAttributes
	// API action.
	//
	// AttributeName is a required field
	AttributeName *string `type:"string" required:"true"`

	// The identifier for the DB snapshot to modify the attributes for.
	//
	// DBSnapshotIdentifier is a required field
	DBSnapshotIdentifier *string `type:"string" required:"true"`

	// A list of DB snapshot attributes to add to the attribute specified by AttributeName.
	//
	// To authorize other AWS accounts to copy or restore a manual snapshot, set
	// this list to include one or more AWS account IDs, or all to make the manual
	// DB snapshot restorable by any AWS account. Do not add the all value for any
	// manual DB snapshots that contain private information that you don't want
	// available to all AWS accounts.
	ValuesToAdd []string `locationNameList:"AttributeValue" type:"list"`

	// A list of DB snapshot attributes to remove from the attribute specified by
	// AttributeName.
	//
	// To remove authorization for other AWS accounts to copy or restore a manual
	// snapshot, set this list to include one or more AWS account identifiers, or
	// all to remove authorization for any AWS account to copy or restore the DB
	// snapshot. If you specify all, an AWS account whose account ID is explicitly
	// added to the restore attribute can still copy or restore the manual DB snapshot.
	ValuesToRemove []string `locationNameList:"AttributeValue" type:"list"`
}

// String returns the string representation
func (s ModifyDBSnapshotAttributeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyDBSnapshotAttributeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyDBSnapshotAttributeInput"}

	if s.AttributeName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AttributeName"))
	}

	if s.DBSnapshotIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBSnapshotIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifyDBSnapshotAttributeOutput struct {
	_ struct{} `type:"structure"`

	// Contains the results of a successful call to the DescribeDBSnapshotAttributes
	// API action.
	//
	// Manual DB snapshot attributes are used to authorize other AWS accounts to
	// copy or restore a manual DB snapshot. For more information, see the ModifyDBSnapshotAttribute
	// API action.
	DBSnapshotAttributesResult *DBSnapshotAttributesResult `type:"structure"`
}

// String returns the string representation
func (s ModifyDBSnapshotAttributeOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyDBSnapshotAttribute = "ModifyDBSnapshotAttribute"

// ModifyDBSnapshotAttributeRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Adds an attribute and values to, or removes an attribute and values from,
// a manual DB snapshot.
//
// To share a manual DB snapshot with other AWS accounts, specify restore as
// the AttributeName and use the ValuesToAdd parameter to add a list of IDs
// of the AWS accounts that are authorized to restore the manual DB snapshot.
// Uses the value all to make the manual DB snapshot public, which means it
// can be copied or restored by all AWS accounts.
//
// Don't add the all value for any manual DB snapshots that contain private
// information that you don't want available to all AWS accounts.
//
// If the manual DB snapshot is encrypted, it can be shared, but only by specifying
// a list of authorized AWS account IDs for the ValuesToAdd parameter. You can't
// use all as a value for that parameter in this case.
//
// To view which AWS accounts have access to copy or restore a manual DB snapshot,
// or whether a manual DB snapshot public or private, use the DescribeDBSnapshotAttributes
// API action. The accounts are returned as values for the restore attribute.
//
//    // Example sending a request using ModifyDBSnapshotAttributeRequest.
//    req := client.ModifyDBSnapshotAttributeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/ModifyDBSnapshotAttribute
func (c *Client) ModifyDBSnapshotAttributeRequest(input *ModifyDBSnapshotAttributeInput) ModifyDBSnapshotAttributeRequest {
	op := &aws.Operation{
		Name:       opModifyDBSnapshotAttribute,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyDBSnapshotAttributeInput{}
	}

	req := c.newRequest(op, input, &ModifyDBSnapshotAttributeOutput{})

	return ModifyDBSnapshotAttributeRequest{Request: req, Input: input, Copy: c.ModifyDBSnapshotAttributeRequest}
}

// ModifyDBSnapshotAttributeRequest is the request type for the
// ModifyDBSnapshotAttribute API operation.
type ModifyDBSnapshotAttributeRequest struct {
	*aws.Request
	Input *ModifyDBSnapshotAttributeInput
	Copy  func(*ModifyDBSnapshotAttributeInput) ModifyDBSnapshotAttributeRequest
}

// Send marshals and sends the ModifyDBSnapshotAttribute API request.
func (r ModifyDBSnapshotAttributeRequest) Send(ctx context.Context) (*ModifyDBSnapshotAttributeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyDBSnapshotAttributeResponse{
		ModifyDBSnapshotAttributeOutput: r.Request.Data.(*ModifyDBSnapshotAttributeOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyDBSnapshotAttributeResponse is the response type for the
// ModifyDBSnapshotAttribute API operation.
type ModifyDBSnapshotAttributeResponse struct {
	*ModifyDBSnapshotAttributeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyDBSnapshotAttribute request.
func (r *ModifyDBSnapshotAttributeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
