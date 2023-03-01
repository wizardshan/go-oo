package service

import "go-oo/example15/bo"

type ItemPrizeCalculator struct {
}

func NewItemPrizeCalculator() *ItemPrizeCalculator {
	return new(ItemPrizeCalculator)
}

func (srv *ItemPrizeCalculator) Handle(item *bo.Item, user *bo.User) {
	if user.IsLevelGold() {
		item.Price -= 1
	} else if user.IsLevelPlatinum() {
		item.Price -= 2
	} else if user.IsLevelDiamond() {
		item.Price -= 3
	}
}
