package parsec

func (p *Parser) Many() Parser {
	return func(s string) (string, interface{}) {
		remaider, result := (*p)(s)
		resultSet := []interface{}{}
		for result != nil {
			resultSet = append(resultSet, result)
			remaider, result = (*p)(remaider)
		}
		return remaider, resultSet
	}
}