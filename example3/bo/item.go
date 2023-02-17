package bo

type Items []*Item

type Item struct {
	ID          int
	Category    int
	Title       string
	Stock       int
	PriceMarket int
	Price       int
}

func (bo *Item) PriceCalculate() {
	if bo.Category == 1 { // 折扣商品价格
		bo.Price = bo.PriceMarket / 2
	} else if bo.Category == 2 { // 试用商品价格
		bo.Price = 0
	} else if bo.Category == 3 { // 返利商品价格
		bo.Price = bo.PriceMarket
	}
}
