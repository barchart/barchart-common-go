package secretsmanager

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/barchart/common-go/pkg/logger"
)

var log = logger.Log

func init() {
	log.SetReportCaller(true)
}

// SecretsManager is a type of AWS Secrets Manager configuration and provider
type SecretsManager struct {
	Region string `validate:"required"`
	sm     *secretsmanager.SecretsManager
}

// isStringJSON returns true/false if the provided string is JSON
func isStringJSON(str string) bool {
	var jsonStr map[string]interface{}
	err := json.Unmarshal([]byte(str), &jsonStr)
	return err == nil
}

// New creates new AWS Secrets Manager instance
func New(region string) *SecretsManager {
	secretsManager := SecretsManager{}
	secretsManager.Region = region

	sess, err := session.NewSession()

	if err != nil {
		log.Println("err " + err.Error())
	}

	secretsManager.sm = secretsmanager.New(sess, aws.NewConfig().WithRegion(region))

	return &secretsManager
}

// GetValue returns value from AWS Secrets Manager.
// Returns 3 variables: value, isJSON, err
func (secretsManager SecretsManager) GetValue(secretName string) (string, bool, error) {
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	}

	secretResult, err := secretsManager.sm.GetSecretValue(input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeDecryptionFailure:
				return "", false, errors.New(fmt.Sprintln(secretsmanager.ErrCodeDecryptionFailure, aerr.Error()))

			case secretsmanager.ErrCodeInternalServiceError:
				return "", false, errors.New(fmt.Sprintln(secretsmanager.ErrCodeInternalServiceError, aerr.Error()))

			case secretsmanager.ErrCodeInvalidParameterException:
				return "", false, errors.New(fmt.Sprintln(secretsmanager.ErrCodeInvalidParameterException, aerr.Error()))

			case secretsmanager.ErrCodeInvalidRequestException:
				return "", false, errors.New(fmt.Sprintln(secretsmanager.ErrCodeInvalidRequestException, aerr.Error()))

			case secretsmanager.ErrCodeResourceNotFoundException:
				return "", false, errors.New(fmt.Sprintln(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error()))
			}
		} else {
			log.Println(err.Error())
		}

		return "", false, err
	}

	var result string
	isJSON := false

	if secretResult.SecretString != nil {
		result = *secretResult.SecretString
	} else {
		decodedBinarySecretBytes := make([]byte, base64.StdEncoding.DecodedLen(len(secretResult.SecretBinary)))
		length, err := base64.StdEncoding.Decode(decodedBinarySecretBytes, secretResult.SecretBinary)
		if err != nil {
			log.Errorln("Base64 Decode Error:", err)
			return "", false, err
		}
		result = string(decodedBinarySecretBytes[:length])
	}

	isJSON = isStringJSON(result)

	return result, isJSON, nil
}
