// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Permanently removes a FlexMatch matchmaking configuration. To delete, specify
// the configuration name. A matchmaking configuration cannot be deleted if it is
// being used in any active matchmaking tickets. Related operations
//
// *
// CreateMatchmakingConfiguration
//
// * DescribeMatchmakingConfigurations
//
// *
// UpdateMatchmakingConfiguration
//
// * DeleteMatchmakingConfiguration
//
// *
// CreateMatchmakingRuleSet
//
// * DescribeMatchmakingRuleSets
//
// *
// ValidateMatchmakingRuleSet
//
// * DeleteMatchmakingRuleSet
func (c *Client) DeleteMatchmakingConfiguration(ctx context.Context, params *DeleteMatchmakingConfigurationInput, optFns ...func(*Options)) (*DeleteMatchmakingConfigurationOutput, error) {
	if params == nil {
		params = &DeleteMatchmakingConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteMatchmakingConfiguration", params, optFns, addOperationDeleteMatchmakingConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteMatchmakingConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request operation.
type DeleteMatchmakingConfigurationInput struct {

	// A unique identifier for a matchmaking configuration. You can use either the
	// configuration name or ARN value.
	//
	// This member is required.
	Name *string
}

type DeleteMatchmakingConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteMatchmakingConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteMatchmakingConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteMatchmakingConfiguration{}, middleware.After)
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
	if err = addOpDeleteMatchmakingConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteMatchmakingConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteMatchmakingConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DeleteMatchmakingConfiguration",
	}
}
