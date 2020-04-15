//Package flags is a wrapper for the default flag package. All types and functions allow checking if a flag was set.
//Example:
//flags.Bool("local", false, "")
//flag.Parse()
//flag.Lookup("local").Value.(*flags.IntType).IsSet()
package flags

import (
	"errors"
	"flag"
	"strconv"
)

// Bool creates bool flag
func Bool(name string, value bool, usage string) {
	v := BoolValue{
		set:   false,
		value: value,
	}
	flag.Var(&v, name, usage)
}

// Float64 creates float64 flag
func Float64(name string, value float64, usage string) {
	v := Float64Value{
		set:   false,
		value: value,
	}

	flag.Var(&v, name, usage)
}

// Int creates int flag
func Int(name string, value int, usage string) {
	v := IntValue{
		set:   false,
		value: value,
	}

	flag.Var(&v, name, usage)
}

// Int64 creates int64 flag
func Int64(name string, value int64, usage string) {
	v := Int64Value{
		set:   false,
		value: value,
	}

	flag.Var(&v, name, usage)
}

// String creates string flag
func String(name string, value string, usage string) {
	v := StringValue{
		set:   false,
		value: value,
	}

	flag.Var(&v, name, usage)
}

// Uint creates uint flag
func Uint(name string, value uint, usage string) {
	v := UintValue{
		set:   false,
		value: value,
	}

	flag.Var(&v, name, usage)
}

// Uint64 creates uint64 flag
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
