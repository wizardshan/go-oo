package bo

const (
	ShopLevelNormal      = 0
	ShopLevelBronzeMedal = 10
	ShopLevelSilverMedal = 20
	ShopLevelGoldMedal   = 30
)

type Shops []*Shop

type Shop struct {
	ID    int
	Name  string
	Level int
}

func (bo *Shop) IsLevelNormal() bool {
	return bo.Level == ShopLevelNormal
}

func (bo *Shop) IsLevelBronzeMedal() bool {
	return bo.Level == ShopLevelBronzeMedal
}

func (bo *Shop) IsLevelSilverMedal() bool {
	return bo.Level == ShopLevelSilverMedal
}

func (bo *Shop) IsLevelGoldMedal() bool {
	return bo.Level == ShopLevelGoldMedal
}
