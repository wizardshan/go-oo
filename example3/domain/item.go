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

	if dom.Category == ItemCategoryRebate {
		return dom.PriceMarket // 返回市场价原价
	}

	return dom.PriceMarket / 2 // 返回市场价的50%
}

// 返利金额计算函数
func (dom *Item) Rebate() *int {
	if dom.Category == ItemCategoryDiscount {
		return nil
	}

	rebate :=  dom.PriceMarket * 5 / 100
	return &rebate
}

func (dom *Item) PriceMarketHidden() bool {
	if dom.Category == ItemCategoryRebate {
		return true
	}

	return false
}

func (dom *Item) StockEnough(number int) bool {
	if dom.Stock >= number {
		return true
	}

	return false
}