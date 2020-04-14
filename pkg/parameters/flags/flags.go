package flags

import (
	"errors"
	"flag"
	"strconv"
)

func Bool(name string, value bool, usage string) {
	v := BoolValue{
		set:   false,
		value: value,
	}
	flag.Var(&v, name, usage)
}

func Float64(name string, value float64, usage string) {
	v := Float64Value{
		set:   false,
		value: value,
	}

	flag.Var(&v, name, usage)
}

func Int(name string, value int, usage string) {
	v := IntValue{
		set:   false,
		value: value,
	}

	flag.Var(&v, name, usage)
}

func Int64(name string, value int64, usage string) {
	v := Int64Value{
		set:   false,
		value: value,
	}

	flag.Var(&v, name, usage)
}

func String(name string, value string, usage string) {
	v := StringValue{
		set:   false,
		value: value,
	}

	flag.Var(&v, name, usage)
}

func Uint(name string, value uint, usage string) {
	v := UintValue{
		set:   false,
		value: value,
	}

	flag.Var(&v, name, usage)
}

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
