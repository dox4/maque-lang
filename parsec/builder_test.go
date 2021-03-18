package parsec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testParser(t *testing.T, parser Parser, src string, value interface{}, remaider string) {
	r, v := parser(src)
	assert.Equal(t, remaider, r)
	assert.Equal(t, value, v)
}

func TestChar(t *testing.T) {
	testParser(t, Char('a'), "abc", 'a', "bc")
	testParser(t, Char('a'), "bc", -1, "bc")
	testParser(t, Char('a'), "", -1, "")
}

func TestKeyword(t *testing.T) {
	testParser(t, Keyword("hello"), "hello, world", "hello", ", world")
	testParser(t, Keyword("hello"), "hello", "hello", "")
	testParser(t, Keyword("hello"), "hell", nil, "hell")
}

func TestOneOf(t *testing.T) {
	testParser(t, OneOf("1234567890"), "0123", '0', "123")
	testParser(t, OneOf("1234567890"), "abc", nil, "abc")
	testParser(t, OneOf("1234567890"), "", nil, "")
}

func white(ch int32) bool {
	return ch == ' '
}
func TestSatisfy(t *testing.T) {
	testParser(t, Satisfy(white), " a", ' ', "a")
	testParser(t, Satisfy(white), "a", nil, "a")
}