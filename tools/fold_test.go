package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFoldl(t *testing.T) {
	assert.Equal(t, 15, Foldl(func(i1, i2 interface{}) interface{} {
		return i1.(int) + i2.(int)
	}, 0, []interface{}{1, 2, 3, 4, 5}))
}

func TestFoldr(t *testing.T) {
	// 5 - (4 - (3 - (2 - (1 - 0))))
	// 3
	assert.Equal(t, 3, Foldr(func(i1, i2 interface{}) interface{} {
		return i1.(int) - i2.(int)
	}, 0, []interface{}{5, 4, 3, 2, 1}))
}
