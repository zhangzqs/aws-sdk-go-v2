// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The Amazon EMR Studio APIs are in preview release for Amazon EMR and are subject
// to change. Creates a new Amazon EMR Studio.
func (c *Client) CreateStudio(ctx context.Context, params *CreateStudioInput, optFns ...func(*Options)) (*CreateStudioOutput, error) {
	if params == nil {
		params = &CreateStudioInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateStudio", params, optFns, addOperationCreateStudioMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateStudioOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateStudioInput struct {

	// Specifies whether the Studio authenticates users using single sign-on (SSO) or
	// IAM. Amazon EMR Studio currently only supports SSO authentication.
	//
	// This member is required.
	AuthMode types.AuthMode

	// The ID of the Amazon EMR Studio Engine security group. The Engine security group
	// allows inbound network traffic from the Workspace security group, and it must be
	// in the same VPC specified by VpcId.
	//
	// This member is required.
	EngineSecurityGroupId *string

	// A descriptive name for the Amazon EMR Studio.
	//
	// This member is required.
	Name *string

	// The IAM role that will be assumed by the Amazon EMR Studio. The service role
	// provides a way for Amazon EMR Studio to interoperate with other AWS services.
	//
	// This member is required.
	ServiceRole *string

	// A list of subnet IDs to associate with the Studio. The subnets must belong to
	// the VPC specified by VpcId. Studio users can create a Workspace in any of the
	// specified subnets.
	//
	// This member is required.
	SubnetIds []string

	// The IAM user role that will be assumed by users and groups logged in to a
	// Studio. The permissions attached to this IAM role can be scoped down for each
	// user or group using session policies.
	//
	// This member is required.
	UserRole *string

	// The ID of the Amazon Virtual Private Cloud (Amazon VPC) to associate with the
	// Studio.
	//
	// This member is required.
	VpcId *string

	// The ID of the Amazon EMR Studio Workspace security group. The Workspace security
	// group allows outbound network traffic to resources in the Engine security group,
	// and it must be in the same VPC specified by VpcId.
	//
	// This member is required.
	WorkspaceSecurityGroupId *string

	// The default Amazon S3 location to back up EMR Studio Workspaces and notebook
	// files. A Studio user can select an alternative Amazon S3 location when creating
	// a Workspace.
	DefaultS3Location *string

	// A detailed description of the Studio.
	Description *string

	// A list of tags to associate with the Studio. Tags are user-defined key-value
	// pairs that consist of a required key string with a maximum of 128 characters,
	// and an optional value string with a maximum of 256 characters.
	Tags []types.Tag
}

type CreateStudioOutput struct {

	// The ID of the Amazon EMR Studio.
	StudioId *string

	// The unique Studio access URL.
	Url *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateStudioMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateStudio{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateStudio{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateStudioValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateStudio(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateStudio(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "CreateStudio",
	}
}
