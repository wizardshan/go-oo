package main

import (
	"github.com/gin-gonic/gin"
	"go-oo/example11/controller"
)

func main() {

	r := gin.Default()

	itemsCtr := controller.NewItems()
	r.GET("example11/items", itemsCtr.Get)

	itemCtr := controller.NewItem()
	r.GET("example11/item", itemCtr.Get)

	r.Run()
}
