package flags

import "strconv"

type Float64Value struct {
	set   bool
	value float64
}

func (f *Float64Value) Set(s string) error {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		err = numError(err)
	}
	f.value = v
	f.set = true

	return err
}

func (f *Float64Value) Get() interface{} { return f.value }

func (f *Float64Value) String() string { return strconv.FormatFloat(f.value, 'g', -1, 64) }

func (f *Float64Value) IsSet() bool { return f.set }
