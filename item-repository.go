package go_oo

import (
	"go-oo/example1/bo"
)

type ItemRepository interface {
	Get() *bo.Item
	All() bo.Items
}
