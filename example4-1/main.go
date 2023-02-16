package main

import (
	"github.com/gin-gonic/gin"
	"go-oo/example4/controller"
)

func main() {

	r := gin.Default()

	itemsCtr := controller.NewItems()
	r.GET("example4/items", itemsCtr.Get)

	itemCtr := controller.NewItem()
	r.GET("example4/item", itemCtr.Get)
	r.Run()
}
