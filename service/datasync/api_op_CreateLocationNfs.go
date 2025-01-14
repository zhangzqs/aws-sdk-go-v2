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

// Creates an endpoint for an Network File System (NFS) file server that DataSync
// can use for a data transfer.
func (c *Client) CreateLocationNfs(ctx context.Context, params *CreateLocationNfsInput, optFns ...func(*Options)) (*CreateLocationNfsOutput, error) {
	if params == nil {
		params = &CreateLocationNfsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLocationNfs", params, optFns, c.addOperationCreateLocationNfsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLocationNfsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// CreateLocationNfsRequest
type CreateLocationNfsInput struct {

	// Specifies the Amazon Resource Names (ARNs) of agents that DataSync uses to
	// connect to your NFS file server. If you are copying data to or from your
	// Snowcone device, see NFS Server on Snowcone (https://docs.aws.amazon.com/datasync/latest/userguide/create-nfs-location.html#nfs-on-snowcone)
	// for more information.
	//
	// This member is required.
	OnPremConfig *types.OnPremConfig

	// Specifies the IP address or domain name of your NFS file server. An agent that
	// is installed on-premises uses this hostname to mount the NFS server in a
	// network. If you are copying data to or from your Snowcone device, see NFS
	// Server on Snowcone (https://docs.aws.amazon.com/datasync/latest/userguide/create-nfs-location.html#nfs-on-snowcone)
	// for more information. You must specify be an IP version 4 address or Domain Name
	// System (DNS)-compliant name.
	//
	// This member is required.
	ServerHostname *string

	// Specifies the subdirectory in the NFS file server that DataSync transfers to or
	// from. The NFS path should be a path that's exported by the NFS server, or a
	// subdirectory of that path. The path should be such that it can be mounted by
	// other NFS clients in your network. To see all the paths exported by your NFS
	// server, run " showmount -e nfs-server-name " from an NFS client that has access
	// to your server. You can specify any directory that appears in the results, and
	// any subdirectory of that directory. Ensure that the NFS export is accessible
	// without Kerberos authentication. To transfer all the data in the folder you
	// specified, DataSync needs to have permissions to read all the data. To ensure
	// this, either configure the NFS export with no_root_squash, or ensure that the
	// permissions for all of the files that you want DataSync allow read access for
	// all users. Doing either enables the agent to read the files. For the agent to
	// access directories, you must additionally enable all execute access. If you are
	// copying data to or from your Snowcone device, see NFS Server on Snowcone (https://docs.aws.amazon.com/datasync/latest/userguide/create-nfs-location.html#nfs-on-snowcone)
	// for more information.
	//
	// This member is required.
	Subdirectory *string

	// Specifies the mount options that DataSync can use to mount your NFS share.
	MountOptions *types.NfsMountOptions

	// Specifies labels that help you categorize, filter, and search for your Amazon
	// Web Services resources. We recommend creating at least a name tag for your
	// location.
	Tags []types.TagListEntry

	noSmithyDocumentSerde
}

// CreateLocationNfsResponse
type CreateLocationNfsOutput struct {

	// The ARN of the transfer location that you created for your NFS file server.
	LocationArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLocationNfsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateLocationNfs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateLocationNfs{}, middleware.After)
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
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateLocationNfsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLocationNfs(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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

func newServiceMetadataMiddleware_opCreateLocationNfs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "CreateLocationNfs",
	}
}
