package parameters

import (
	"strconv"
)

func isAWSSecret(secretManager []bool) bool {
	awsSecret := false
	if len(secretManager) > 0 && secretManager[0] == true {
		awsSecret = true
	}
	return awsSecret
}

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
