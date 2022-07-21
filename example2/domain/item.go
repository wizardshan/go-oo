package domain

type Items []*Item

type Item struct {
	ID       int
	Category int
	Title    string
	Stock    int

	PriceMarket int
}

// 价格计算函数
func (dom *Item) Price() int {

	if dom.Category == 2 {
		return dom.PriceMarket // 返回市场价原价
	}

	return dom.PriceMarket / 2 // 返回市场价的50%
}

// 返利金额计算函数
func (dom *Item) Rebate() int {
	if dom.Category == 1 {
		return 0
	}

	return dom.PriceMarket * 5 / 100
}