package bo

const (
	UserLevelNormal   = 0
	UserLevelGold     = 10
	UserLevelPlatinum = 20
	UserLevelDiamond  = 30
)

type User struct {
	ID     int
	Mobile string
	Level  int
}

func (bo *User) IsLevelNormal() bool {
	return bo.Level == UserLevelNormal
}

func (bo *User) IsLevelGold() bool {
	return bo.Level == UserLevelGold
}

func (bo *User) IsLevelPlatinum() bool {
	return bo.Level == UserLevelPlatinum
}

func (bo *User) IsLevelDiamond() bool {
	return bo.Level == UserLevelDiamond
}
