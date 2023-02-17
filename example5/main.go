package main

import (
	"github.com/gin-gonic/gin"
	"go-oo/example5/controller"
)

func main() {

	srv := gin.Default()

	itemCtr := controller.NewItem()
	srv.GET("/items", itemCtr.All)
	srv.GET("/item", itemCtr.Get)

	srv.Run()
}
