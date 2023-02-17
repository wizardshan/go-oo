package bo

const (
	ItemCategoryDiscount = iota + 1
	ItemCategoryTrial
	ItemCategoryRebate
	ItemCategoryVirtual
)

// 价格计算器接口
type ItemPriceCalculator interface {
	Price() int
}

// 返利计算器接口
type ItemRebateCalculator interface {
	Rebate() int
}

// 市场价是否隐藏接口
type ItemPriceMarketHidden interface {
	PriceMarketHidden() bool
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

	stockHandler := &ItemStock{Number: bo.Stock}
	switch bo.Category {
	case ItemCategoryDiscount:
		instance = &ItemDiscount{
			Item:             bo,
			ItemStockHandler: stockHandler,
		}
	case ItemCategoryTrial:
		instance = &ItemTrial{
			Item:             bo,
			ItemStockHandler: stockHandler,
		}
	case ItemCategoryRebate:
		instance = &ItemRebate{
			Item:             bo,
			ItemStockHandler: stockHandler,
		}
	case ItemCategoryVirtual:
		instance = &ItemVirtual{
			Item: bo,
		}
	}

	bo.Instance = instance
}
