package parameters

import (
	"flag"
	"github.com/barchart/common-go/pkg/configuration"
	"github.com/barchart/common-go/pkg/configuration/aws/secretsmanager"
	"github.com/barchart/common-go/pkg/logger"
	"github.com/barchart/common-go/pkg/parameters/flags"
	"os"
	"strings"
)

// Parameter is a struct defines a Parameter
type Parameter struct {
	Name         string
	DefaultValue interface{}
	Usage        string
	Required     bool
	AWSSecret    bool
	valueType    string
}

// parameters is a struct that holds the collection of Parameter
type parameters struct {
	collection map[string]Parameter
}

var (
	sm            *secretsmanager.SecretsManager
	smError       error
	log           = logger.Log
	defaultParams = &parameters{collection: map[string]Parameter{}}
)

func init() {

}

// Add is alias for AddString
func Add(name string, value string, usage string, required bool, awsSecret ...bool) {
	AddString(name, value, usage, required, awsSecret...)
}

// AddBool defines a bool Parameter with specified name, default value, and usage string.
func AddBool(name string, value bool, usage string, required bool, awsSecret ...bool) {
	defaultParams.collection[name] = Parameter{
		Name:         name,
		DefaultValue: value,
		Usage:        usage,
		Required:     required,
		AWSSecret:    isAWSSecret(awsSecret),
		valueType:    boolType,
	}

	flags.Bool(name, value, usage)
}

// AddFloat64 defines a float64 Parameter with specified name, default value, and usage string.
func AddFloat64(name string, value float64, usage string, required bool, awsSecret ...bool) {
	defaultParams.collection[name] = Parameter{
		Name:         name,
		DefaultValue: value,
		Usage:        usage,
		Required:     required,
		AWSSecret:    isAWSSecret(awsSecret),
		valueType:    float64Type,
	}

	flags.Float64(name, value, usage)
}

// AddInt defines a int Parameter with specified name, default value, and usage string.
func AddInt(name string, value int, usage string, required bool, awsSecret ...bool) {
	defaultParams.collection[name] = Parameter{
		Name:         name,
		DefaultValue: value,
		Usage:        usage,
		Required:     required,
		AWSSecret:    isAWSSecret(awsSecret),
		valueType:    intType,
	}

	flags.Int(name, value, usage)
}

// AddInt64 defines a int64 Parameter with specified name, default value, and usage string.
func AddInt64(name string, value int64, usage string, required bool, awsSecret ...bool) {
	defaultParams.collection[name] = Parameter{
		Name:         name,
		DefaultValue: value,
		Usage:        usage,
		Required:     required,
		AWSSecret:    isAWSSecret(awsSecret),
		valueType:    int64Type,
	}

	flags.Int64(name, value, usage)
}

// AddString defines a string Parameter with specified name, default value, and usage string.
func AddString(name string, value string, usage string, required bool, awsSecret ...bool) {
	defaultParams.collection[name] = Parameter{
		Name:         name,
		DefaultValue: value,
		Usage:        usage,
		Required:     required,
		AWSSecret:    isAWSSecret(awsSecret),
		valueType:    stringType,
	}

	flags.String(name, value, usage)
}

// AddUint defines a uint Parameter with specified name, default value, and usage string.
func AddUint(name string, value uint, usage string, required bool, awsSecret ...bool) {
	defaultParams.collection[name] = Parameter{
		Name:         name,
		DefaultValue: value,
		Usage:        usage,
		Required:     required,
		AWSSecret:    isAWSSecret(awsSecret),
		valueType:    uintType,
	}

	flags.Uint(name, value, usage)
}

// AddUint64 defines a uint64 Parameter with specified name, default value, and usage string.
func AddUint64(name string, value uint64, usage string, required bool, awsSecret ...bool) {
	defaultParams.collection[name] = Parameter{
		Name:         name,
		DefaultValue: value,
		Usage:        usage,
		Required:     required,
		AWSSecret:    isAWSSecret(awsSecret),
		valueType:    uint64Type,
	}
	flags.Uint64(name, value, usage)
}

// Parse returns map of values of all defined parameters
func Parse() map[string]interface{} {
	result := make(map[string]interface{})
	missing := make([]string, 0, 1)

	flags.String(AwsRegionSecrets, "us-east-1", "The AWS Secrets Manager region")

	if flag.Parsed() {
		log.Panic("flags have already parsed")
	}

	flag.Parse()

	configuration.SetSecretsManager(getAWSSecretsRegion())
	smm, smmError := configuration.GetSecretsManager()
	sm = &smm
	smError = smmError

	if defaultParams.collection == nil {
		log.Panic("parameters wasn't added")
	}

	for _, param := range defaultParams.collection {
		flg := flag.Lookup(param.Name)
		value, isSet := getValueFromFlag(flg, param.valueType)

		if !isSet {
			envValueString := os.Getenv(param.Name)
			if envValueString != "" {
				envValue := convertString(envValueString, param.valueType)
				result[param.Name] = envValue
			} else {
				secretValue := getValueFromAWSSecretsManager(param)
				if secretValue != nil {
					result[param.Name] = secretValue
				} else {
					if param.Required {
						missing = append(missing, param.Name)
					} else {
						result[param.Name] = value
					}
				}
			}
		} else {
			result[param.Name] = value
		}
	}

	if len(missing) > 0 {
		log.Panicf("missing required parameters: [ %v ]", strings.Join(missing, ","))
	}

	return result
}

func GetCollection() map[string]Parameter {
	return defaultParams.collection
}
