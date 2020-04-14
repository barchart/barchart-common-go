package parameters

import (
	"flag"
	"fmt"
	"github.com/barchart/barchart-common-go/pkg/configuration"
	_ "github.com/barchart/barchart-common-go/pkg/configuration"
	fl "github.com/barchart/barchart-common-go/pkg/parameters/flags"
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

	fl.Bool(name, value, usage)
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

	fl.Float64(name, value, usage)
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

	fl.Int(name, value, usage)
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

	fl.Int64(name, value, usage)
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

	fl.String(name, value, usage)
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

	fl.Uint(name, value, usage)
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
	fl.Uint64(name, value, usage)
}

func getAWSSecretsRegion() string {
	flg := flag.Lookup(AwsRegionSecrets)
	region := flg.Value.String()
	if !flg.Value.(*fl.StringValue).IsSet() {
		region = os.Getenv(AwsRegionSecrets)
		if region == "" {
			region = flg.DefValue
		}
	}

	return region
}

func (p parameters) Parse() map[string]interface{} {
	result := make(map[string]interface{})
	missing := make([]string, 0, 1)

	fl.String(AwsRegionSecrets, "us-east-1", "The AWS Secrets Manager region")

	flag.Parse()

	config := configuration.GetConfig()
	config.SetSecretsManager(getAWSSecretsRegion())
	sm, smError := configuration.GetConfig().GetSecretsManager()

	if p.collection == nil {
		panic("parameters wasn't added")
	}

	for _, param := range p.collection {
		flg := flag.Lookup(param.Name)
		flagValueString := flg.Value.String()
		flagValue := convertString(flagValueString, param.valueType)

		envValueString := os.Getenv(param.Name)
		envValue := convertString(envValueString, param.valueType)

		switch param.valueType {
		case boolType:
			{
				if param.AWSSecret {
					if smError == nil {
						value, _, err := sm.GetValue(param.Name)
						if err == nil {
							result[param.Name] = convertString(value, param.valueType)
							continue
						}
					}
				}

				if flg.Value.(*fl.BoolValue).IsSet() {
					result[param.Name] = flagValue
				} else {
					if envValueString != "" {
						result[param.Name] = envValue
					} else {
						if param.Required {
							missing = append(missing, param.Name)
						} else {
							result[param.Name] = flagValue
						}
					}
				}
			}
		case float64Type:
			{
				if param.AWSSecret {
					if smError == nil {
						value, _, err := sm.GetValue(param.Name)
						if err == nil {
							result[param.Name] = convertString(value, param.valueType)
							continue
						}
					}
				}

				if flg.Value.(*fl.Float64Value).IsSet() {
					result[param.Name] = flagValue
				} else {
					if envValueString != "" {
						result[param.Name] = envValue
					} else {
						if param.Required {
							missing = append(missing, param.Name)
						} else {
							result[param.Name] = flagValue
						}
					}
				}
			}
		case intType:
			{
				if param.AWSSecret {
					if smError == nil {
						value, _, err := sm.GetValue(param.Name)
						if err == nil {
							result[param.Name] = convertString(value, param.valueType)
							continue
						}
					}
				}

				if flg.Value.(*fl.IntValue).IsSet() {
					result[param.Name] = flagValue
				} else {
					if envValueString != "" {
						result[param.Name] = envValue
					} else {
						if param.Required {
							missing = append(missing, param.Name)
						} else {
							result[param.Name] = flagValue
						}
					}
				}
			}
		case int64Type:
			{
				if param.AWSSecret {
					if smError == nil {
						value, _, err := sm.GetValue(param.Name)
						if err == nil {
							result[param.Name] = convertString(value, param.valueType)
							continue
						}
					}
				}

				if flg.Value.(*fl.Int64Value).IsSet() {
					result[param.Name] = flagValue
				} else {
					if envValueString != "" {
						result[param.Name] = envValue
					} else {
						if param.Required {
							missing = append(missing, param.Name)
						} else {
							result[param.Name] = flagValue
						}
					}
				}
			}
		case stringType:
			{
				if param.AWSSecret {
					if smError == nil {
						value, _, err := sm.GetValue(param.Name)
						if err == nil {
							result[param.Name] = convertString(value, param.valueType)
							continue
						}
					}
				}

				if flg.Value.(*fl.StringValue).IsSet() {
					result[param.Name] = flagValue
				} else {
					if envValueString != "" {
						result[param.Name] = envValue
					} else {
						if param.Required {
							missing = append(missing, param.Name)
						} else {
							result[param.Name] = flagValue
						}
					}
				}
			}
		case uintType:
			{
				if param.AWSSecret {
					if smError == nil {
						value, _, err := sm.GetValue(param.Name)
						if err == nil {
							result[param.Name] = convertString(value, param.valueType)
							continue
						}
					}
				}

				if flg.Value.(*fl.UintValue).IsSet() {
					result[param.Name] = flagValue
				} else {
					if envValueString != "" {
						result[param.Name] = envValue
					} else {
						if param.Required {
							missing = append(missing, param.Name)
						} else {
							result[param.Name] = flagValue
						}
					}
				}
			}
		case uint64Type:
			{
				if param.AWSSecret {
					if smError == nil {
						value, _, err := sm.GetValue(param.Name)
						if err == nil {
							result[param.Name] = convertString(value, param.valueType)
							continue
						}
					}
				}

				if flg.Value.(*fl.Uint64Value).IsSet() {
					result[param.Name] = flagValue
				} else {
					if envValueString != "" {
						result[param.Name] = envValue
					} else {
						if param.Required {
							missing = append(missing, param.Name)
						} else {
							result[param.Name] = flagValue
						}
					}
				}
			}
		}
	}

	if len(missing) > 0 {
		panic(fmt.Sprintf("missing required parameters: [ %v ]", strings.Join(missing, ",")))
	}

	return result
}
