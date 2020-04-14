package ses

type SES struct {
	From   string `validate:"required"`
	Region string `validate:"required"`
	Domain string `validate:"required"`
}
