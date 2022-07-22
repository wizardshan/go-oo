package domain

// 体验商品
type ItemExperience struct {
	*Item
}

// 实现价格计算器接口
func (dom *ItemExperience) Price() int {
	return 0
}
