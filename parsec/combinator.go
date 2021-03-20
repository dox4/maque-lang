package parsec

import (
	"github.com/dox4/maque-lang/option"
)

type Mapper = func(interface{}) interface{}
type Biop = func(interface{}, interface{}) interface{}

func (p Parser) Accumulate(mapper Mapper, reducer Biop, base interface{}) Parser {

	return func(s string) (string, *option.Option) {
		for {
			remainder, result := p(s)
			if result.IsNil() {
				return s, option.OfNilable(base)
			}
			temp := mapper(result.Get())
			base = reducer(base, temp)
			s = remainder
		}
	}
}

func (p Parser) Many() Parser {
	return func(s string) (string, *option.Option) {
		remainder, result := (p)(s)
		var resultSet []interface{} = nil
		for result.IsPresent() {
			resultSet = append(resultSet, result.Get())
			remainder, result = (p)(remainder)
		}
		// if resultSet is nil and reunt option.OfNilable(resultSet)
		// option.value will be nil<[]interface{}>
		// while will cause option.IsNil() false
		if resultSet == nil {
			return remainder, option.OfNil()
		}
		return remainder, option.OfValue(resultSet)
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

func (p Parser) Or(alters ...Parser) Parser {
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

func (p Parser) TakeLeft(other Parser) Parser {
	return func(s string) (string, *option.Option) {
		remainder, result := p(s)
		if result.IsPresent() {
			remainder, ignore := other(remainder)
			if ignore.IsPresent() {
				return remainder, result
			}
		}
		return s, option.OfNil()
	}
}

func (p Parser) TakeRight(other Parser) Parser {
	return func(s string) (string, *option.Option) {
		remainder, ignore := p(s)
		if ignore.IsPresent() {
			remainder, result := other(remainder)
			if result.IsPresent() {
				return remainder, result
			}
		}
		return s, option.OfNil()
	}
}

func (p Parser) PackedBy(left, right Parser) Parser {
	return left.TakeRight(p).TakeLeft(right)
}

func (p Parser) Option() Parser {
	return func(s string) (string, *option.Option) {
		remainder, result := p(s)
		if result.IsNil() {
			return s, option.OfNil()
		}
		return remainder, result
	}
}

func (p Parser) Skip() Parser {
	return func(s string) (string, *option.Option) {
		remainder, result := p(s)
		if result.IsNil() {
			return s, option.OfNil()
		}
		return remainder, option.OfNil()
	}
}
