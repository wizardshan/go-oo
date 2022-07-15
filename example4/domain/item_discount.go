package domain

// 折扣商品
type ItemDiscount struct {
	*Item
	*ItemPriceVIP
}

// 实现价格计算器接口
func (dom *ItemDiscount) Price() int {
	return dom.PriceMarket / 2
}

// 实现VIP价格计算器接口
func (dom *ItemDiscount) PriceVIP() *int {
	return dom.ItemPriceVIP.Calculate(dom.Price())
}


