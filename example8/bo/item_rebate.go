package bo

// 返利商品
type ItemRebate struct {
	*Item
}

// 实现价格计算器接口
func (bo *ItemRebate) PriceCalculate() {
	bo.Price = bo.PriceMarket
}

// 返利计算函数
func (bo *ItemRebate) RebateCalculate() {
	bo.Rebate = bo.PriceMarket * 5 / 100
}

func (bo *ItemRebate) PriceMarketHidden() bool {
	return true
}
