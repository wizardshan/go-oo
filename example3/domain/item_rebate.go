package domain

// 返利商品
type ItemRebate struct {
	*Item
}

// 实现价格计算器接口
func (dom *ItemRebate) Price() int {
	return dom.PriceMarket
}