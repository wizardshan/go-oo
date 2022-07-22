package domain

const (
	ItemCategoryDiscount = iota + 1
	ItemCategoryRebate
)

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
		return 0
	}

	return dom.PriceMarket / 2
}