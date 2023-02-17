package bo

// 试用商品
type ItemTrial struct {
	*Item
	ItemStockHandler
}

// 实现价格计算器接口
func (bo *ItemTrial) Price() int {
	return 0
}
