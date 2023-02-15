package entity

func (ent *Item) Mapping() *domain.Item {
	/**************** mapping start ****************/
	dom := new(domain.Item)
	dom.ID = ent.ID
	dom.Category = ent.Category
	dom.Title = ent.Title
	dom.Stock = ent.Stock
	dom.PriceMarket = ent.PriceMarket
	return dom

	/**************** mapping end  ****************/
}
