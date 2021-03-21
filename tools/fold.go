package tools

import (
	"github.com/dox4/maque-lang/parsec"
)

func Foldl(reducer parsec.Biop, base interface{}, list []interface{}) interface{} {
	return foldl(reducer, base, list, 0)
}

func foldl(reducer parsec.Biop, base interface{}, list []interface{}, index int) interface{} {
	if index >= len(list) {
		return base
	}
	return foldl(reducer, reducer(base, list[index]), list, index+1)
}

func Foldr(reducer parsec.Biop, base interface{}, list []interface{}) interface{} {
	return foldr(reducer, base, list, len(list)-1)
}

func foldr(reducer parsec.Biop, base interface{}, list []interface{}, index int) interface{} {
	if index < 0 {
		return base
	}
	return foldr(reducer, reducer(list[index], base), list, index-1)
}
