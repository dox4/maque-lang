package syntax

import (
	"unicode"

	"github.com/dox4/maque-lang/parsec"
)

var Expr parsec.Parser

var optionSpace = parsec.OneOf(" \t").Many()

func leftGrop(ch int) parsec.Parser {
	return parsec.Char(ch).TakeLeft(optionSpace)
}

func rightGroup(ch int) parsec.Parser {
	return optionSpace.TakeRight(parsec.Char(ch))
}

var name = parsec.Satisfy(func(i int32) bool {
	return unicode.IsLetter(i)
}).Many().Map(parsec.VecToStr).Map(func(i interface{}) interface{} {
	return NewNameExpr()
})

// ## expression
// 0. primary: literal
var PrimaryExpr = NumberLiteral.Or(BoolLiteral, NullLiteral, GroupExpr)
// 1. group: (expression)
var leftParen = leftGrop('(')
var rightParen = rightGroup(')')
var GroupExpr = Expr.PackedBy(leftParen, rightParen)
// 2. member access: name.member
// var MemberAccess = name.And(parsec.Char('.')).And(name)
// 3. function call: expr()
// 4. math multi: expr * expr, expr / expr
var MultiExpr = Expr.And(parsec.Char('*').Or(parsec.Char('/'))).And(Expr)
// 5. math add: expr + expr, expr - expr
// 6. assignment: name = expr
// 7. function declaration