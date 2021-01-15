// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the Challenge-Handshake Authentication Protocol (CHAP) credentials for a
// specified iSCSI target. By default, a gateway does not have CHAP enabled;
// however, for added security, you might use it. This operation is supported in
// the volume and tape gateway types. When you update CHAP credentials, all
// existing connections on the target are closed and initiators must reconnect with
// the new credentials.
func (c *Client) UpdateChapCredentials(ctx context.Context, params *UpdateChapCredentialsInput, optFns ...func(*Options)) (*UpdateChapCredentialsOutput, error) {
	if params == nil {
		params = &UpdateChapCredentialsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateChapCredentials", params, optFns, addOperationUpdateChapCredentialsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateChapCredentialsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A JSON object containing one or more of the following fields:
//
// *
// UpdateChapCredentialsInput$InitiatorName
//
// *
// UpdateChapCredentialsInput$SecretToAuthenticateInitiator
//
// *
// UpdateChapCredentialsInput$SecretToAuthenticateTarget
//
// *
// UpdateChapCredentialsInput$TargetARN
type UpdateChapCredentialsInput struct {

	// The iSCSI initiator that connects to the target.
	//
	// This member is required.
	InitiatorName *string

	// The secret key that the initiator (for example, the Windows client) must provide
	// to participate in mutual CHAP with the target. The secret key must be between 12
	// and 16 bytes when encoded in UTF-8.
	//
	// This member is required.
	SecretToAuthenticateInitiator *string

	// The Amazon Resource Name (ARN) of the iSCSI volume target. Use the
	// DescribeStorediSCSIVolumes operation to return the TargetARN for specified
	// VolumeARN.
	//
	// This member is required.
	TargetARN *string

	// The secret key that the target must provide to participate in mutual CHAP with
	// the initiator (e.g. Windows client). Byte constraints: Minimum bytes of 12.
	// Maximum bytes of 16. The secret key must be between 12 and 16 bytes when encoded
	// in UTF-8.
	SecretToAuthenticateTarget *string
}

// A JSON object containing the following fields:
type UpdateChapCredentialsOutput struct {

	// The iSCSI initiator that connects to the target. This is the same initiator name
	// specified in the request.
	InitiatorName *string

	// The Amazon Resource Name (ARN) of the target. This is the same target specified
	// in the request.
	TargetARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateChapCredentialsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateChapCredentials{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateChapCredentials{}, middleware.After)
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
	if err = addOpUpdateChapCredentialsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateChapCredentials(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateChapCredentials(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "UpdateChapCredentials",
	}
}
