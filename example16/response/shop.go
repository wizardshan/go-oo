package response

import (
	"go-oo/example16/bo"
)

type Shops []*Shop

type Shop struct {
	ID    int    `json:"id"`
	Level int    `json:"level"`
	Name  string `json:"name"`
}

func (resp *Shop) Mapping(boShop *bo.Shop) {
	if boShop == nil {
		return
	}

	resp.ID = boShop.ID
	resp.Level = boShop.Level
	resp.Name = boShop.Name
}
