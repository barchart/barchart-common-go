package parameters

import (
	"flag"
	"github.com/barchart/common-go/pkg/parameters/flags"
	"os"
	"strconv"
)

// convertStrings converts a string to a typeValue and returns it
func convertString(str string, typeValue string) interface{} {
	switch typeValue {
	case boolType:
		{
			value, _ := strconv.ParseBool(str)
			return value
		}
	case float64Type:
		{
			value, _ := strconv.ParseFloat(str, 64)
			return value
		}
	case intType:
		{
			value, _ := strconv.ParseInt(str, 10, 32)
			return int(value)
		}
	case int64Type:
		{
			value, _ := strconv.ParseInt(str, 10, 64)
			return value
		}
	case stringType:
		{
			return str
		}
	case uintType:
		{
			value, _ := strconv.ParseUint(str, 10, 32)
			return uint(value)
		}
	case uint64Type:
		{
			value, _ := strconv.ParseUint(str, 10, 64)
			return value
		}
	}

	return str
}

// getAWSSecretsRegion gets AwsRegionSecrets value from a AWS-REGION-SECRETS flag or env variable and returns it
func getAWSSecretsRegion() string {
	flg := flag.Lookup(AwsRegionSecrets)
	region := flg.Value.String()
	if !flg.Value.(*flags.StringValue).IsSet() {
		region = os.Getenv(AwsRegionSecrets)
		if region == "" {
			region = flg.DefValue
		}
	}

	return region
}

// getValueFromAWSSecretsManager returns a parameter value from the AWS Secrets Manager
func getValueFromAWSSecretsManager(param parameter) interface{} {
	if param.AWSSecret {
		if sm != nil && smError == nil {
			value, _, err := sm.GetValue(param.Name)
			if err == nil {
				return convertString(value, param.valueType)
			}
		}
	}

	return nil
}

// getValueFromFlag returns the value of the desired type from the flag
func getValueFromFlag(flg *flag.Flag, typeValue string) (interface{}, bool) {
	switch typeValue {
	case boolType:
		{
			return flg.Value.(*flags.BoolValue).Get(), flg.Value.(*flags.BoolValue).IsSet()
		}
	case float64Type:
		{
			return flg.Value.(*flags.Float64Value).Get(), flg.Value.(*flags.Float64Value).IsSet()
		}
	case intType:
		{
			return flg.Value.(*flags.IntValue).Get(), flg.Value.(*flags.IntValue).IsSet()
		}
	case int64Type:
		{
			return flg.Value.(*flags.Int64Value).Get(), flg.Value.(*flags.Int64Value).IsSet()
		}
	case stringType:
		{
			return flg.Value.(*flags.StringValue).Get(), flg.Value.(*flags.StringValue).IsSet()
		}
	case uintType:
		{
			return flg.Value.(*flags.UintValue).Get(), flg.Value.(*flags.UintValue).IsSet()
		}
	case uint64Type:
		{
			return flg.Value.(*flags.Uint64Value).Get(), flg.Value.(*flags.Uint64Value).IsSet()
		}
	}

	return nil, false
}

// isAWSSecret returns is AWS Secrets Manager argument was provided
func isAWSSecret(awsSecret []bool) bool {
	isAwsSecret := false
	if len(awsSecret) > 0 && awsSecret[0] == true {
		isAwsSecret = true
	}
	return isAwsSecret
}
