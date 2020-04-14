package flags

import "strconv"

type Uint64Value struct {
	set   bool
	value uint64
}

func (i *Uint64Value) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, strconv.IntSize)
	if err != nil {
		err = numError(err)
	}
	i.value = v
	i.set = true
	return err
}

func (i *Uint64Value) Get() interface{} { return i.value }

func (i *Uint64Value) String() string { return strconv.FormatUint(i.value, 10) }

func (i *Uint64Value) IsSet() bool { return i.set }
