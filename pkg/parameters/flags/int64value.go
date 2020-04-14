package flags

import "strconv"

type Int64Value struct {
	set   bool
	value int64
}

func (i *Int64Value) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, strconv.IntSize)
	if err != nil {
		err = numError(err)
	}
	i.value = v
	i.set = true

	return err
}

func (i *Int64Value) Get() interface{} { return i.value }

func (i *Int64Value) String() string { return strconv.Itoa(int(i.value)) }

func (i *Int64Value) IsSet() bool { return i.set }
