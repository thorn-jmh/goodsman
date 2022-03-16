package main

import (
	handler "goodsman/handlers"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", handler.Ping)
	apiGroup := r.Group("/api")
	{
		apiGroup.GET("/users/authority", handler.GetUserAuth)
		apiGroup.GET("/goods/msg", handler.GetGoodsMsg)

		apiGroup.POST("/goods/borrow", handler.BorrowGoods)
		apiGroup.POST("/goods/new")
		apiGroup.POST("/goods/state")
	}
	return r
}
