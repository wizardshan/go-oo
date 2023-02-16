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
	if bo.Category == 1 {
		bo.Price = bo.PriceMarket / 2
	} else if bo.Category == 2 {
		bo.Price = 0
	} else if bo.Category == 3 {
		bo.Price = bo.PriceMarket
	}
}
