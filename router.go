package main

import (
	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping")
	apiGroup := r.Group("/api")
	{
		apiGroup.GET("/users/authority")
		apiGroup.GET("goods/msg")

		apiGroup.POST("goods/borrow")
		apiGroup.POST("/goods/new")
		apiGroup.POST("/goods/state")
	}
	return r
}
