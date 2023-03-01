package bo

// 返利商品
type ItemRebate struct {
	*Item
}

// 实现价格计算器接口
func (bo *ItemRebate) PriceCalculate() {
	bo.Price = bo.PriceMarket
}
