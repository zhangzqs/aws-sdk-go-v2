// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Import single or multi-volume disk images or EBS snapshots into an Amazon
// Machine Image (AMI). For more information, see Importing a VM as an Image Using
// VM Import/Export
// (https://docs.aws.amazon.com/vm-import/latest/userguide/vmimport-image-import.html)
// in the VM Import/Export User Guide.
func (c *Client) ImportImage(ctx context.Context, params *ImportImageInput, optFns ...func(*Options)) (*ImportImageOutput, error) {
	if params == nil {
		params = &ImportImageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ImportImage", params, optFns, addOperationImportImageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ImportImageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ImportImageInput struct {

	// The architecture of the virtual machine. Valid values: i386 | x86_64 | arm64
	Architecture *string

	// The client-specific data.
	ClientData *types.ClientData

	// The token to enable idempotency for VM import requests.
	ClientToken *string

	// A description string for the import image task.
	Description *string

	// Information about the disk containers.
	DiskContainers []types.ImageDiskContainer

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool

	// Specifies whether the destination AMI of the imported image should be encrypted.
	// The default CMK for EBS is used unless you specify a non-default AWS Key
	// Management Service (AWS KMS) CMK using KmsKeyId. For more information, see
	// Amazon EBS Encryption
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html) in the
	// Amazon Elastic Compute Cloud User Guide.
	Encrypted bool

	// The target hypervisor platform. Valid values: xen
	Hypervisor *string

	// An identifier for the symmetric AWS Key Management Service (AWS KMS) customer
	// master key (CMK) to use when creating the encrypted AMI. This parameter is only
	// required if you want to use a non-default CMK; if this parameter is not
	// specified, the default CMK for EBS is used. If a KmsKeyId is specified, the
	// Encrypted flag must also be set. The CMK identifier may be provided in any of
	// the following formats:
	//
	// * Key ID
	//
	// * Key alias. The alias ARN contains the
	// arn:aws:kms namespace, followed by the Region of the CMK, the AWS account ID of
	// the CMK owner, the alias namespace, and then the CMK alias. For example,
	// arn:aws:kms:us-east-1:012345678910:alias/ExampleAlias.
	//
	// * ARN using key ID. The
	// ID ARN contains the arn:aws:kms namespace, followed by the Region of the CMK,
	// the AWS account ID of the CMK owner, the key namespace, and then the CMK ID. For
	// example,
	// arn:aws:kms:us-east-1:012345678910:key/abcd1234-a123-456a-a12b-a123b4cd56ef.
	//
	// *
	// ARN using key alias. The alias ARN contains the arn:aws:kms namespace, followed
	// by the Region of the CMK, the AWS account ID of the CMK owner, the alias
	// namespace, and then the CMK alias. For example,
	// arn:aws:kms:us-east-1:012345678910:alias/ExampleAlias.
	//
	// AWS parses KmsKeyId
	// asynchronously, meaning that the action you call may appear to complete even
	// though you provided an invalid identifier. This action will eventually report
	// failure. The specified CMK must exist in the Region that the AMI is being copied
	// to. Amazon EBS does not support asymmetric CMKs.
	KmsKeyId *string

	// The ARNs of the license configurations.
	LicenseSpecifications []types.ImportImageLicenseConfigurationRequest

	// The license type to be used for the Amazon Machine Image (AMI) after importing.
	// By default, we detect the source-system operating system (OS) and apply the
	// appropriate license. Specify AWS to replace the source-system license with an
	// AWS license, if appropriate. Specify BYOL to retain the source-system license,
	// if appropriate. To use BYOL, you must have existing licenses with rights to use
	// these licenses in a third party cloud, such as AWS. For more information, see
	// Prerequisites
	// (https://docs.aws.amazon.com/vm-import/latest/userguide/vmimport-image-import.html#prerequisites-image)
	// in the VM Import/Export User Guide.
	LicenseType *string

	// The operating system of the virtual machine. Valid values: Windows | Linux
	Platform *string

	// The name of the role to use when not using the default role, 'vmimport'.
	RoleName *string

	// The tags to apply to the image being imported.
	TagSpecifications []types.TagSpecification
}

type ImportImageOutput struct {

	// The architecture of the virtual machine.
	Architecture *string

	// A description of the import task.
	Description *string

	// Indicates whether the AMI is encrypted.
	Encrypted bool

	// The target hypervisor of the import task.
	Hypervisor *string

	// The ID of the Amazon Machine Image (AMI) created by the import task.
	ImageId *string

	// The task ID of the import image task.
	ImportTaskId *string

	// The identifier for the symmetric AWS Key Management Service (AWS KMS) customer
	// master key (CMK) that was used to create the encrypted AMI.
	KmsKeyId *string

	// The ARNs of the license configurations.
	LicenseSpecifications []types.ImportImageLicenseConfigurationResponse

	// The license type of the virtual machine.
	LicenseType *string

	// The operating system of the virtual machine.
	Platform *string

	// The progress of the task.
	Progress *string

	// Information about the snapshots.
	SnapshotDetails []types.SnapshotDetail

	// A brief status of the task.
	Status *string

	// A detailed status message of the import task.
	StatusMessage *string

	// Any tags assigned to the image being imported.
	Tags []types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationImportImageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpImportImage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpImportImage{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opImportImage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opImportImage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ImportImage",
	}
}
