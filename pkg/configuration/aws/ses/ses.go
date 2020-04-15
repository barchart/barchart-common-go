package ses

// SES is a type of AWS SES configuration
type SES struct {
	From   string `validate:"required"`
	Region string `validate:"required"`
	Domain string `validate:"required"`
}
