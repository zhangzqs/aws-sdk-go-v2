// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an endpoint for an Amazon EFS file system.
func (c *Client) CreateLocationEfs(ctx context.Context, params *CreateLocationEfsInput, optFns ...func(*Options)) (*CreateLocationEfsOutput, error) {
	if params == nil {
		params = &CreateLocationEfsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLocationEfs", params, optFns, addOperationCreateLocationEfsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLocationEfsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// CreateLocationEfsRequest
type CreateLocationEfsInput struct {

	// The subnet and security group that the Amazon EFS file system uses. The security
	// group that you provide needs to be able to communicate with the security group
	// on the mount target in the subnet specified. The exact relationship between
	// security group M (of the mount target) and security group S (which you provide
	// for DataSync to use at this stage) is as follows:
	//
	// * Security group M (which you
	// associate with the mount target) must allow inbound access for the Transmission
	// Control Protocol (TCP) on the NFS port (2049) from security group S. You can
	// enable inbound connections either by IP address (CIDR range) or security
	// group.
	//
	// * Security group S (provided to DataSync to access EFS) should have a
	// rule that enables outbound connections to the NFS port on one of the file
	// system’s mount targets. You can enable outbound connections either by IP address
	// (CIDR range) or security group. For information about security groups and mount
	// targets, see Security Groups for Amazon EC2 Instances and Mount Targets in the
	// Amazon EFS User Guide.
	//
	// This member is required.
	Ec2Config *types.Ec2Config

	// The Amazon Resource Name (ARN) for the Amazon EFS file system.
	//
	// This member is required.
	EfsFilesystemArn *string

	// A subdirectory in the location’s path. This subdirectory in the EFS file system
	// is used to read data from the EFS source location or write data to the EFS
	// destination. By default, AWS DataSync uses the root directory. Subdirectory must
	// be specified with forward slashes. For example, /path/to/folder.
	Subdirectory *string

	// The key-value pair that represents a tag that you want to add to the resource.
	// The value can be an empty string. This value helps you manage, filter, and
	// search for your resources. We recommend that you create a name tag for your
	// location.
	Tags []types.TagListEntry
}

// CreateLocationEfs
type CreateLocationEfsOutput struct {

	// The Amazon Resource Name (ARN) of the Amazon EFS file system location that is
	// created.
	LocationArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateLocationEfsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateLocationEfs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateLocationEfs{}, middleware.After)
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
	if err = addOpCreateLocationEfsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLocationEfs(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLocationEfs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "CreateLocationEfs",
	}
}
