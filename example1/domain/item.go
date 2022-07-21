package domain

type Items []*Item

type Item struct {
	ID          int
	Title       string
	Stock       int
	PriceMarket int
}

// 价格计算函数
func (dom *Item) Price() int {
	return dom.PriceMarket / 2
}

func (dom *Item) StockEnough(number int) bool {
	if dom.Stock >= number {
		return true
	}

	return false
}
