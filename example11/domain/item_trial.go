package domain

// 试用商品
type ItemTrial struct {
	*Item
}

// 实现价格计算器接口
func (dom *ItemTrial) Price() int {
	return 0
}
