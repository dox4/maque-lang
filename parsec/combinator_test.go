package parsec

import (
	"testing"
	"unicode"

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
	}).Many().Map(func(i interface{}) interface{} {
		arr := i.([]interface{})
		runes := make([]rune, len(arr))
		for i, v := range arr {
			runes[i] = v.(rune)
		}
		return string(runes)
	})
	testComplexParser(t, name, "name", "name", "", comp)
}

func TestParser_Option(t *testing.T) {
	testComplexParser(t, Char(' ').Option(), " a", ' ', "a", comp)
	testComplexParser(t, Char(' ').Option(), "a", nil, "a", comp)
}

func TestParser_Skip(t *testing.T) {
	testComplexParser(t, Char(' ').Many().Skip(), "            a", nil, "a", comp)
	testComplexParser(t, Char(' ').Many().Skip(), "a", nil, "a", comp)
}
