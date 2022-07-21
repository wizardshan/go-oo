package domain

const (
	ItemCategoryRebate = iota + 1
	ItemCategoryDiscount
	ItemCategoryExperience
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
	} else if dom.Category == ItemCategoryExperience {
		return 0 // 体验商品返回价格0
	}

	return 0
}

// VIP价格计算
func (dom *Item) PriceVIP() *int {
	// 返利商品和折扣商品有VIP价格
	if dom.Category == ItemCategoryRebate || dom.Category == ItemCategoryDiscount {
		priceVIP := dom.Price() - 1
		return &priceVIP
	}
	return nil
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
