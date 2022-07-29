package domain

// 虚拟商品
type ItemVirtual struct {
	*Item
}

// 实现价格计算器接口
func (dom *ItemVirtual) Price() int {
	return 0
}