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
		apiGroup.GET("/users/authority")
		apiGroup.GET("/goods/msg")

		apiGroup.POST("/goods/borrow", handler.BorrowGoods)
		apiGroup.POST("/goods/new")
		apiGroup.POST("/goods/state")
	}
	return r
}
