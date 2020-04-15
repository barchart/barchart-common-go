package s3

// S3 is a type of S3 configuration
type S3 struct {
	Region string `validate:"required"`
	Bucket string `validate:"required"`
}
