package parsec

import (
	// "strconv"
	"testing"
	"unicode"

	"github.com/dox4/maque-lang/pair"
	"github.com/stretchr/testify/assert"
)

func testComplexParser(t *testing.T, parser Parser, src string, wantedValue interface{}, wantedRemainder string, equal func(interface{}, interface{}) bool) {
	remainder, opt := parser(src)
	assert.Equal(t, wantedRemainder, remainder)
	if wantedValue == nil {
		assert.True(t, opt.IsNil())
	} else {
		assert.True(t, equal(wantedValue, opt.Get()))
	}
}

func testComplexParserFailed(t *testing.T, parser Parser, src string) {
	remainder, opt := parser(src)
	assert.Equal(t, src, remainder)
	assert.True(t, opt.IsNil())
}

func compVector(i1, i2 interface{}) bool {
	vector1, vector2 := i1.([]interface{}), i2.([]interface{})
	if len(vector1) != len(vector2) {
		return false
	}
	for i := 0; i < len(vector1); i++ {
		if vector1[i] != vector2[i] {
			return false
		}
	}
	return true
}

func comp(a, b interface{}) bool {
	return a == b
}

func TestParser_Many(t *testing.T) {
	manyA := Char('a').Many()
	testComplexParser(t, manyA, "aaab", []interface{}{'a', 'a', 'a'}, "b", compVector)
}

func TestParser_Seq(t *testing.T) {
	a, b, c := Char('a'), Char('b'), Char('c')
	abc := a.Seq(b, c)
	testComplexParser(t, abc, "abc", []interface{}{'a', 'b', 'c'}, "", compVector)
	testComplexParserFailed(t, abc, "ab")
	testComplexParserFailed(t, abc, "ac")
	testComplexParserFailed(t, abc, "")
	testComplexParserFailed(t, abc, "b")
}

func TestParser_Or(t *testing.T) {
	a, b, c := Char('a'), Char('b'), Char('c')
	abc := a.Or(b, c)
	testComplexParser(t, abc, "abc", 'a', "bc", comp)
	testComplexParser(t, abc, "bc", 'b', "c", comp)
	testComplexParser(t, abc, "c", 'c', "", comp)
	testComplexParserFailed(t, abc, "dfg")
}

func TestParser_TakeLeft(t *testing.T) {
	a := Char('a')
	ignoreB := a.TakeLeft(Char('b'))
	testComplexParser(t, ignoreB, "ab", 'a', "", comp)
}

func TestParser_PackedBy(t *testing.T) {
	a := Char('a')
	packedA := a.PackedBy(Char('('), Char(')'))
	testComplexParser(t, packedA, "(a)", 'a', "", comp)
	testComplexParserFailed(t, packedA, "(a")
	testComplexParserFailed(t, packedA, "a)")
}

func TestParser_Accumulate(t *testing.T) {
	digit := OneOf("1234567890").Accumulate(func(i interface{}) interface{} {
		return i.(int32) - '0'
	}, func(i1, i2 interface{}) interface{} {
		return i1.(int32)*10 + i2.(int32)
	}, int32(0))
	testComplexParser(t, digit, "123", int32(123), "", comp)

	name := Satisfy(func(ch int32) bool {
		return unicode.IsLetter(ch)
	}).Many().Map(VecToStr)
	testComplexParser(t, name, "name", "name", "", comp)
}

func TestParser_Option(t *testing.T) {
	testComplexParser(t, Char(' ').Option(nil), " a", ' ', "a", comp)
	testComplexParser(t, Char(' ').Option(nil), "a", nil, "a", comp)
	ab := Char('a').Seq(Char('b').Option(""))
	testComplexParser(t, ab, "ab", []interface{}{'a', 'b'}, "", compVector)
	testComplexParser(t, ab, "a", []interface{}{'a', ""}, "", compVector)
}

func TestParser_Skip(t *testing.T) {
	testComplexParser(t, Char(' ').Many().Skip(), "            a", nil, "a", comp)
	testComplexParser(t, Char(' ').Many().Skip(), "a", nil, "a", comp)
}

func TestParser_AtLeast(t *testing.T) {
	a3 := Char('a').AtLeast(3)
	testComplexParser(t, a3, "aaaab", []interface{}{'a', 'a', 'a', 'a'}, "b", compVector)
	testComplexParser(t, a3, "aaab", []interface{}{'a', 'a', 'a'}, "b", compVector)
	testComplexParser(t, a3, "aaa", []interface{}{'a', 'a', 'a'}, "", compVector)
	testComplexParserFailed(t, a3, "aa")
	testComplexParserFailed(t, a3, "ab")
	b0 := Char('b').AtLeast(0)
	testComplexParser(t, b0, "bbba", []interface{}{'b', 'b', 'b'}, "a", compVector)
}

func TestCalculator1(t *testing.T) {
	testComplexParser(t, FloatValue, "123", 123.0, "", comp)
	testComplexParser(t, FloatValue, "0", 0.0, "", comp)
	testComplexParser(t, FloatValue, "123.45", 123.45, "", comp)
	testComplexParser(t, FloatValue, "0.123", 0.123, "", comp)
	testComplexParser(t, FloatValue, "12.0", 12.0, "", comp)
	testComplexParser(t, FloatValue, "02.1", 0.0, "2.1", comp)
}

func compPair(i interface{}, j interface{}) bool {
	p1 := i.(*pair.Pair)
	p2 := j.(*pair.Pair)
	return p1.First() == p2.First() && p1.Second() == p2.Second()
}
func TestParser_And(t *testing.T) {
	ab := Char('a').And(Char('b'))
	testComplexParser(t, ab, "ab", pair.NewPair('a', 'b'), "", compPair)
	testComplexParserFailed(t, ab, "bc")
	testComplexParserFailed(t, ab, "ac")
}


func TestParser_ChainLeft(t *testing.T) {
	digit := Digit1to9.Map(func(i interface{}) interface{} {
		return i.(int32) - '0'
	})
	addSome := Char('+').And(digit).Many()
	expr := digit.ChainLeft(addSome, func(i1, i2 interface{}) interface{} {
		return i1.(int32) + i2.(*pair.Pair).Second().(int32)
	})
	testComplexParser(t, expr, "1+2+3+4", int32(10), "", comp)
	testComplexParser(t, expr, "1", int32(1), "", comp)
}


func TestParser_Expression(t *testing.T) {
	Calculator := Expression.Map(func(i interface{}) interface{} {
		return i.(Expr).execute()
	})
	testComplexParser(t, Calculator, " 1.23 - 3.45", -2.22, "", comp)
	testComplexParser(t, Calculator, "1.0-3", -2.0, "", comp)
	testComplexParser(t, Calculator, "4.5 * 2.0 - 3*2", 3.0, "", comp)
	testComplexParser(t, Calculator, "1", 1.0, "", comp)
}