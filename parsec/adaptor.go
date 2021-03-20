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
