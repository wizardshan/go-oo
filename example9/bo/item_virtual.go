package bo

// 虚拟商品
type ItemVirtual struct {
	*Item
}

// 实现价格计算器接口
func (bo *ItemVirtual) Price() int {
	return 0
}
