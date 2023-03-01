package bo

const (
	ShopLevelNormal   = 0
	ShopLevelGold     = 10
	ShopLevelPlatinum = 20
	ShopLevelDiamond  = 30
)

type Shop struct {
	ID    int
	Name  string
	Level int
}
