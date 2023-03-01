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

// VIP价格计算器接口
type ItemPriceVIPCalculator interface {
	PriceVIP() int
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
	PriceVIP    int
	Rebate      int

	Instance interface{}

	Shop *Shop
}

func (bo *Item) OfInstance() {

	var instance interface{}

	switch bo.Category {
	case ItemCategoryDiscount:
		instance = &ItemDiscount{
			Item: bo,
		}
	case ItemCategoryRebate:
		instance = &ItemRebate{
			Item: bo,
		}
	}

	bo.Instance = instance
}
