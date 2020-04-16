package parameters

import (
	"flag"
	"fmt"
	"github.com/barchart/common-go/pkg/configuration"
	"github.com/barchart/common-go/pkg/configuration/aws/secretsmanager"
	"github.com/barchart/common-go/pkg/parameters/flags"
	"os"
	"strings"
)

// parameter is a struct defines a parameter
type parameter struct {
	Name         string
	DefaultValue interface{}
	Usage        string
	Required     bool
	AWSSecret    bool
	valueType    string
}

// parameters is a struct that holds the collection of parameter
type parameters struct {
	collection map[string]parameter
}

var (
	sm      *secretsmanager.SecretsManager
	smError error
)

// New return instance of parameters
func New() *parameters {
	return &parameters{collection: map[string]parameter{}}
}

// Add is alias for AddString
func (p *parameters) Add(name string, value string, usage string, required bool, awsSecret ...bool) {
	p.AddString(name, value, usage, required, awsSecret...)
}

// AddBool defines a bool parameter with specified name, default value, and usage string.
func (p *parameters) AddBool(name string, value bool, usage string, required bool, awsSecret ...bool) {
	p.collection[name] = parameter{
		Name:         name,
		DefaultValue: value,
		Usage:        usage,
		Required:     required,
		AWSSecret:    isAWSSecret(awsSecret),
		valueType:    boolType,
	}

	flags.Bool(name, value, usage)
}

// AddFloat64 defines a float64 parameter with specified name, default value, and usage string.
func (p *parameters) AddFloat64(name string, value float64, usage string, required bool, awsSecret ...bool) {
	p.collection[name] = parameter{
		Name:         name,
		DefaultValue: value,
		Usage:        usage,
		Required:     required,
		AWSSecret:    isAWSSecret(awsSecret),
		valueType:    float64Type,
	}

	flags.Float64(name, value, usage)
}

// AddInt defines a int parameter with specified name, default value, and usage string.
func (p *parameters) AddInt(name string, value int, usage string, required bool, awsSecret ...bool) {
	p.collection[name] = parameter{
		Name:         name,
		DefaultValue: value,
		Usage:        usage,
		Required:     required,
		AWSSecret:    isAWSSecret(awsSecret),
		valueType:    intType,
	}

	flags.Int(name, value, usage)
}

// AddInt64 defines a int64 parameter with specified name, default value, and usage string.
func (p *parameters) AddInt64(name string, value int64, usage string, required bool, awsSecret ...bool) {
	p.collection[name] = parameter{
		Name:         name,
		DefaultValue: value,
		Usage:        usage,
		Required:     required,
		AWSSecret:    isAWSSecret(awsSecret),
		valueType:    int64Type,
	}

	flags.Int64(name, value, usage)
}

// AddString defines a string parameter with specified name, default value, and usage string.
func (p *parameters) AddString(name string, value string, usage string, required bool, awsSecret ...bool) {
	p.collection[name] = parameter{
		Name:         name,
		DefaultValue: value,
		Usage:        usage,
		Required:     required,
		AWSSecret:    isAWSSecret(awsSecret),
		valueType:    stringType,
	}

	flags.String(name, value, usage)
}

// AddUint defines a uint parameter with specified name, default value, and usage string.
func (p *parameters) AddUint(name string, value uint, usage string, required bool, awsSecret ...bool) {
	p.collection[name] = parameter{
		Name:         name,
		DefaultValue: value,
		Usage:        usage,
		Required:     required,
		AWSSecret:    isAWSSecret(awsSecret),
		valueType:    uintType,
	}

	flags.Uint(name, value, usage)
}

// AddUint64 defines a uint64 parameter with specified name, default value, and usage string.
func (p *parameters) AddUint64(name string, value uint64, usage string, required bool, awsSecret ...bool) {
	p.collection[name] = parameter{
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
func (p parameters) Parse() map[string]interface{} {
	result := make(map[string]interface{})
	missing := make([]string, 0, 1)

	flags.String(AwsRegionSecrets, "us-east-1", "The AWS Secrets Manager region")

	if flag.Parsed() {
		panic("flags have already parsed")
	}

	flag.Parse()

	config := configuration.GetConfig()
	config.SetSecretsManager(getAWSSecretsRegion())
	smm, smmError := configuration.GetConfig().GetSecretsManager()
	sm = &smm
	smError = smmError

	if p.collection == nil {
		panic("parameters wasn't added")
	}

	for _, param := range p.collection {
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
		panic(fmt.Sprintf("missing required parameters: [ %v ]", strings.Join(missing, ",")))
	}

	return result
}
