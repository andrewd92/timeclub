package price_list

func (p PriceList) Marshall() []interface{} {
	result := make([]interface{}, len(p.prices))

	for i, price := range p.prices {
		result[i] = price.Marshal()
	}

	return result
}
