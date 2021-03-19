package parsec

import (
	"testing"

	"github.com/dox4/maque-lang/option"
	"github.com/stretchr/testify/assert"
)

func testParser(t *testing.T, parser Parser, src string, value *option.Option, remaider string) {
	r, v := parser(src)
	assert.Equal(t, remaider, r)
	assert.Equal(t, value, v)
}

func TestChar(t *testing.T) {
	testParser(t, Char('a'), "abc", option.OfValue('a'), "bc")
	testParser(t, Char('a'), "bc", option.OfNil(), "bc")
	testParser(t, Char('a'), "", option.OfNil(), "")
}

func TestKeyword(t *testing.T) {
	testParser(t, Keyword("hello"), "hello, world", option.OfValue("hello"), ", world")
	testParser(t, Keyword("hello"), "hello", option.OfValue("hello"), "")
	testParser(t, Keyword("hello"), "hell", option.OfNil(), "hell")
}

func TestOneOf(t *testing.T) {
	testParser(t, OneOf("1234567890"), "0123", option.OfValue('0'), "123")
	testParser(t, OneOf("1234567890"), "abc", option.OfNil(), "abc")
	testParser(t, OneOf("1234567890"), "", option.OfNil(), "")
}

func white(ch int32) bool {
	return ch == ' '
}
func TestSatisfy(t *testing.T) {
	testParser(t, Satisfy(white), " a", option.OfValue(' '), "a")
	testParser(t, Satisfy(white), "a", option.OfNil(), "a")
}
