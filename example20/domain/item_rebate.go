package domain

// 返利商品
type ItemRebate struct {
	*Item
}

// 实现价格计算器接口
func (dom *ItemRebate) Price() int {
	return dom.PriceMarket
}

// 实现返利计算器接口 封装变化
func (dom *ItemRebate) Rebate() *int {
	rebate := dom.PriceMarket * 5 / 100
	return &rebate
}

// 实现市场价显示接口 封装变化
func (dom *ItemRebate) PriceMarketHidden() bool {
	return true
}
