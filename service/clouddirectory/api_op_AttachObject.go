// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Attaches an existing object to another object. An object can be accessed in two
// ways:
//
// * Using the path
//
// * Using ObjectIdentifier
func (c *Client) AttachObject(ctx context.Context, params *AttachObjectInput, optFns ...func(*Options)) (*AttachObjectOutput, error) {
	if params == nil {
		params = &AttachObjectInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AttachObject", params, optFns, addOperationAttachObjectMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AttachObjectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AttachObjectInput struct {

	// The child object reference to be attached to the object.
	//
	// This member is required.
	ChildReference *types.ObjectReference

	// Amazon Resource Name (ARN) that is associated with the Directory where both
	// objects reside. For more information, see arns.
	//
	// This member is required.
	DirectoryArn *string

	// The link name with which the child object is attached to the parent.
	//
	// This member is required.
	LinkName *string

	// The parent object reference.
	//
	// This member is required.
	ParentReference *types.ObjectReference
}

type AttachObjectOutput struct {

	// The attached ObjectIdentifier, which is the child ObjectIdentifier.
	AttachedObjectIdentifier *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAttachObjectMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAttachObject{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAttachObject{}, middleware.After)
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
	if err = addOpAttachObjectValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAttachObject(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAttachObject(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "AttachObject",
	}
}
