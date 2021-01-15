// Code generated by smithy-go-codegen DO NOT EDIT.

package batch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/batch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an AWS Batch compute environment.
func (c *Client) UpdateComputeEnvironment(ctx context.Context, params *UpdateComputeEnvironmentInput, optFns ...func(*Options)) (*UpdateComputeEnvironmentOutput, error) {
	if params == nil {
		params = &UpdateComputeEnvironmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateComputeEnvironment", params, optFns, addOperationUpdateComputeEnvironmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateComputeEnvironmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for UpdateComputeEnvironment.
type UpdateComputeEnvironmentInput struct {

	// The name or full Amazon Resource Name (ARN) of the compute environment to
	// update.
	//
	// This member is required.
	ComputeEnvironment *string

	// Details of the compute resources managed by the compute environment. Required
	// for a managed compute environment. For more information, see Compute
	// Environments
	// (https://docs.aws.amazon.com/batch/latest/userguide/compute_environments.html)
	// in the AWS Batch User Guide.
	ComputeResources *types.ComputeResourceUpdate

	// The full Amazon Resource Name (ARN) of the IAM role that allows AWS Batch to
	// make calls to other AWS services on your behalf. For more information, see AWS
	// Batch service IAM role
	// (https://docs.aws.amazon.com/batch/latest/userguide/service_IAM_role.html) in
	// the AWS Batch User Guide. If your specified role has a path other than /, then
	// you must either specify the full role ARN (this is recommended) or prefix the
	// role name with the path. Depending on how you created your AWS Batch service
	// role, its ARN might contain the service-role path prefix. When you only specify
	// the name of the service role, AWS Batch assumes that your ARN does not use the
	// service-role path prefix. Because of this, we recommend that you specify the
	// full ARN of your service role when you create compute environments.
	ServiceRole *string

	// The state of the compute environment. Compute environments in the ENABLED state
	// can accept jobs from a queue and scale in or out automatically based on the
	// workload demand of its associated queues. If the state is ENABLED, then the AWS
	// Batch scheduler can attempt to place jobs from an associated job queue on the
	// compute resources within the environment. If the compute environment is managed,
	// then it can scale its instances out or in automatically, based on the job queue
	// demand. If the state is DISABLED, then the AWS Batch scheduler doesn't attempt
	// to place jobs within the environment. Jobs in a STARTING or RUNNING state
	// continue to progress normally. Managed compute environments in the DISABLED
	// state don't scale out. However, they scale in to minvCpus value after instances
	// become idle.
	State types.CEState
}

type UpdateComputeEnvironmentOutput struct {

	// The Amazon Resource Name (ARN) of the compute environment.
	ComputeEnvironmentArn *string

	// The name of the compute environment. Up to 128 letters (uppercase and
	// lowercase), numbers, hyphens, and underscores are allowed.
	ComputeEnvironmentName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateComputeEnvironmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateComputeEnvironment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateComputeEnvironment{}, middleware.After)
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
	if err = addOpUpdateComputeEnvironmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateComputeEnvironment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateComputeEnvironment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "batch",
		OperationName: "UpdateComputeEnvironment",
	}
}
