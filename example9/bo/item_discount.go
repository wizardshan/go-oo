package bo

// 折扣商品
type ItemDiscount struct {
	*Item
}

// 实现价格计算器接口
func (bo *ItemDiscount) Price() int {
	return bo.PriceMarket / 2
}
