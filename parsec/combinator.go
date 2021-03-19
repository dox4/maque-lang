package parsec

import (

	"github.com/dox4/maque-lang/option"
)

func (p Parser) Many() Parser {
	return func(s string) (string, *option.Option) {
		remainder, result := (p)(s)
		var resultSet []interface{} = nil
		for result.IsPresent() {
			resultSet = append(resultSet, result.Get())
			remainder, result = (p)(remainder)
		}
		return remainder, option.OfNilable(resultSet)
	}
}

func (p Parser) Seq(others ...Parser) Parser {
	return func(s string) (string, *option.Option) {
		remainder, result := p(s)
		if result.IsNil() {
			return s, option.OfNil()
		}
		var resultSet []interface{} = []interface{}{result.Get()}
		for _, ap := range others {
			remainder, result = ap(remainder)
			if result.IsNil() {
				return s, option.OfNil()
			} else {
				resultSet = append(resultSet, result.Get())
			}
		}
		return remainder, option.OfValue(resultSet)
	}
}

func (p Parser) Or(alters...Parser) Parser {
	return func(s string) (string, *option.Option) {
		remainder, result := p(s)
		if result.IsPresent() {
			return remainder, result
		}
		for _, ap := range alters {
			remainder, result = ap(s)
			if result.IsPresent() {
				return remainder, result
			}
		}
		return s, option.OfNil()
	}
}
