package domain

type Items []*Item

type Item struct {
	ID          int
	Category    int
	Title       string
	Stock       int
	PriceMarket int
}

// 价格计算函数
func (dom *Item) Price() int {
	if dom.Category == 2 {
		return 0
	} else if dom.Category == 3 {
		return dom.PriceMarket
	}

	return dom.PriceMarket / 2
}