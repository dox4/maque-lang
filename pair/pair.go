package pair

type Pair struct {
	first interface{}
	second interface{}
}

func NewPair(first interface{}, second interface{}) *Pair {
	return &Pair{
		first,
		second,
	}
}

func (p *Pair) First() interface{} {
	return p.first
}

func (p *Pair) Second() interface{} {
	return p.second
}