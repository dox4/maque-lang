package parser

type Number struct {
	Value float64
}

func ParseDigit(src string) (string, int) {
	if len(src) == 0 {
		return src, -1
	}
	if src[0] >= '0' && src[0] <= '9' {
		return src[1:], int(src[0])
	}
	return src, -1
}

func ParseNumber(src string) (string, int) {
	var value = 0	
	remaider, digit := ParseDigit(src)
	for ; digit != -1; remaider, digit = ParseDigit(remaider) {
		value = value * 10 + digit - '0'
	}
	return remaider, value
}
