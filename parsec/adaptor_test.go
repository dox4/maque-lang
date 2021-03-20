package parsec

import (
	"testing"
)

func TestParser_Map(t *testing.T) {
	digit := OneOf("1234567890").Map(func(i interface{}) interface{} {
		return i.(int32) - '0'
	})
	testComplexParser(t, digit, "123", int32(1), "23", comp)
	testComplexParserFailed(t, digit, "abc")
}
