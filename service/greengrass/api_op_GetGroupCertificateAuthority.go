// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retreives the CA associated with a group. Returns the public key of the CA.
func (c *Client) GetGroupCertificateAuthority(ctx context.Context, params *GetGroupCertificateAuthorityInput, optFns ...func(*Options)) (*GetGroupCertificateAuthorityOutput, error) {
	if params == nil {
		params = &GetGroupCertificateAuthorityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetGroupCertificateAuthority", params, optFns, addOperationGetGroupCertificateAuthorityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetGroupCertificateAuthorityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetGroupCertificateAuthorityInput struct {

	// The ID of the certificate authority.
	//
	// This member is required.
	CertificateAuthorityId *string

	// The ID of the Greengrass group.
	//
	// This member is required.
	GroupId *string
}

type GetGroupCertificateAuthorityOutput struct {

	// The ARN of the certificate authority for the group.
	GroupCertificateAuthorityArn *string

	// The ID of the certificate authority for the group.
	GroupCertificateAuthorityId *string

	// The PEM encoded certificate for the group.
	PemEncodedCertificate *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetGroupCertificateAuthorityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetGroupCertificateAuthority{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetGroupCertificateAuthority{}, middleware.After)
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
	if err = addOpGetGroupCertificateAuthorityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetGroupCertificateAuthority(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetGroupCertificateAuthority(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "GetGroupCertificateAuthority",
	}
}
