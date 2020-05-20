// Package flags is a wrapper for the default flag package. All types and functions allow checking if a flag was set.
//
// Example:
//
// flags.Bool("local", false, "")
//
// flag.Parse()
//
// flag.Lookup("local").Value.(*flags.IntType).IsSet()
package flags

import (
	"errors"
	"flag"
	"strconv"

	"github.com/barchart/common-go/pkg/configuration/database"
)

// Bool defines a bool flag with specified name, default value, and usage string.
func Bool(name string, value bool, usage string) {
	v := BoolValue{
		set:   false,
		value: value,
	}
	flag.Var(&v, name, usage)
}

// Bool defines a database flag with specified name, default value, and usage string.
func Database(name string, value database.Database, usage string) {
	v := DatabaseValue{
		set:   false,
		value: value,
	}
	flag.Var(&v, name, usage)
}

// Float64 defines a float64 flag with specified name, default value, and usage string.
func Float64(name string, value float64, usage string) {
	v := Float64Value{
		set:   false,
		value: value,
	}

	flag.Var(&v, name, usage)
}

// Int defines a int flag with specified name, default value, and usage string.
func Int(name string, value int, usage string) {
	v := IntValue{
		set:   false,
		value: value,
	}

	flag.Var(&v, name, usage)
}

// Int64 defines a int64 flag with specified name, default value, and usage string.
func Int64(name string, value int64, usage string) {
	v := Int64Value{
		set:   false,
		value: value,
	}

	flag.Var(&v, name, usage)
}

// String defines a string flag with specified name, default value, and usage string.
func String(name string, value string, usage string) {
	v := StringValue{
		set:   false,
		value: value,
	}

	flag.Var(&v, name, usage)
}

// Uint defines a uint flag with specified name, default value, and usage string.
func Uint(name string, value uint, usage string) {
	v := UintValue{
		set:   false,
		value: value,
	}

	flag.Var(&v, name, usage)
}

// Uint64 defines a uint64 flag with specified name, default value, and usage string.
func Uint64(name string, value uint64, usage string) {
	v := Uint64Value{
		set:   false,
		value: value,
	}

	flag.Var(&v, name, usage)
}

// errParse is returned by Set if a flag's value fails to parse, such as with an invalid integer for Int.
// It then gets wrapped through failf to provide more information.
var errParse = errors.New("parse error")

// errRange is returned by Set if a flag's value is out of range.
// It then gets wrapped through failf to provide more information.
var errRange = errors.New("value out of range")

func numError(err error) error {
	ne, ok := err.(*strconv.NumError)
	if !ok {
		return err
	}
	if ne.Err == strconv.ErrSyntax {
		return errParse
	}
	if ne.Err == strconv.ErrRange {
		return errRange
	}
	return err
}
