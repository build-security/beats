// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
)

type CreateNetworkAclEntryInput struct {
	_ struct{} `type:"structure"`

	// The IPv4 network range to allow or deny, in CIDR notation (for example 172.16.0.0/24).
	// We modify the specified CIDR block to its canonical form; for example, if
	// you specify 100.68.0.18/18, we modify it to 100.68.0.0/18.
	CidrBlock *string `locationName:"cidrBlock" type:"string"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// Indicates whether this is an egress rule (rule is applied to traffic leaving
	// the subnet).
	//
	// Egress is a required field
	Egress *bool `locationName:"egress" type:"boolean" required:"true"`

	// ICMP protocol: The ICMP or ICMPv6 type and code. Required if specifying protocol
	// 1 (ICMP) or protocol 58 (ICMPv6) with an IPv6 CIDR block.
	IcmpTypeCode *IcmpTypeCode `locationName:"Icmp" type:"structure"`

	// The IPv6 network range to allow or deny, in CIDR notation (for example 2001:db8:1234:1a00::/64).
	Ipv6CidrBlock *string `locationName:"ipv6CidrBlock" type:"string"`

	// The ID of the network ACL.
	//
	// NetworkAclId is a required field
	NetworkAclId *string `locationName:"networkAclId" type:"string" required:"true"`

	// TCP or UDP protocols: The range of ports the rule applies to. Required if
	// specifying protocol 6 (TCP) or 17 (UDP).
	PortRange *PortRange `locationName:"portRange" type:"structure"`

	// The protocol number. A value of "-1" means all protocols. If you specify
	// "-1" or a protocol number other than "6" (TCP), "17" (UDP), or "1" (ICMP),
	// traffic on all ports is allowed, regardless of any ports or ICMP types or
	// codes that you specify. If you specify protocol "58" (ICMPv6) and specify
	// an IPv4 CIDR block, traffic for all ICMP types and codes allowed, regardless
	// of any that you specify. If you specify protocol "58" (ICMPv6) and specify
	// an IPv6 CIDR block, you must specify an ICMP type and code.
	//
	// Protocol is a required field
	Protocol *string `locationName:"protocol" type:"string" required:"true"`

	// Indicates whether to allow or deny the traffic that matches the rule.
	//
	// RuleAction is a required field
	RuleAction RuleAction `locationName:"ruleAction" type:"string" required:"true" enum:"true"`

	// The rule number for the entry (for example, 100). ACL entries are processed
	// in ascending order by rule number.
	//
	// Constraints: Positive integer from 1 to 32766. The range 32767 to 65535 is
	// reserved for internal use.
	//
	// RuleNumber is a required field
	RuleNumber *int64 `locationName:"ruleNumber" type:"integer" required:"true"`
}

// String returns the string representation
func (s CreateNetworkAclEntryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateNetworkAclEntryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateNetworkAclEntryInput"}

	if s.Egress == nil {
		invalidParams.Add(aws.NewErrParamRequired("Egress"))
	}

	if s.NetworkAclId == nil {
		invalidParams.Add(aws.NewErrParamRequired("NetworkAclId"))
	}

	if s.Protocol == nil {
		invalidParams.Add(aws.NewErrParamRequired("Protocol"))
	}
	if len(s.RuleAction) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("RuleAction"))
	}

	if s.RuleNumber == nil {
		invalidParams.Add(aws.NewErrParamRequired("RuleNumber"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateNetworkAclEntryOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateNetworkAclEntryOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateNetworkAclEntry = "CreateNetworkAclEntry"

// CreateNetworkAclEntryRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Creates an entry (a rule) in a network ACL with the specified rule number.
// Each network ACL has a set of numbered ingress rules and a separate set of
// numbered egress rules. When determining whether a packet should be allowed
// in or out of a subnet associated with the ACL, we process the entries in
// the ACL according to the rule numbers, in ascending order. Each network ACL
// has a set of ingress rules and a separate set of egress rules.
//
// We recommend that you leave room between the rule numbers (for example, 100,
// 110, 120, ...), and not number them one right after the other (for example,
// 101, 102, 103, ...). This makes it easier to add a rule between existing
// ones without having to renumber the rules.
//
// After you add an entry, you can't modify it; you must either replace it,
// or create an entry and delete the old one.
//
// For more information about network ACLs, see Network ACLs (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_ACLs.html)
// in the Amazon Virtual Private Cloud User Guide.
//
//    // Example sending a request using CreateNetworkAclEntryRequest.
//    req := client.CreateNetworkAclEntryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateNetworkAclEntry
func (c *Client) CreateNetworkAclEntryRequest(input *CreateNetworkAclEntryInput) CreateNetworkAclEntryRequest {
	op := &aws.Operation{
		Name:       opCreateNetworkAclEntry,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateNetworkAclEntryInput{}
	}

	req := c.newRequest(op, input, &CreateNetworkAclEntryOutput{})
	req.Handlers.Unmarshal.Remove(ec2query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return CreateNetworkAclEntryRequest{Request: req, Input: input, Copy: c.CreateNetworkAclEntryRequest}
}

// CreateNetworkAclEntryRequest is the request type for the
// CreateNetworkAclEntry API operation.
type CreateNetworkAclEntryRequest struct {
	*aws.Request
	Input *CreateNetworkAclEntryInput
	Copy  func(*CreateNetworkAclEntryInput) CreateNetworkAclEntryRequest
}

// Send marshals and sends the CreateNetworkAclEntry API request.
func (r CreateNetworkAclEntryRequest) Send(ctx context.Context) (*CreateNetworkAclEntryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateNetworkAclEntryResponse{
		CreateNetworkAclEntryOutput: r.Request.Data.(*CreateNetworkAclEntryOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateNetworkAclEntryResponse is the response type for the
// CreateNetworkAclEntry API operation.
type CreateNetworkAclEntryResponse struct {
	*CreateNetworkAclEntryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateNetworkAclEntry request.
func (r *CreateNetworkAclEntryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
