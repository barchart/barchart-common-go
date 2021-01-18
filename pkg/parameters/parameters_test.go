package parameters

import (
	"flag"
	"os"
	"strconv"
	"testing"

	"github.com/barchart/common-go/pkg/configuration/database"
	. "github.com/smartystreets/goconvey/convey"
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

func TestParametersBasic(t *testing.T) {
	Convey("Test basic functions", t, func() {
		Convey("Parsed", func() {
			Convey("must return true after execution of Parse() function", func() {
				isParsed := Parsed()
				So(isParsed, ShouldEqual, true)
			})
		})

		Convey("Get results", func() {
			Convey("get result function must return parsed data", func() {
				newResult := GetResults()
				So(newResult, ShouldResemble, result)
			})
		})

		Convey("Parse params twice", func() {
			Convey("parse must return results from the first execution Parse() function", func() {
				newResult := Parse()
				So(newResult, ShouldResemble, result)
			})
		})
	})
}

func TestParametersGet(t *testing.T) {
	Convey("Test getting values", t, func() {
		Convey("Get values directly from map[string]interface{}", func() {
			Convey("test string", func() {
				So(result[keyString], ShouldEqual, expectedString)
			})
			Convey("test int", func() {
				So(result[keyInt], ShouldEqual, expectedInt)
			})
			Convey("test int64", func() {
				So(result[keyInt64], ShouldEqual, expectedInt64)
			})
			Convey("test float64", func() {
				So(result[keyFloat64], ShouldEqual, expectedFloat64)
			})
			Convey("test bool", func() {
				So(result[keyBool], ShouldEqual, expectedBool)
			})
			Convey("test add (string)", func() {
				So(result[keyAdd], ShouldEqual, expectedAdd)
			})
			Convey("test uint ", func() {
				So(result[keyUint], ShouldEqual, expectedUint)
			})
			Convey("test uint64", func() {
				So(result[keyUint64], ShouldEqual, expectedUint64)
			})
			Convey("test default field", func() {
				So(result[keyDefaultField], ShouldEqual, expectedDefaultField)
			})
			Convey("test database", func() {
				So(result[keyDatabase], ShouldResemble, expectedDatabase)
			})
		})

		Convey("Get values by using functions", func() {
			Convey("test GetString", func() {
				So(result.GetString(keyString), ShouldEqual, expectedString)
			})
			Convey("test GetInt", func() {
				So(result.GetInt(keyInt), ShouldEqual, expectedInt)
			})
			Convey("test GetInt64", func() {
				So(result.GetInt64(keyInt64), ShouldEqual, expectedInt64)
			})
			Convey("test GetFloat64", func() {
				So(result.GetFloat64(keyFloat64), ShouldEqual, expectedFloat64)
			})
			Convey("test GetBool", func() {
				So(result.GetBool(keyBool), ShouldEqual, expectedBool)
			})
			Convey("test GetUint ", func() {
				So(result.GetUint(keyUint), ShouldEqual, expectedUint)
			})
			Convey("test GetUint64", func() {
				So(result.GetUint64(keyUint64), ShouldEqual, expectedUint64)
			})
			Convey("test GetDatabase", func() {
				So(result.GetDatabase(keyDatabase), ShouldResemble, expectedDatabase)
			})
		})

		Convey("Get values by using safe functions", func() {
			Convey("a value must exist, an error must be nil", func() {
				Convey("test GetStringSafe", func() {
					value, err := result.GetStringSafe(keyString)
					So(value, ShouldEqual, expectedString)
					So(err, ShouldEqual, nil)
				})
				Convey("test GetIntSafe", func() {
					value, err := result.GetIntSafe(keyInt)
					So(value, ShouldEqual, expectedInt)
					So(err, ShouldEqual, nil)
				})
				Convey("test GetInt64Safe", func() {
					value, err := result.GetInt64Safe(keyInt64)
					So(value, ShouldEqual, expectedInt64)
					So(err, ShouldEqual, nil)
				})
				Convey("test GetFloat64Safe", func() {
					value, err := result.GetFloat64Safe(keyFloat64)
					So(value, ShouldEqual, expectedFloat64)
					So(err, ShouldEqual, nil)
				})
				Convey("test GetBoolSafe", func() {
					value, err := result.GetBoolSafe(keyBool)
					So(value, ShouldEqual, expectedBool)
					So(err, ShouldEqual, nil)
				})
				Convey("test GetUintSafe ", func() {
					value, err := result.GetUintSafe(keyUint)
					So(value, ShouldEqual, expectedUint)
					So(err, ShouldEqual, nil)
				})
				Convey("test GetUint64Safe", func() {
					value, err := result.GetUint64Safe(keyUint64)
					So(value, ShouldEqual, expectedUint64)
					So(err, ShouldEqual, nil)
				})
				Convey("test GetDatabaseSafe", func() {
					value, err := result.GetDatabaseSafe(keyDatabase)
					So(value, ShouldResemble, expectedDatabase)
					So(err, ShouldEqual, nil)
				})
			})

			Convey("a value must be equal to the default value for a variable type, an error mustn't be equal to nil ", func() {
				Convey("test GetStringSafe", func() {
					value, err := result.GetStringSafe(keyNotExist)
					So(value, ShouldEqual, "")
					So(err, ShouldNotEqual, nil)
				})
				Convey("test GetIntSafe", func() {
					value, err := result.GetIntSafe(keyNotExist)
					So(value, ShouldEqual, 0)
					So(err, ShouldNotEqual, nil)
				})
				Convey("test GetInt64Safe", func() {
					value, err := result.GetInt64Safe(keyNotExist)
					So(value, ShouldEqual, 0)
					So(err, ShouldNotEqual, nil)
				})
				Convey("test GetFloat64Safe", func() {
					value, err := result.GetFloat64Safe(keyNotExist)
					So(value, ShouldEqual, 0)
					So(err, ShouldNotEqual, nil)
				})
				Convey("test GetBoolSafe", func() {
					value, err := result.GetBoolSafe(keyNotExist)
					So(value, ShouldEqual, false)
					So(err, ShouldNotEqual, nil)
				})
				Convey("test GetUintSafe ", func() {
					value, err := result.GetUintSafe(keyNotExist)
					So(value, ShouldEqual, 0)
					So(err, ShouldNotEqual, nil)
				})
				Convey("test GetUint64Safe", func() {
					value, err := result.GetUint64Safe(keyNotExist)
					So(value, ShouldEqual, 0)
					So(err, ShouldNotEqual, nil)
				})
				Convey("test GetDatabaseSafe", func() {
					value, err := result.GetDatabaseSafe(keyNotExist)
					So(value, ShouldResemble, database.Database{})
					So(err, ShouldNotEqual, nil)
				})
			})
		})
	})
}
