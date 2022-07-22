package domain

const (
	ItemCategoryDiscount = iota + 1
	ItemCategoryRebate
)

type Items []*Item

// 基类商品
type Item struct {
	ID       int
	Category int
	Title    string
	Stock    int

	PriceMarket int
}

func (dom *Item) OfInstance() interface{} {
	switch dom.Category {
	case ItemCategoryRebate:
		return dom.OfInstanceRebate()
	case ItemCategoryDiscount:
		return dom.OfInstanceDiscount()
	}

	return nil
}

func (dom *Item) OfInstanceRebate() *ItemRebate {
	return &ItemRebate{
		Item: dom,
	}
}

func (dom *Item) OfInstanceDiscount() *ItemDiscount {
	return &ItemDiscount{
		Item: dom,
	}
}

func (dom *Item) StockEnough(number int) bool {
	if dom.Stock >= number {
		return true
	}

	return false
}
