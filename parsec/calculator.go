package parsec

import (
	"strconv"

	"github.com/dox4/maque-lang/pair"
)

var Digit1to9 Parser = OneOf("123456789")
var Digit0to9 Parser = OneOf("1234567890")

var Integer Parser = Digit1to9.
	Map(SingletonList).
	Seq(Digit0to9.Many()).
	Map(Concat).
	Or(Char('0').Map(SingletonList))
var Frac Parser = Char('.').
	Map(SingletonList).
	Seq(Digit0to9.AtLeast(1)).
	Map(Concat)
var Float Parser = Integer.Seq(Frac.Option([]interface{}{})).Map(Concat)

var Add Parser = Char('+')
var Sub Parser = Char('-')
var Mul Parser = Char('*')
var Div Parser = Char('/')

var VecToStr Mapper = func(i interface{}) interface{} {
	arr := i.([]interface{})
	runes := make([]rune, len(arr))
	for i, v := range arr {
		runes[i] = v.(rune)
	}
	return string(runes)
}

var FloatLiteral = Float.Map(VecToStr)
var FloatValue = FloatLiteral.Map(func(i interface{}) interface{} {
	v, _ := strconv.ParseFloat(i.(string), 64)
	return v
})

type Expr interface {
	execute() float64
}

type BiExpr struct {
	left  Expr
	right Expr
}

type AddExpr struct {
	BiExpr
}

type SubExpr struct {
	BiExpr
}

type MulExpr struct {
	BiExpr
}

type DivExpr struct {
	BiExpr
}

type PrimaryExpr struct {
	value float64
}

func (e *AddExpr) execute() float64 {
	return e.left.execute() + e.right.execute()
}

func (e *SubExpr) execute() float64 {
	return e.left.execute() - e.right.execute()
}

func (e *MulExpr) execute() float64 {
	return e.left.execute() * e.right.execute()
}

func (e *DivExpr) execute() float64 {
	return e.left.execute() / e.right.execute()
}

func (e *PrimaryExpr) execute() float64 {
	return e.value
}

func NewPrimaryExpr(value float64) Expr {
	return &PrimaryExpr{
		value: value,
	}
}

func NewAddExpr(left Expr, right Expr) Expr {
	return &AddExpr{
		BiExpr{
			left:  left,
			right: right,
		},
	}
}

func NewSubExpr(left Expr, right Expr) Expr {
	return &SubExpr{
		BiExpr{
			left:  left,
			right: right,
		},
	}
}

func NewMulExpr(left Expr, right Expr) Expr {
	return &MulExpr{
		BiExpr{
			left:  left,
			right: right,
		},
	}
}

func NewDivExpr(left Expr, right Expr) Expr {
	return &DivExpr{
		BiExpr{
			left:  left,
			right: right,
		},
	}
}

var OptionSpace Parser = OneOf(" \t").Many()

var Primary Parser = FloatValue.PackedBy(OptionSpace, OptionSpace).Map(func(i interface{}) interface{} {
	return NewPrimaryExpr(i.(float64))
})

func BiopBuilder(op int32) func(left Expr, right Expr) Expr {
	switch op {
	case '*':
		return NewMulExpr
	case '/':
		return NewDivExpr
	case '+':
		return NewAddExpr
	case '-':
		return NewSubExpr
	default:
		panic("unrecognized operator: " + string(op))
	}
}

var GenericBiopBuilder = func(i1, i2 interface{}) interface{} {
	p := i2.(*pair.Pair)
	return BiopBuilder(p.First().(int32))(i1.(Expr), p.Second().(Expr))
}

var Term Parser = Primary.ChainLeft(Char('*').Or(Char('/')).And(Primary).Many(), GenericBiopBuilder)

var Expression Parser = Term.ChainLeft(Char('+').Or(Char('-')).And(Term).Many(), GenericBiopBuilder)
