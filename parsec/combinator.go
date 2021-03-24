package parsec

import (
	"github.com/dox4/maque-lang/decltype"
	"github.com/dox4/maque-lang/option"
	"github.com/dox4/maque-lang/pair"
	"github.com/dox4/maque-lang/tools"
)

type Biop = decltype.Biop

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
// Many return the parser that parse the pattern like {pattern}*
// which indicator the pattern will match 0 or more times
func (p Parser) Many() Parser {
	return func(s string) (string, *option.Option) {
		remainder, result := (p)(s)
		var resultSet []interface{} = []interface{}{}
		for result.IsPresent() {
			resultSet = append(resultSet, result.Get())
			remainder, result = (p)(remainder)
		}
		return remainder, option.OfValue(resultSet)
	}
}

func (p Parser) AtLeast(atLeast int) Parser {
	if atLeast == 0 {
		return p.Many()
	}
	return func(s string) (string, *option.Option) {
		var resultSet []interface{} = nil
		var remainder string = s
		var result *option.Option
		count := atLeast
		for count > 0 {
			remainder, result = p(remainder)
			if result.IsNil() {
				return s, option.OfNil()
			} else {
				resultSet = append(resultSet, result.Get())
			}
			count = count - 1
		}
		for {
			remainder2, result := p(remainder)
			if result.IsNil() {
				return remainder, option.OfValue(resultSet)
			} else {
				resultSet = append(resultSet, result.Get())
			}
			remainder = remainder2
		}
	}
}

func (p Parser) Seq(others ...Parser) Parser {
	return func(s string) (string, *option.Option) {
		remainder := s
		remainder, result := p(remainder)
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

func (p Parser) Option(defaultValue interface{}) Parser {
	return func(s string) (string, *option.Option) {
		remainder, result := p(s)
		if result.IsNil() {
			return s, option.OfNilable(defaultValue)
		}
		return remainder, result
	}
}

func (p Parser) Skip() Parser {
	return func(s string) (string, *option.Option) {
		remainder, _ := p(s)
		return remainder, option.OfNil()
	}
}

func (p Parser) And(right Parser) Parser {
	return func(s string) (string, *option.Option) {
		remainder, result1 := p(s)
		if result1.IsNil() {
			return s, option.OfNil()
		}
		remainder, result2 := right(remainder)
		if result2.IsNil() {
			return s, option.OfNil()
		}
		return remainder, option.OfValue(pair.NewPair(result1.Get(), result2.Get()))
	}
}

func (p Parser) ChainLeft(tails Parser, reducer Biop) Parser {
	return func(s string) (string, *option.Option) {
		remainder, result := p(s)
		if result.IsNil() {
			return s, option.OfNil()
		}
		remainder2, result2 := tails(remainder)
		if result2.IsNil() {
			return s, option.OfNil()
		}
		return remainder2, option.OfValue(tools.Foldl(reducer, result.Get(), result2.Get().([]interface{})))
	}
}
