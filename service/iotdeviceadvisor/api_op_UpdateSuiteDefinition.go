// Code generated by smithy-go-codegen DO NOT EDIT.

package iotdeviceadvisor

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotdeviceadvisor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates a Device Advisor test suite.
func (c *Client) UpdateSuiteDefinition(ctx context.Context, params *UpdateSuiteDefinitionInput, optFns ...func(*Options)) (*UpdateSuiteDefinitionOutput, error) {
	if params == nil {
		params = &UpdateSuiteDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateSuiteDefinition", params, optFns, addOperationUpdateSuiteDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateSuiteDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateSuiteDefinitionInput struct {

	// Updates a Device Advisor test suite with suite definition id.
	//
	// This member is required.
	SuiteDefinitionId *string

	// Updates a Device Advisor test suite with suite definition configuration.
	SuiteDefinitionConfiguration *types.SuiteDefinitionConfiguration
}

type UpdateSuiteDefinitionOutput struct {

	// Updates a Device Advisor test suite with TimeStamp of when it was created.
	CreatedAt *time.Time

	// Updates a Device Advisor test suite with TimeStamp of when it was updated.
	LastUpdatedAt *time.Time

	// Updates a Device Advisor test suite with Amazon Resource name.
	SuiteDefinitionArn *string

	// Updates a Device Advisor test suite with suite UUID.
	SuiteDefinitionId *string

	// Updates a Device Advisor test suite with suite definition name.
	SuiteDefinitionName *string

	// Updates a Device Advisor test suite with suite definition version.
	SuiteDefinitionVersion *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateSuiteDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateSuiteDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateSuiteDefinition{}, middleware.After)
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
	if err = addOpUpdateSuiteDefinitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSuiteDefinition(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateSuiteDefinition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotdeviceadvisor",
		OperationName: "UpdateSuiteDefinition",
	}
}
