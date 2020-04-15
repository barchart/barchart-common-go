package aws

import (
	"github.com/barchart/common-go/pkg/configuration/aws/dynamo"
	"github.com/barchart/common-go/pkg/configuration/aws/s3"
	"github.com/barchart/common-go/pkg/configuration/aws/secretsmanager"
	"github.com/barchart/common-go/pkg/configuration/aws/ses"
	"github.com/barchart/common-go/pkg/configuration/aws/sns"
	"github.com/barchart/common-go/pkg/configuration/aws/sqs"
)

// AWS is a type of AWS configuration
type AWS struct {
	Dynamo         *map[string]dynamo.Dynamo
	SNS            *map[string]sns.SNS
	SQS            *map[string]sqs.SQS
	SES            *map[string]ses.SES
	S3             *map[string]s3.S3
	SecretsManager *secretsmanager.SecretsManager
}
