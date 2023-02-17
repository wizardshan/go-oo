package bo

const (
	ItemCategoryDiscount = iota + 1
	ItemCategoryTrial
	ItemCategoryRebate
)

// 价格计算器接口
type ItemPriceCalculator interface {
	Price() int
}

// 返利计算器接口
type ItemRebateCalculator interface {
	Rebate() int
}

type Items []*Item

type Item struct {
	ID          int
	Category    int
	Title       string
	Stock       int
	PriceMarket int
	Price       int
	Rebate      int

	Instance interface{}
}

func (bo *Item) OfInstance() {

	var instance interface{}
	switch bo.Category {
	case ItemCategoryDiscount:
		instance = &ItemDiscount{
			Item: bo,
		}
	case ItemCategoryTrial:
		instance = &ItemTrial{
			Item: bo,
		}
	case ItemCategoryRebate:
		instance = &ItemRebate{
			Item: bo,
		}
	}

	bo.Instance = instance
}