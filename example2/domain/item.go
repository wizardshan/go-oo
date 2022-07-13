package domain

type ItemPriceCalculator interface {
	Price() int
}

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
