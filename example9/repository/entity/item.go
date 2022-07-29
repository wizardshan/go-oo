package entity

import (
	"go-oo/example9/domain"
)

type Items []*Item

type Item struct {
	ID    int
	Category int
	Title string
	Stock int
	PriceMarket int
}

func (ent *Item) Mapping() *domain.Item {
	dom := new(domain.Item)
	dom.ID = ent.ID
	dom.Category = ent.Category
	dom.Title = ent.Title
	dom.PriceMarket = ent.PriceMarket


	switch ent.Category {
	case domain.ItemCategoryDiscount:
		dom.Instance =  &domain.ItemDiscount{
			Item: dom,
			ItemStockHandler: &domain.ItemStock{Number: ent.Stock},
		}
		break
	case domain.ItemCategoryTrial:
		dom.Instance = &domain.ItemTrial{
			Item: dom,
			ItemStockHandler: &domain.ItemStock{Number: ent.Stock},
		}
		break
	case domain.ItemCategoryRebate:
		dom.Instance = &domain.ItemRebate{
			Item: dom,
			ItemStockHandler: &domain.ItemStock{Number: ent.Stock},
		}
		break
	case domain.ItemCategoryVirtual:
		dom.Instance = &domain.ItemVirtual{
			Item: dom,
		}
		break
	}

	return dom
}

func (ent Items) Mapping() domain.Items {
	entItemsLen := len(ent)
	dom := make(domain.Items, entItemsLen)
	if entItemsLen > 0 {
		for entItemsIndex := 0; entItemsIndex < entItemsLen; entItemsIndex++ {
			dom[entItemsIndex] = ent[entItemsIndex].Mapping()
		}
		return dom
	}

	return nil
}
