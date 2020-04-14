package s3

type S3 struct {
	Region string `validate:"required"`
	Bucket string `validate:"required"`
}
