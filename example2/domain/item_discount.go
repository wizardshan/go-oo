package domain

type ItemDiscount struct {
	*Item
}

func (dom *ItemDiscount) Price() int {
	return dom.PriceMarket / 2
}
