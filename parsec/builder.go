package parsec

import "strings"

type Parser func(string) (string, interface{})

func Char(ch int) Parser {
	return func(src string) (string, interface{}) {
		if len(src) == 0 || int(src[0]) != ch {
			return src, -1
		}
		return src[1:], int32(src[0])
	}
}

func Keyword(kw string) Parser {
	return func(s string) (string, interface{}) {
		if len(s) < len(kw) || !strings.HasPrefix(s, kw) {
			return s, nil
		}
		return s[len(kw):], kw
	}
}

func OneOf(set string) Parser {
	return func(s string) (string, interface{}) {
		if len(s) == 0 || !strings.ContainsRune(set, rune(s[0])) {
			return s, nil
		}
		return s[1:], int32(s[0])
	}
}

func Satisfy(cond func(int32) bool) Parser {
	return func(s string) (string, interface{}) {
		if len(s) == 0 || !cond(int32(s[0])) {
			return s, nil
		}
		return s[1:], int32(s[0])
	}
}

func (p Parser) Accumulate() Parser {
	return nil
}

// var Blank Parser = func(s string) (string, interface{}) {

// }
