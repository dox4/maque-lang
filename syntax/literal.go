package syntax

import (
	"github.com/dox4/maque-lang/pair"
	"github.com/dox4/maque-lang/parsec"
)

// number
var NumberLiteral parsec.Parser = parsec.Char('+').Or(parsec.Char('-')).TakeLeft(parsec.Char(' ').Many()).Many().And(parsec.FloatValue).Map(func(i interface{}) interface{} {
	pair := i.(*pair.Pair)
	signs := pair.First().([]interface{})
	value := pair.Second().(float64)
	if len(signs) == 0 {
		return value
	}
	sign := 1
	for _, v := range signs {
		if v == '-' {
			sign = 0 - sign
		}
	}
	if sign == 1 {
		return value
	} else {
		return -value
	}
}).Map(Value)

// bool
var falseLiteral parsec.Parser = parsec.Keyword("false").Map(func(i interface{}) interface{} { return false })
var trueLiteral parsec.Parser = parsec.Keyword("true").Map(func(i interface{}) interface{} { return true })
var BoolLiteral = falseLiteral.Or(trueLiteral).Map(Value)

// null
var NullLiteral parsec.Parser = parsec.Keyword("null").Map(func(_ interface{}) interface{} {
	return Value(nil)
})
