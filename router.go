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
		apiGroup.GET("/users/id", handler.GetUserId)
		apiGroup.GET("/users/authority", handler.GetUserAuth)
		apiGroup.GET("/users/records/goods_id", handler.GetRecordsByEmployeeIdAndGoodsId)
		apiGroup.GET("/goods/msg", handler.GetGoodsMsg)

		apiGroup.POST("/goods/borrow", handler.BorrowGoods)
		apiGroup.POST("/goods/new", handler.AddNewGoods)
		apiGroup.POST("/goods/state", handler.ChangeGoodsState)
		// apiGroup.POST("/users/return_goods", handler.ReturnGoods) // TODO: It has bug.
		apiGroup.POST("/users/return_goods/all", handler.ReturnAllGoods)

		apiGroup.POST("/users/add/manager", handler.AddManager)
		apiGroup.POST("/users/del/manager", handler.DeleteManager)
	}
	eventGroup := r.Group("/event")
	{
		eventGroup.POST("/received/msg", handler.ReplyCheck)
	}
	return r
}
