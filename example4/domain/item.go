package domain

const (
	ItemCategoryDiscount = iota + 1
	ItemCategoryTrial
	ItemCategoryRebate
)

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
	if dom.Category == ItemCategoryTrial {
		return 0
	} else if dom.Category == ItemCategoryRebate {
		return dom.PriceMarket
	}

	return dom.PriceMarket / 2
}

// 返利计算函数
func (dom *Item) Rebate() int {
	if dom.Category != ItemCategoryRebate {
		return 0
	}
	return dom.PriceMarket * 5 / 100
}
