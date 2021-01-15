// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts an asynchronous job to transcribe speech to text.
func (c *Client) StartTranscriptionJob(ctx context.Context, params *StartTranscriptionJobInput, optFns ...func(*Options)) (*StartTranscriptionJobOutput, error) {
	if params == nil {
		params = &StartTranscriptionJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartTranscriptionJob", params, optFns, addOperationStartTranscriptionJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartTranscriptionJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartTranscriptionJobInput struct {

	// An object that describes the input media for a transcription job.
	//
	// This member is required.
	Media *types.Media

	// The name of the job. You can't use the strings "." or ".." by themselves as the
	// job name. The name must also be unique within an AWS account. If you try to
	// create a transcription job with the same name as a previous transcription job,
	// you get a ConflictException error.
	//
	// This member is required.
	TranscriptionJobName *string

	// An object that contains the request parameters for content redaction.
	ContentRedaction *types.ContentRedaction

	// Set this field to true to enable automatic language identification. Automatic
	// language identification is disabled by default. You receive a
	// BadRequestException error if you enter a value for a LanguageCode.
	IdentifyLanguage *bool

	// Provides information about how a transcription job is executed. Use this field
	// to indicate that the job can be queued for deferred execution if the concurrency
	// limit is reached and there are no slots available to immediately run the job.
	JobExecutionSettings *types.JobExecutionSettings

	// The language code for the language used in the input media file.
	LanguageCode types.LanguageCode

	// An object containing a list of languages that might be present in your
	// collection of audio files. Automatic language identification chooses a language
	// that best matches the source audio from that list.
	LanguageOptions []types.LanguageCode

	// The format of the input media file.
	MediaFormat types.MediaFormat

	// The sample rate, in Hertz, of the audio track in the input media file. If you do
	// not specify the media sample rate, Amazon Transcribe determines the sample rate.
	// If you specify the sample rate, it must match the sample rate detected by Amazon
	// Transcribe. In most cases, you should leave the MediaSampleRateHertz field blank
	// and let Amazon Transcribe determine the sample rate.
	MediaSampleRateHertz *int32

	// Choose the custom language model you use for your transcription job in this
	// parameter.
	ModelSettings *types.ModelSettings

	// The location where the transcription is stored. If you set the OutputBucketName,
	// Amazon Transcribe puts the transcript in the specified S3 bucket. When you call
	// the GetTranscriptionJob operation, the operation returns this location in the
	// TranscriptFileUri field. If you enable content redaction, the redacted
	// transcript appears in RedactedTranscriptFileUri. If you enable content redaction
	// and choose to output an unredacted transcript, that transcript's location still
	// appears in the TranscriptFileUri. The S3 bucket must have permissions that allow
	// Amazon Transcribe to put files in the bucket. For more information, see
	// Permissions Required for IAM User Roles
	// (https://docs.aws.amazon.com/transcribe/latest/dg/security_iam_id-based-policy-examples.html#auth-role-iam-user).
	// You can specify an AWS Key Management Service (KMS) key to encrypt the output of
	// your transcription using the OutputEncryptionKMSKeyId parameter. If you don't
	// specify a KMS key, Amazon Transcribe uses the default Amazon S3 key for
	// server-side encryption of transcripts that are placed in your S3 bucket. If you
	// don't set the OutputBucketName, Amazon Transcribe generates a pre-signed URL, a
	// shareable URL that provides secure access to your transcription, and returns it
	// in the TranscriptFileUri field. Use this URL to download the transcription.
	OutputBucketName *string

	// The Amazon Resource Name (ARN) of the AWS Key Management Service (KMS) key used
	// to encrypt the output of the transcription job. The user calling the
	// StartTranscriptionJob operation must have permission to use the specified KMS
	// key. You can use either of the following to identify a KMS key in the current
	// account:
	//
	// * KMS Key ID: "1234abcd-12ab-34cd-56ef-1234567890ab"
	//
	// * KMS Key Alias:
	// "alias/ExampleAlias"
	//
	// You can use either of the following to identify a KMS key
	// in the current account or another account:
	//
	// * Amazon Resource Name (ARN) of a
	// KMS Key: "arn:aws:kms:region:account
	// ID:key/1234abcd-12ab-34cd-56ef-1234567890ab"
	//
	// * ARN of a KMS Key Alias:
	// "arn:aws:kms:region:account ID:alias/ExampleAlias"
	//
	// If you don't specify an
	// encryption key, the output of the transcription job is encrypted with the
	// default Amazon S3 key (SSE-S3). If you specify a KMS key to encrypt your output,
	// you must also specify an output location in the OutputBucketName parameter.
	OutputEncryptionKMSKeyId *string

	// You can specify a location in an Amazon S3 bucket to store the output of your
	// transcription job. If you don't specify an output key, Amazon Transcribe stores
	// the output of your transcription job in the Amazon S3 bucket you specified. By
	// default, the object key is "your-transcription-job-name.json". You can use
	// output keys to specify the Amazon S3 prefix and file name of the transcription
	// output. For example, specifying the Amazon S3 prefix, "folder1/folder2/", as an
	// output key would lead to the output being stored as
	// "folder1/folder2/your-transcription-job-name.json". If you specify
	// "my-other-job-name.json" as the output key, the object key is changed to
	// "my-other-job-name.json". You can use an output key to change both the prefix
	// and the file name, for example "folder/my-other-job-name.json". If you specify
	// an output key, you must also specify an S3 bucket in the OutputBucketName
	// parameter.
	OutputKey *string

	// A Settings object that provides optional settings for a transcription job.
	Settings *types.Settings
}

type StartTranscriptionJobOutput struct {

	// An object containing details of the asynchronous transcription job.
	TranscriptionJob *types.TranscriptionJob

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStartTranscriptionJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartTranscriptionJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartTranscriptionJob{}, middleware.After)
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
	if err = addOpStartTranscriptionJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartTranscriptionJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartTranscriptionJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "StartTranscriptionJob",
	}
}
