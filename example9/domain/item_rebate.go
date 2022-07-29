package domain

// 返利商品
type ItemRebate struct {
	*Item
	ItemStockHandler
}

// 实现价格计算器接口
func (dom *ItemRebate) Price() int {
	return dom.PriceMarket
}

// 返利计算函数
func (dom *ItemRebate) Rebate() int {
	return dom.PriceMarket * 5 / 100
}

func (dom *ItemRebate) PriceMarketHidden() bool {
	return true
}