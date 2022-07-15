package domain

const (
	ItemCategoryRebate = iota + 1
	ItemCategoryDiscount
)

type Items []*Item

type Item struct {
	ID       int
	Category int
	Title    string
	Stock    int

	PriceMarket int
}

func (dom *Item) Price() int {
	if dom.Category == ItemCategoryRebate {
		return dom.PriceMarket // 返回市场价
	} else if dom.Category == ItemCategoryDiscount {
		return dom.PriceMarket / 2 // 返回市场价除以2
	}

	return 0
}

func (dom *Item) Rebate() *int {
	if dom.Category == ItemCategoryRebate {
		rebate := dom.PriceMarket * 5 / 100
		return &rebate
	}

	return nil
}

func (dom *Item) PriceMarketHidden() bool {
	if dom.Category == ItemCategoryRebate {
		return false
	}

	return true
}

func (dom *Item) StockEnough(number int) bool {
	if dom.Stock >= number {
		return true
	}

	return false
}
