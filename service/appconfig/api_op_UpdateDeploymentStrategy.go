// Code generated by smithy-go-codegen DO NOT EDIT.

package appconfig

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appconfig/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a deployment strategy.
func (c *Client) UpdateDeploymentStrategy(ctx context.Context, params *UpdateDeploymentStrategyInput, optFns ...func(*Options)) (*UpdateDeploymentStrategyOutput, error) {
	if params == nil {
		params = &UpdateDeploymentStrategyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDeploymentStrategy", params, optFns, addOperationUpdateDeploymentStrategyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDeploymentStrategyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDeploymentStrategyInput struct {

	// The deployment strategy ID.
	//
	// This member is required.
	DeploymentStrategyId *string

	// Total amount of time for a deployment to last.
	DeploymentDurationInMinutes int32

	// A description of the deployment strategy.
	Description *string

	// The amount of time AppConfig monitors for alarms before considering the
	// deployment to be complete and no longer eligible for automatic roll back.
	FinalBakeTimeInMinutes int32

	// The percentage of targets to receive a deployed configuration during each
	// interval.
	GrowthFactor float32

	// The algorithm used to define how percentage grows over time. AWS AppConfig
	// supports the following growth types: Linear: For this type, AppConfig processes
	// the deployment by increments of the growth factor evenly distributed over the
	// deployment time. For example, a linear deployment that uses a growth factor of
	// 20 initially makes the configuration available to 20 percent of the targets.
	// After 1/5th of the deployment time has passed, the system updates the percentage
	// to 40 percent. This continues until 100% of the targets are set to receive the
	// deployed configuration. Exponential: For this type, AppConfig processes the
	// deployment exponentially using the following formula: G*(2^N). In this formula,
	// G is the growth factor specified by the user and N is the number of steps until
	// the configuration is deployed to all targets. For example, if you specify a
	// growth factor of 2, then the system rolls out the configuration as follows:
	// 2*(2^0)
	//     2*(2^1)
	//
	// 2*(2^2) Expressed numerically, the deployment rolls out as
	// follows: 2% of the targets, 4% of the targets, 8% of the targets, and continues
	// until the configuration has been deployed to all targets.
	GrowthType types.GrowthType
}

type UpdateDeploymentStrategyOutput struct {

	// Total amount of time the deployment lasted.
	DeploymentDurationInMinutes int32

	// The description of the deployment strategy.
	Description *string

	// The amount of time AppConfig monitored for alarms before considering the
	// deployment to be complete and no longer eligible for automatic roll back.
	FinalBakeTimeInMinutes int32

	// The percentage of targets that received a deployed configuration during each
	// interval.
	GrowthFactor float32

	// The algorithm used to define how percentage grew over time.
	GrowthType types.GrowthType

	// The deployment strategy ID.
	Id *string

	// The name of the deployment strategy.
	Name *string

	// Save the deployment strategy to a Systems Manager (SSM) document.
	ReplicateTo types.ReplicateTo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateDeploymentStrategyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDeploymentStrategy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDeploymentStrategy{}, middleware.After)
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
	if err = addOpUpdateDeploymentStrategyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDeploymentStrategy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDeploymentStrategy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appconfig",
		OperationName: "UpdateDeploymentStrategy",
	}
}
