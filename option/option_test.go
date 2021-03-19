package option

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOfNilable(t *testing.T) {
	o1 := OfNilable(nil)
	assert.True(t, o1.IsNil())
	o2 := OfNilable(1)
	assert.False(t, o2.IsNil())
}

func TestOfValue(t *testing.T) {
	o1 := OfValue(1)
	assert.True(t, o1.IsPresent())
}

func TestOption_Get(t *testing.T) {
	o1 := OfValue(1)
	assert.Equal(t, 1, o1.Get())
}

func TestOption_OrElse(t *testing.T) {
	o1 := OfNilable(nil)
	assert.Equal(t, true, o1.IsNil())
	assert.Equal(t, 1, o1.OrElse(1))
	assert.Equal(t, 2, o1.OrElseGet(func() interface{} { return 2 }))
}
