package syntax

import (
	"testing"

	"github.com/dox4/maque-lang/parsec"
	"github.com/stretchr/testify/assert"
)

func assertSimpleValue(t *testing.T, expected Expression, actual Expression) {
	assert.Equal(t, expected.Type(), actual.Type())
	assert.Equal(t, expected.Eval(), actual.Eval())
}

func testLiteralSuccessHelper(t *testing.T, parser parsec.Parser, src string, expectedValue Expression) {
	remainder, result := parser(src)
	assert.Equal(t, "", remainder)
	assert.True(t, result.IsPresent())
	assertSimpleValue(t, expectedValue, result.Get().(Expression))
}

func TestLiteral(t *testing.T) {
	testLiteralSuccessHelper(t, NumberLiteral, "+ +++ +123", Value(123.0).(Expression))
	testLiteralSuccessHelper(t, NumberLiteral, "-- - -+-123", Value(-123.0).(Expression))
	testLiteralSuccessHelper(t, BoolLiteral, "false", Value(false).(Expression))
	testLiteralSuccessHelper(t, BoolLiteral, "true", Value(true).(Expression))
	testLiteralSuccessHelper(t, NullLiteral, "null", Value(nil).(Expression))
}
