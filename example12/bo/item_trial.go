package bo

// 试用商品
type ItemTrial struct {
	*Item
}

// 实现价格计算器接口
func (bo *ItemTrial) Price() int {
	return 0
}

// 实现VIP价格计算器接口
func (bo *ItemTrial) PriceVIP() int {
	return 0
}
