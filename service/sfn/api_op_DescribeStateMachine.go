// Code generated by smithy-go-codegen DO NOT EDIT.

package sfn

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sfn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes a state machine. This operation is eventually consistent. The results
// are best effort and may not reflect very recent updates and changes.
func (c *Client) DescribeStateMachine(ctx context.Context, params *DescribeStateMachineInput, optFns ...func(*Options)) (*DescribeStateMachineOutput, error) {
	if params == nil {
		params = &DescribeStateMachineInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeStateMachine", params, optFns, addOperationDescribeStateMachineMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeStateMachineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeStateMachineInput struct {

	// The Amazon Resource Name (ARN) of the state machine to describe.
	//
	// This member is required.
	StateMachineArn *string
}

type DescribeStateMachineOutput struct {

	// The date the state machine is created.
	//
	// This member is required.
	CreationDate *time.Time

	// The Amazon States Language definition of the state machine. See Amazon States
	// Language
	// (https://docs.aws.amazon.com/step-functions/latest/dg/concepts-amazon-states-language.html).
	//
	// This member is required.
	Definition *string

	// The name of the state machine. A name must not contain:
	//
	// * white space
	//
	// *
	// brackets < > { } [ ]
	//
	// * wildcard characters ? *
	//
	// * special characters " # % \ ^
	// | ~ ` $ & , ; : /
	//
	// * control characters (U+0000-001F, U+007F-009F)
	//
	// To enable
	// logging with CloudWatch Logs, the name should only contain 0-9, A-Z, a-z, - and
	// _.
	//
	// This member is required.
	Name *string

	// The Amazon Resource Name (ARN) of the IAM role used when creating this state
	// machine. (The IAM role maintains security by granting Step Functions access to
	// AWS resources.)
	//
	// This member is required.
	RoleArn *string

	// The Amazon Resource Name (ARN) that identifies the state machine.
	//
	// This member is required.
	StateMachineArn *string

	// The type of the state machine (STANDARD or EXPRESS).
	//
	// This member is required.
	Type types.StateMachineType

	// The LoggingConfiguration data type is used to set CloudWatch Logs options.
	LoggingConfiguration *types.LoggingConfiguration

	// The current status of the state machine.
	Status types.StateMachineStatus

	// Selects whether AWS X-Ray tracing is enabled.
	TracingConfiguration *types.TracingConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeStateMachineMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeStateMachine{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeStateMachine{}, middleware.After)
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
	if err = addOpDescribeStateMachineValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStateMachine(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeStateMachine(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "states",
		OperationName: "DescribeStateMachine",
	}
}
