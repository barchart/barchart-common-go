package sqs

// SQS is a type of SQS configuration
type SQS struct {
	Prefix string `validate:"required"`
	Region string `validate:"required"`
	Queue  string `validate:"required"`
}
