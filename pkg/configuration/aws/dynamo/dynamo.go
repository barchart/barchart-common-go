package dynamo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)
import "github.com/aws/aws-sdk-go/aws/session"

type Dynamo struct {
	Prefix string `validate:"required"`
	Region string `validate:"required"`
}

func (d Dynamo) NewDynamo() *dynamodb.DynamoDB {
	mySession := session.Must(session.NewSession())
	dynamo := dynamodb.New(mySession, aws.NewConfig().WithRegion(d.Region))

	return dynamo
}
