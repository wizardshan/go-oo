package domain

// 折扣商品
type ItemDiscount struct {
	*Item

}

// 实现价格计算器接口
func (dom *ItemDiscount) Price() int {
	return dom.PriceMarket / 2
}