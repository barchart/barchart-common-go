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
	"log"
)

type cache struct {
	Value  string
	isJSON bool
}

type SecretsManager struct {
	Region string `validate:"required"`
	sm     *secretsmanager.SecretsManager
	cache  map[string]cache
}

func isStringJSON(str string) bool {
	var jsonStr map[string]interface{}
	err := json.Unmarshal([]byte(str), &jsonStr)
	return err == nil
}

func New(region string) *SecretsManager {
	secretsManager := SecretsManager{}
	secretsManager.Region = region

	if secretsManager.cache == nil {
		secretsManager.cache = map[string]cache{}
	}

	sess, err := session.NewSession()

	if err != nil {
		log.Println("err " + err.Error())
	}

	secretsManager.sm = secretsmanager.New(sess, aws.NewConfig().WithRegion(region))

	return &secretsManager
}

func (secretsManager SecretsManager) ClearCache() {
	secretsManager.cache = map[string]cache{}
}

func (secretsManager SecretsManager) GetValue(secretName string) (string, bool, error) {
	cached := secretsManager.cache[secretName]
	if cached.Value != "" {
		return cached.Value, cached.isJSON, nil
	}

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
			fmt.Println(err.Error())
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
			fmt.Println("Base64 Decode Error:", err)
			return "", false, err
		}
		result = string(decodedBinarySecretBytes[:length])
	}
	isJSON = isStringJSON(result)
	secretsManager.cache[secretName] = cache{Value: result, isJSON: isJSON}

	return result, isJSON, nil
}
