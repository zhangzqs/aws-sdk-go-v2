// Code generated by smithy-go-codegen DO NOT EDIT.

package kafka

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disassociates one or more Scram Secrets from an Amazon MSK cluster.
func (c *Client) BatchDisassociateScramSecret(ctx context.Context, params *BatchDisassociateScramSecretInput, optFns ...func(*Options)) (*BatchDisassociateScramSecretOutput, error) {
	if params == nil {
		params = &BatchDisassociateScramSecretInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchDisassociateScramSecret", params, optFns, addOperationBatchDisassociateScramSecretMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchDisassociateScramSecretOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Disassociates sasl scram secrets to cluster.
type BatchDisassociateScramSecretInput struct {

	// The Amazon Resource Name (ARN) of the cluster to be updated.
	//
	// This member is required.
	ClusterArn *string

	// List of AWS Secrets Manager secret ARNs.
	//
	// This member is required.
	SecretArnList []string
}

type BatchDisassociateScramSecretOutput struct {

	// The Amazon Resource Name (ARN) of the cluster.
	ClusterArn *string

	// List of errors when disassociating secrets to cluster.
	UnprocessedScramSecrets []types.UnprocessedScramSecret

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchDisassociateScramSecretMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchDisassociateScramSecret{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchDisassociateScramSecret{}, middleware.After)
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
	if err = addOpBatchDisassociateScramSecretValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchDisassociateScramSecret(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchDisassociateScramSecret(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kafka",
		OperationName: "BatchDisassociateScramSecret",
	}
}
