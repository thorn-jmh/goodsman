package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"goodsman/response"
)

type BorrowGoodsRequests struct {
	EmployeeId string `form:"employee_id" json:"employee_id" binding:"required"`
	GoodsId    string `form:"goods_id" json:"goods_id" binding:"required"`
	GoodsNum   int    `form:"goods_num" json:"goods_num" binding:"required"`
}

func BorrowGoods(c *gin.Context) {
	var req BorrowGoodsRequests

	if err := c.Bind(&req); err != nil {
		logrus.Error(err)
		response.Error(c, response.PARAMS_ERROR)
		return
	}

}
