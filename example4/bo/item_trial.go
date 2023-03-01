package bo

// 试用商品
type ItemTrial struct {
	*Item
}

// 实现价格计算器接口
func (bo *ItemTrial) PriceCalculate() {
	bo.Price = 0
}
