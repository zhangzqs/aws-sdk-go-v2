// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a Network File System (NFS) file share on an existing file gateway. In
// Storage Gateway, a file share is a file system mount point backed by Amazon S3
// cloud storage. Storage Gateway exposes file shares using an NFS interface. This
// operation is only supported for file gateways. File gateway requires AWS
// Security Token Service (AWS STS) to be activated to enable you to create a file
// share. Make sure AWS STS is activated in the AWS Region you are creating your
// file gateway in. If AWS STS is not activated in the AWS Region, activate it. For
// information about how to activate AWS STS, see Activating and deactivating AWS
// STS in an AWS Region
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_enable-regions.html)
// in the AWS Identity and Access Management User Guide. File gateway does not
// support creating hard or symbolic links on a file share.
func (c *Client) CreateNFSFileShare(ctx context.Context, params *CreateNFSFileShareInput, optFns ...func(*Options)) (*CreateNFSFileShareOutput, error) {
	if params == nil {
		params = &CreateNFSFileShareInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateNFSFileShare", params, optFns, addOperationCreateNFSFileShareMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateNFSFileShareOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// CreateNFSFileShareInput
type CreateNFSFileShareInput struct {

	// A unique string value that you supply that is used by file gateway to ensure
	// idempotent file share creation.
	//
	// This member is required.
	ClientToken *string

	// The Amazon Resource Name (ARN) of the file gateway on which you want to create a
	// file share.
	//
	// This member is required.
	GatewayARN *string

	// The ARN of the backend storage used for storing file data. A prefix name can be
	// added to the S3 bucket name. It must end with a "/".
	//
	// This member is required.
	LocationARN *string

	// The ARN of the AWS Identity and Access Management (IAM) role that a file gateway
	// assumes when it accesses the underlying storage.
	//
	// This member is required.
	Role *string

	// Refresh cache information.
	CacheAttributes *types.CacheAttributes

	// The list of clients that are allowed to access the file gateway. The list must
	// contain either valid IP addresses or valid CIDR blocks.
	ClientList []string

	// The default storage class for objects put into an Amazon S3 bucket by the file
	// gateway. The default value is S3_INTELLIGENT_TIERING. Optional. Valid Values:
	// S3_STANDARD | S3_INTELLIGENT_TIERING | S3_STANDARD_IA | S3_ONEZONE_IA
	DefaultStorageClass *string

	// The name of the file share. Optional. FileShareName must be set if an S3 prefix
	// name is set in LocationARN.
	FileShareName *string

	// A value that enables guessing of the MIME type for uploaded objects based on
	// file extensions. Set this value to true to enable MIME type guessing, otherwise
	// set to false. The default value is true. Valid Values: true | false
	GuessMIMETypeEnabled *bool

	// Set to true to use Amazon S3 server-side encryption with your own AWS KMS key,
	// or false to use a key managed by Amazon S3. Optional. Valid Values: true | false
	KMSEncrypted *bool

	// The Amazon Resource Name (ARN) of a symmetric customer master key (CMK) used for
	// Amazon S3 server-side encryption. Storage Gateway does not support asymmetric
	// CMKs. This value can only be set when KMSEncrypted is true. Optional.
	KMSKey *string

	// File share default values. Optional.
	NFSFileShareDefaults *types.NFSFileShareDefaults

	// The notification policy of the file share.
	NotificationPolicy *string

	// A value that sets the access control list (ACL) permission for objects in the S3
	// bucket that a file gateway puts objects into. The default value is private.
	ObjectACL types.ObjectACL

	// A value that sets the write status of a file share. Set this value to true to
	// set the write status to read-only, otherwise set to false. Valid Values: true |
	// false
	ReadOnly *bool

	// A value that sets who pays the cost of the request and the cost associated with
	// data download from the S3 bucket. If this value is set to true, the requester
	// pays the costs; otherwise, the S3 bucket owner pays. However, the S3 bucket
	// owner always pays the cost of storing data. RequesterPays is a configuration for
	// the S3 bucket that backs the file share, so make sure that the configuration on
	// the file share is the same as the S3 bucket configuration. Valid Values: true |
	// false
	RequesterPays *bool

	// A value that maps a user to anonymous user. Valid values are the following:
	//
	// *
	// RootSquash: Only root is mapped to anonymous user.
	//
	// * NoSquash: No one is mapped
	// to anonymous user.
	//
	// * AllSquash: Everyone is mapped to anonymous user.
	Squash *string

	// A list of up to 50 tags that can be assigned to the NFS file share. Each tag is
	// a key-value pair. Valid characters for key and value are letters, spaces, and
	// numbers representable in UTF-8 format, and the following special characters: + -
	// = . _ : / @. The maximum length of a tag's key is 128 characters, and the
	// maximum length for a tag's value is 256.
	Tags []types.Tag
}

// CreateNFSFileShareOutput
type CreateNFSFileShareOutput struct {

	// The Amazon Resource Name (ARN) of the newly created file share.
	FileShareARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateNFSFileShareMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateNFSFileShare{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateNFSFileShare{}, middleware.After)
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
	if err = addOpCreateNFSFileShareValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateNFSFileShare(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateNFSFileShare(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "CreateNFSFileShare",
	}
}
