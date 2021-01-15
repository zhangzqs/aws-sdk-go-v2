// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Tests the functionality of repository triggers by sending information to the
// trigger target. If real data is available in the repository, the test sends data
// from the last commit. If no data is available, sample data is generated.
func (c *Client) TestRepositoryTriggers(ctx context.Context, params *TestRepositoryTriggersInput, optFns ...func(*Options)) (*TestRepositoryTriggersOutput, error) {
	if params == nil {
		params = &TestRepositoryTriggersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TestRepositoryTriggers", params, optFns, addOperationTestRepositoryTriggersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TestRepositoryTriggersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a test repository triggers operation.
type TestRepositoryTriggersInput struct {

	// The name of the repository in which to test the triggers.
	//
	// This member is required.
	RepositoryName *string

	// The list of triggers to test.
	//
	// This member is required.
	Triggers []types.RepositoryTrigger
}

// Represents the output of a test repository triggers operation.
type TestRepositoryTriggersOutput struct {

	// The list of triggers that were not tested. This list provides the names of the
	// triggers that could not be tested, separated by commas.
	FailedExecutions []types.RepositoryTriggerExecutionFailure

	// The list of triggers that were successfully tested. This list provides the names
	// of the triggers that were successfully tested, separated by commas.
	SuccessfulExecutions []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationTestRepositoryTriggersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpTestRepositoryTriggers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpTestRepositoryTriggers{}, middleware.After)
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
	if err = addOpTestRepositoryTriggersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opTestRepositoryTriggers(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opTestRepositoryTriggers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "TestRepositoryTriggers",
	}
}
