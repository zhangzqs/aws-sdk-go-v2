// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpoint

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a message template for messages that were sent through the email
// channel.
func (c *Client) DeleteEmailTemplate(ctx context.Context, params *DeleteEmailTemplateInput, optFns ...func(*Options)) (*DeleteEmailTemplateOutput, error) {
	if params == nil {
		params = &DeleteEmailTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteEmailTemplate", params, optFns, addOperationDeleteEmailTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteEmailTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteEmailTemplateInput struct {

	// The name of the message template. A template name must start with an
	// alphanumeric character and can contain a maximum of 128 characters. The
	// characters can be alphanumeric characters, underscores (_), or hyphens (-).
	// Template names are case sensitive.
	//
	// This member is required.
	TemplateName *string

	// The unique identifier for the version of the message template to update,
	// retrieve information about, or delete. To retrieve identifiers and other
	// information for all the versions of a template, use the Template Versions
	// resource. If specified, this value must match the identifier for an existing
	// template version. If specified for an update operation, this value must match
	// the identifier for the latest existing version of the template. This restriction
	// helps ensure that race conditions don't occur. If you don't specify a value for
	// this parameter, Amazon Pinpoint does the following:
	//
	// * For a get operation,
	// retrieves information about the active version of the template.
	//
	// * For an update
	// operation, saves the updates to (overwrites) the latest existing version of the
	// template, if the create-new-version parameter isn't used or is set to false.
	//
	// *
	// For a delete operation, deletes the template, including all versions of the
	// template.
	Version *string
}

type DeleteEmailTemplateOutput struct {

	// Provides information about an API request or response.
	//
	// This member is required.
	MessageBody *types.MessageBody

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteEmailTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteEmailTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteEmailTemplate{}, middleware.After)
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
	if err = addOpDeleteEmailTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteEmailTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteEmailTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mobiletargeting",
		OperationName: "DeleteEmailTemplate",
	}
}
