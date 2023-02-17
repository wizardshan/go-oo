package main

import (
	"github.com/gin-gonic/gin"
	"go-oo/example6/controller"
)

func main() {

	r := gin.Default()

	itemsCtr := controller.NewItems()
	r.GET("example6/items", itemsCtr.Get)

	itemCtr := controller.NewItem()
	r.GET("example6/item", itemCtr.Get)
	r.Run()
}
