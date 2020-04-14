package sns

type SNS struct {
	Region string `validate:"required"`
	Topic  string `validate:"required"`
	Prefix string `validate:"required"`
}
