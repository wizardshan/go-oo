package bo

type Items []*Item

type Item struct {
	ID          int
	Title       string
	Stock       int
	PriceMarket int
	Price       int
}

func (bo *Item) PriceCalculate() {
	bo.Price = bo.PriceMarket / 2 // 折扣计算逻辑
}
