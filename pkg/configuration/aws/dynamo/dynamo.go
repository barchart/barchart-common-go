package dynamo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Dynamo is a type of DynamoDB configuration
type Dynamo struct {
	Prefix string `validate:"required"`
	Region string `validate:"required"`
}

// New creates a new instance of AWS DynamoDB
func (d Dynamo) New() *dynamodb.DynamoDB {
	mySession := session.Must(session.NewSession())
	dynamo := dynamodb.New(mySession, aws.NewConfig().WithRegion(d.Region))

	return dynamo
}
