package flags

type StringValue struct {
	set   bool
	value string
}

func (s *StringValue) Set(val string) error {
	s.value = val
	s.set = true

	return nil
}

func (s StringValue) Get() interface{} { return s.value }

func (s StringValue) String() string { return s.value }

func (s StringValue) IsSet() bool { return s.set }
