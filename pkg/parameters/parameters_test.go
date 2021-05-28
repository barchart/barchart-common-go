package parameters

import (
	"flag"
	"os"
	"strconv"
	"testing"

	"github.com/barchart/common-go/pkg/configuration/database"
	"github.com/stretchr/testify/assert"
)

const (
	expectAdd          = "ADD"
	expectString       = "STRING"
	expectBool         = "true"
	expectInt          = "100"
	expectInt64        = "100"
	expectFloat64      = "100.100"
	expectUint         = "100"
	expectUint64       = "100"
	expectDefaultField = "100"
	expectDatabase     = "{\"provider\":\"mysql\",\"host\":\"https://example.com\",\"port\":54321,\"database\":\"database\",\"username\":\"user\",\"password\":\"password\"}"
)

const (
	expectedInt          = 100
	expectedInt64        = int64(100)
	expectedFloat64      = float64(100.100)
	expectedUint         = uint(100)
	expectedUint64       = uint64(100)
	expectedBool         = true
	expectedAdd          = "ADD"
	expectedString       = "STRING"
	expectedDefaultField = "100"
)

var expectedDatabase = database.Database{
	Provider: "mysql",
	Host:     "https://example.com",
	Port:     54321,
	Database: "database",
	Username: "user",
	Password: "password",
}

const (
	keyInt           = "INT"
	keyInt64         = "INT64"
	keyString        = "STRING"
	keyBool          = "BOOL"
	keyFloat64       = "FLOAT64"
	keyUint          = "UINT"
	keyUint64        = "UINT64"
	keyAdd           = "ADD"
	keyDatabase      = "DATABASE"
	keyDefaultField  = "DEFAULT_FIELD"
	keyRequiredField = "REQUIRED_FIELD"
	keyNotExist      = "NOT_EXIST"
)

var result Results
var required, _ = strconv.ParseBool(os.Getenv("REQUIRED"))

func setup() {
	manual, _ := strconv.ParseBool(os.Getenv("MANUAL"))

	if !manual {
		_ = flag.Set(keyAdd, expectAdd)
		_ = flag.Set(keyString, expectString)
		_ = flag.Set(keyBool, expectBool)
		_ = flag.Set(keyInt, expectInt)
		_ = flag.Set(keyInt64, expectInt64)
		_ = flag.Set(keyFloat64, expectFloat64)
		_ = flag.Set(keyUint, expectUint)
		_ = flag.Set(keyDatabase, expectDatabase)
		_ = os.Setenv(keyUint64, expectUint64)
	}
}

func TestMain(m *testing.M) {
	Add(keyDefaultField, expectDefaultField, "The default Parameter", false)
	Add(keyAdd, "default", "The string Parameter", false)
	AddString(keyString, "default", "The string Parameter", false)
	AddBool(keyBool, false, "The bool Parameter from env", false)
	AddInt(keyInt, 50, "The int Parameter", false)
	AddInt64(keyInt64, 50, "The int64 Parameter", false)
	AddFloat64(keyFloat64, 50.50, "The float64 Parameter", false)
	AddUint(keyUint, 50, "The uint Parameter", false)
	AddUint64(keyUint64, 50, "The uint64 Parameter", false)
	AddDatabase(keyDatabase, database.Database{}, "The database Parameter", false)

	if required {
		AddBool(keyRequiredField, false, "The required Parameter", true)
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %s", err)
			}
		}()
	}

	setup()

	result = Parse()
	os.Exit(m.Run())
}

func TestParameters_Basic(t *testing.T) {
	isParsed := Parsed()
	assert.True(t, isParsed, "must return true after execution of Parse() function")

	newResult := GetResults()
	assert.Equal(t, result, newResult, "get result function must return parsed data")

	doubleParsedResult := Parse()
	assert.Equal(t, result, doubleParsedResult, "parse must return results from the first execution Parse() function")
}

func TestResultsGetValueDirectly(t *testing.T) {
	assert.Equal(t, expectedString, result[keyString])
	assert.Equal(t, expectedInt, result[keyInt])
	assert.Equal(t, expectedInt64, result[keyInt64])
	assert.Equal(t, expectedFloat64, result[keyFloat64])
	assert.Equal(t, expectedBool, result[keyBool])
	assert.Equal(t, expectedAdd, result[keyAdd])
	assert.Equal(t, expectedUint, result[keyUint])
	assert.Equal(t, expectedUint64, result[keyUint64])
	assert.Equal(t, expectedDefaultField, result[keyDefaultField])
	assert.Equal(t, expectedDatabase, result[keyDatabase])
}

func TestResults_GetString(t *testing.T) {
	assert.Equal(t, expectedString, result.GetString(keyString), "should return a string value from results")
}

func TestResults_GetInt(t *testing.T) {
	assert.Equal(t, expectedInt, result.GetInt(keyInt), "should return an int value from results")
}

func TestResults_GetInt64(t *testing.T) {
	assert.Equal(t, expectedInt64, result.GetInt64(keyInt64), "should return an int64 value from results")
}

func TestResults_GetFloat64(t *testing.T) {
	assert.Equal(t, expectedFloat64, result.GetFloat64(keyFloat64), "should return a float64 value from results")
}

func TestResults_GetBool(t *testing.T) {
	assert.Equal(t, expectedBool, result.GetBool(keyBool), "should return a bool value from results")
}

func TestResults_GetUint(t *testing.T) {
	assert.Equal(t, expectedUint, result.GetUint(keyUint), "should return a uint value from results")
}

func TestResults_GetUint64(t *testing.T) {
	assert.Equal(t, expectedUint64, result.GetUint64(keyUint64), "should return a uint64 value from results")
}

func TestResults_GetDatabase(t *testing.T) {
	assert.Equalf(t, expectedDatabase, result.GetDatabase(keyDatabase), "should return a database value from results")
}

func TestResults_GetStringSafe(t *testing.T) {
	value, err := result.GetStringSafe(keyString)
	assert.Equal(t, expectedString, value, "should return a string value from result")
	assert.Nil(t, err, "an error should be nil")
}

func TestResults_GetIntSafe(t *testing.T) {
	value, err := result.GetIntSafe(keyInt)
	assert.Equal(t, expectedInt, value, "should return an int value from result")
	assert.Nil(t, err, "an error should be nil")
}

func TestResults_GetInt64Safe(t *testing.T) {
	value, err := result.GetInt64Safe(keyInt64)
	assert.Equal(t, expectedInt64, value, "should return an int64 value from result")
	assert.Nil(t, err, "an error should be nil")
}

func TestResults_GetFloat64Safe(t *testing.T) {
	value, err := result.GetFloat64Safe(keyFloat64)
	assert.Equal(t, expectedFloat64, value, "should return a float64 value from result")
	assert.Nil(t, err, "an error should be nil")
}

func TestResults_GetBoolSafe(t *testing.T) {
	value, err := result.GetBoolSafe(keyBool)
	assert.Equal(t, expectedBool, value, "should return a bool value from result")
	assert.Nil(t, err, "an error should be nil")
}

func TestResults_GetUintSafe(t *testing.T) {
	value, err := result.GetUintSafe(keyUint)
	assert.Equal(t, expectedUint, value, "should return a uint value from result")
	assert.Nil(t, err, "an error should be nil")
}

func TestResults_GetUint64Safe(t *testing.T) {
	value, err := result.GetUint64Safe(keyUint64)
	assert.Equal(t, expectedUint64, value, "should return a uint64 value from result")
	assert.Nil(t, err, "an error should be nil")
}

func TestResults_GetDatabaseSafe(t *testing.T) {
	value, err := result.GetDatabaseSafe(keyDatabase)
	assert.Equal(t, expectedDatabase, value, "should return a database value from result")
	assert.Nil(t, err, "an error should be nil")
}

func TestResults_GetStringNotExist(t *testing.T) {
	value, err := result.GetStringSafe(keyNotExist)
	assert.Equal(t, "", value, "should return a default value")
	assert.NotNil(t, err, "an error should not be nil")
}

func TestResults_GetIntSafeNotExist(t *testing.T) {
	value, err := result.GetIntSafe(keyNotExist)
	assert.Equal(t, 0, value, "should return a default value")
	assert.NotNil(t, err, "an error should not be nil")
}

func TestResults_GetInt64SafeNotExist(t *testing.T) {
	value, err := result.GetInt64Safe(keyNotExist)
	assert.Equal(t, int64(0), value, "should return a default value")
	assert.NotNil(t, err, "an error should not be nil")
}

func TestResults_GetFloat64SafeNotExist(t *testing.T) {
	value, err := result.GetFloat64Safe(keyNotExist)
	assert.Equal(t, float64(0), value, "should return a default value")
	assert.NotNil(t, err, "an error should not be nil")
}

func TestResults_GetBoolSafeNotExist(t *testing.T) {
	value, err := result.GetBoolSafe(keyNotExist)
	assert.Equal(t, false, value, "should return a default value")
	assert.NotNil(t, err, "an error should not be nil")
}

func TestResults_GetUintSafeNotExist(t *testing.T) {
	value, err := result.GetUintSafe(keyNotExist)
	assert.Equal(t, uint(0), value, "should return a default value")
	assert.NotNil(t, err, "an error should not be nil")
}

func TestResults_GetUint64SafeNotExist(t *testing.T) {
	value, err := result.GetUint64Safe(keyNotExist)
	assert.Equal(t, uint64(0), value, "should return a default value")
	assert.NotNil(t, err, "an error should not be nil")
}

func TestResults_GetDatabaseSafeNotExist(t *testing.T) {
	value, err := result.GetDatabaseSafe(keyNotExist)
	assert.Equal(t, database.Database{}, value, "should return a default value")
	assert.NotNil(t, err, "an error should not be nil")
}
