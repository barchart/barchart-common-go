package flags

import "strconv"

type IntValue struct {
	set   bool
	value int
}

func (i *IntValue) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, strconv.IntSize)
	if err != nil {
		err = numError(err)
	}
	i.value = int(v)
	i.set = true

	return err
}

func (i *IntValue) Get() interface{} { return i.value }

func (i *IntValue) String() string { return strconv.Itoa(i.value) }

func (i *IntValue) IsSet() bool { return i.set }
