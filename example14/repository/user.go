package repository

import (
	"go-oo/example14/bo"
	"go-oo/example14/repository/entity"
)

type User struct {
}

func NewUser() *User {
	return new(User)
}

func (repo *User) Get() *bo.User {

	user := new(entity.User)
	user.ID = 1
	user.Level = 10

	return user.Mapping()
}
