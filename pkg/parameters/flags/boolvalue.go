package flags

import (
	"strconv"
)

type BoolValue struct {
	set   bool
	value bool
}

func (b *BoolValue) Set(s string) error {
	v, err := strconv.ParseBool(s)
	if err != nil {
		err = errParse
	} else {
		b.set = true
	}
	b.value = v
	return err
}

func (b *BoolValue) Get() interface{} { return b.value }

func (b *BoolValue) String() string { return strconv.FormatBool(b.value) }

func (b *BoolValue) IsBoolFlag() bool { return true }

func (b *BoolValue) IsSet() bool { return b.set }
