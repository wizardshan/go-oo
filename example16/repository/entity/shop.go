package entity

import "go-oo/example16/bo"

type Shops []*Shop

type Shop struct {
	ID    int
	Name  string
	Level int
}

func (ent *Shop) Mapping() *bo.Shop {
	bo := new(bo.Shop)
	bo.ID = ent.ID
	bo.Name = ent.Name
	bo.Level = ent.Level

	return bo
}

func (ents Shops) Mapping() bo.Shops {
	entsLen := len(ents)
	bos := make(bo.Shops, entsLen)
	if entsLen > 0 {
		for entIndex := 0; entIndex < entsLen; entIndex++ {
			bos[entIndex] = ents[entIndex].Mapping()
		}
		return bos
	}

	return nil
}
