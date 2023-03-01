package entity

import (
	"go-oo/example14/bo"
)

type Users []*User

type User struct {
	ID     int
	Mobile string
	Level  int
}

func (ent *User) Mapping() *bo.User {
	bo := new(bo.User)
	bo.ID = ent.ID
	bo.Mobile = ent.Mobile
	bo.Level = ent.Level

	return bo
}
