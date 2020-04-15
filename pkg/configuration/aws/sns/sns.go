package sns

// SNS is a type of AWS SNS configuration
type SNS struct {
	Region string `validate:"required"`
	Topic  string `validate:"required"`
	Prefix string `validate:"required"`
}
