package bo

// 返利商品
type ItemRebate struct {
	*Item
}

// 实现价格计算器接口
func (bo *ItemRebate) Price() int {
	return bo.PriceMarket
}

// 实现VIP价格计算器接口
func (bo *ItemRebate) PriceVIP() int {
	return bo.Price() - 2
}

// 返利计算函数
func (bo *ItemRebate) Rebate() int {
	return bo.PriceMarket * 5 / 100
}

func (bo *ItemRebate) PriceMarketHidden() bool {
	return true
}