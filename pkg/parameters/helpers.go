package parameters

import (
	"flag"
	fl "github.com/barchart/common-go/pkg/parameters/flags"
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
	if !flg.Value.(*fl.StringValue).IsSet() {
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
			return flg.Value.(*fl.BoolValue).Get(), flg.Value.(*fl.BoolValue).IsSet()
		}
	case float64Type:
		{
			return flg.Value.(*fl.Float64Value).Get(), flg.Value.(*fl.Float64Value).IsSet()
		}
	case intType:
		{
			return flg.Value.(*fl.IntValue).Get(), flg.Value.(*fl.IntValue).IsSet()
		}
	case int64Type:
		{
			return flg.Value.(*fl.Int64Value).Get(), flg.Value.(*fl.Int64Value).IsSet()
		}
	case stringType:
		{
			return flg.Value.(*fl.StringValue).Get(), flg.Value.(*fl.StringValue).IsSet()
		}
	case uintType:
		{
			return flg.Value.(*fl.UintValue).Get(), flg.Value.(*fl.UintValue).IsSet()
		}
	case uint64Type:
		{
			return flg.Value.(*fl.Uint64Value).Get(), flg.Value.(*fl.Uint64Value).IsSet()
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
