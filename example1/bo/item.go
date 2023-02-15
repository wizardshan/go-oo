package bo

type Items []*Item

type Item struct {
	ID          int
	Title       string
	Stock       int
	PriceMarket int
	Price int
}