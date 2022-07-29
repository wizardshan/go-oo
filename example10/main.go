package main

import (
	"github.com/gin-gonic/gin"
	"go-oo/example10/controller"
)

func main() {

	r := gin.Default()

	itemsCtr := controller.NewItems()
	r.GET("example10/items", itemsCtr.Get)

	itemCtr := controller.NewItem()
	r.GET("example10/item", itemCtr.Get)

	r.Run()
}
