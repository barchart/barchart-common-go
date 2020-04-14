package flags

import "strconv"

type UintValue struct {
	set   bool
	value uint
}

func (i *UintValue) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, strconv.IntSize)
	if err != nil {
		err = numError(err)
	}
	i.value = uint(v)
	i.set = true
	return err
}

func (i *UintValue) Get() interface{} { return i.value }

func (i *UintValue) String() string { return strconv.FormatUint(uint64(i.value), 10) }

func (i *UintValue) IsSet() bool { return i.set }
