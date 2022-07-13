package domain

type ItemRebate struct {
	*Item
}

func (dom *ItemRebate) Price() int {
	return dom.PriceMarket
}
