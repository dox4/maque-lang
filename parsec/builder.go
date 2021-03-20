package parsec

import (
	"strings"

	"github.com/dox4/maque-lang/option"
)

type Parser func(string) (string, *option.Option)

func Char(ch int) Parser {
	return func(src string) (string, *option.Option) {
		if len(src) == 0 || int(src[0]) != ch {
			return src, option.OfNilable(nil)
		}
		return src[1:], option.OfValue(int32(src[0]))
	}
}

func Keyword(kw string) Parser {
	return func(s string) (string, *option.Option) {
		if len(s) < len(kw) || !strings.HasPrefix(s, kw) {
			return s, option.OfNil()
		}
		return s[len(kw):], option.OfValue(kw)
	}
}

func OneOf(set string) Parser {
	return func(s string) (string, *option.Option) {
		if len(s) == 0 || !strings.ContainsRune(set, rune(s[0])) {
			return s, option.OfNil()
		}
		return s[1:], option.OfValue(int32(s[0]))
	}
}

func Satisfy(cond func(int32) bool) Parser {
	return func(s string) (string, *option.Option) {
		if len(s) == 0 || !cond(int32(s[0])) {
			return s, option.OfNil()
		}
		return s[1:], option.OfValue(int32(s[0]))
	}
}
