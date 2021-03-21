package parsec

import (
	"github.com/dox4/maque-lang/option"
)

func (p Parser) Map(convertor Mapper) Parser {
	return func(s string) (string, *option.Option) {
		remainder, result := p(s)
		if result.IsPresent() {
			return remainder, option.OfNilable(convertor(result.Get()))
		}
		return s, option.OfNil()
	}
}

type Mapper = func(interface{}) interface{}
type Biop = func(interface{}, interface{}) interface{}

var SingletonList Mapper = func(i interface{}) interface{} {
	return []interface{}{i}
}
var Concat Mapper = func(i interface{}) interface{} {
	a := i.([]interface{})
	a0 := a[0].([]interface{})
	a1 := a[1].([]interface{})
	return append(a0, a1...)
}
