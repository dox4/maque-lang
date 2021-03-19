package option

type Option struct {
	value interface{}
}

func OfNil() *Option {
	return OfNilable(nil)
}

func OfNilable(v interface{}) *Option {
	return &Option{
		value: v,
	}
}

func OfValue(v interface{}) *Option {
	if v == nil {
		panic("Option.OfValue(v interface{}): v must not be nil.")
	}
	return &Option{
		value: v,
	}
}

func (o *Option) IsNil() bool {
	return o.value == nil
}

func (o *Option) IsPresent() bool {
	return !o.IsNil()
}

func (o *Option) Get() interface{} {
	if o.IsNil() {
		panic("Option.Get(): Option.value is nil.")
	}
	return o.value
}

func (o *Option) OrElse(elseValue interface{}) interface{} {
	if o.value == nil {
		return elseValue
	} else {
		return o.value
	}
}

func (o *Option) OrElseGet(getter func() interface{}) interface{} {
	if o.value == nil {
		return getter()
	} else {
		return o.value
	}
}
