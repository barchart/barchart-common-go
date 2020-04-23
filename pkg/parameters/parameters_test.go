package parameters

import (
	"flag"
	"github.com/stretchr/testify/require"
	"os"
	"strconv"
	"testing"
)

const (
	expectAdd     = "ADD"
	expectString  = "STRING"
	expectBool    = "true"
	expectInt     = "100"
	expectInt64   = "100"
	expectFloat64 = "100.100"
	expectUint    = "100"
	expectUint64  = "100"
	expectDefault = "100"
)

var (
	expectedInt     = 100
	expectedInt64   = int64(100)
	expectedFloat64 = float64(100.100)
	expectedUint    = uint(100)
	expectedUint64  = uint64(100)
)

var result map[string]interface{}
var required, _ = strconv.ParseBool(os.Getenv("REQUIRED"))

func setup() {
	manual, _ := strconv.ParseBool(os.Getenv("MANUAL"))

	if !manual {
		_ = flag.Set("ADD", expectAdd)
		_ = flag.Set("STRING", expectString)
		_ = flag.Set("BOOL", expectBool)
		_ = flag.Set("INT", expectInt)
		_ = flag.Set("INT64", expectInt64)
		_ = flag.Set("FLOAT64", expectFloat64)
		_ = flag.Set("UINT", expectUint)
		_ = os.Setenv("UINT64", expectUint64)
	}
}

func TestMain(m *testing.M) {
	Add("DEFAULT_FIELD", expectDefault, "The default Parameter", false)
	Add("ADD", "default", "The string Parameter", false)
	AddString("STRING", "default", "The string Parameter", false)
	AddBool("BOOL", false, "The bool Parameter from env", false)
	AddInt("INT", 50, "The int Parameter", false)
	AddInt64("INT64", 50, "The int64 Parameter", false)
	AddFloat64("FLOAT64", 50.50, "The float64 Parameter", false)
	AddUint("UINT", 50, "The uint Parameter", false)
	AddUint64("UINT64", 50, "The uint64 Parameter", false)

	if required {
		AddBool("REQUIRED_FIELD", false, "The required Parameter", true)
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

func TestParameters_AddString(t *testing.T) {
	require.Equal(t, expectString, result["STRING"])
}

func TestParameters_AddInt(t *testing.T) {
	require.Equal(t, expectedInt, result["INT"])
}

func TestParameters_AddInt64(t *testing.T) {
	require.Equal(t, expectedInt64, result["INT64"])
}

func TestParameters_AddFloat64(t *testing.T) {
	require.Equal(t, expectedFloat64, result["FLOAT64"])
}

func TestParameters_AddBool(t *testing.T) {
	require.Equal(t, true, result["BOOL"])
}

func TestParameters_Add(t *testing.T) {
	require.Equal(t, expectAdd, result["ADD"])
}

func TestParameters_AddUint(t *testing.T) {
	require.Equal(t, expectedUint, result["UINT"])
}

func TestParameters_AddUint64(t *testing.T) {
	require.Equal(t, expectedUint64, result["UINT64"])
}

func TestParameters_AddDefaultField(t *testing.T) {
	require.Equal(t, expectDefault, result["DEFAULT_FIELD"])
}
