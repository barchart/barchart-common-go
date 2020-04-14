package sqs

type SQS struct {
	Prefix string `validate:"required"`
	Region string `validate:"required"`
	Queue  string `validate:"required"`
}
