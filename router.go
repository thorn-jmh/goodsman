package main

import (
	handler "goodsman/handlers"
	"goodsman/midware"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", handler.Ping)
	apiGroup := r.Group("/api", midware.GetAuthInfo())
	{
		apiGroup.GET("/users/authority", handler.GetUserAuth)
		apiGroup.GET("/goods/msg", handler.GetGoodsMsg)

		apiGroup.POST("/goods/borrow", handler.BorrowGoods)
		apiGroup.POST("/goods/new", handler.AddNewGoods)
		apiGroup.POST("/goods/state", handler.ChangeGoodsState)
	}
	return r
}
