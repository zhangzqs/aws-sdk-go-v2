// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Registers a container image to your Amazon Lightsail container service. This
// action is not required if you install and use the Lightsail Control
// (lightsailctl) plugin to push container images to your Lightsail container
// service. For more information, see Pushing and managing container images on your
// Amazon Lightsail container services in the Lightsail Dev Guide.
func (c *Client) RegisterContainerImage(ctx context.Context, params *RegisterContainerImageInput, optFns ...func(*Options)) (*RegisterContainerImageOutput, error) {
	if params == nil {
		params = &RegisterContainerImageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterContainerImage", params, optFns, addOperationRegisterContainerImageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterContainerImageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterContainerImageInput struct {

	// The digest of the container image to be registered.
	//
	// This member is required.
	Digest *string

	// The label for the container image when it's registered to the container service.
	// Use a descriptive label that you can use to track the different versions of your
	// registered container images. Use the GetContainerImages action to return the
	// container images registered to a Lightsail container service. The label is the
	// portion of the following image name example:
	//
	// * :container-service-1..1
	//
	// If the
	// name of your container service is mycontainerservice, and the label that you
	// specify is mystaticwebsite, then the name of the registered container image will
	// be :mycontainerservice.mystaticwebsite.1. The number at the end of these image
	// name examples represents the version of the registered container image. If you
	// push and register another container image to the same Lightsail container
	// service, with the same label, then the version number for the new registered
	// container image will be 2. If you push and register another container image, the
	// version number will be 3, and so on.
	//
	// This member is required.
	Label *string

	// The name of the container service for which to register a container image.
	//
	// This member is required.
	ServiceName *string
}

type RegisterContainerImageOutput struct {

	// Describes a container image that is registered to an Amazon Lightsail container
	// service.
	ContainerImage *types.ContainerImage

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRegisterContainerImageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRegisterContainerImage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRegisterContainerImage{}, middleware.After)
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
	if err = addOpRegisterContainerImageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterContainerImage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRegisterContainerImage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "RegisterContainerImage",
	}
}
