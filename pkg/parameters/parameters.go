package parameters

import (
	"flag"
	"fmt"
	"github.com/barchart/barchart-common-go/pkg/configuration"
	"github.com/barchart/barchart-common-go/pkg/configuration/aws/secretsmanager"
	"github.com/barchart/barchart-common-go/pkg/parameters/flags"
	"log"
	"os"
	"strings"
)

type parameter struct {
	Name         string
	DefaultValue interface{}
	Usage        string
	Required     bool
	AWSSecret    bool
	valueType    string
}

type parameters struct {
	collection map[string]parameter
}

var (
	sm      *secretsmanager.SecretsManager
	smError error
)

func New() *parameters {
	return &parameters{collection: map[string]parameter{}}
}

func (p *parameters) Add(name string, value string, usage string, required bool, awsSecret ...bool) {
	p.AddString(name, value, usage, required, awsSecret...)
}

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

func (p parameters) Parse() map[string]interface{} {
	result := make(map[string]interface{})
	missing := make([]string, 0, 1)

	flags.String(AwsRegionSecrets, "us-east-1", "The AWS Secrets Manager region")

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
				log.Printf("PARAM: %v: %T", param.Name, envValue)
				result[param.Name] = envValue
			} else {
				secretValue := getValueFromAWSSecretsManager(param)
				if secretValue != nil {
					log.Printf("PARAM: %v: %T", param.Name, secretValue)
					result[param.Name] = secretValue
				} else {
					if param.Required {
						missing = append(missing, param.Name)
					} else {
						log.Printf("PARAM: %v: %T", param.Name, value)
						result[param.Name] = value
					}
				}
			}
		} else {
			log.Printf("PARAM: %v: %T", param.Name, value)
			result[param.Name] = value
		}
	}

	if len(missing) > 0 {
		panic(fmt.Sprintf("missing required parameters: [ %v ]", strings.Join(missing, ",")))
	}

	return result
}
