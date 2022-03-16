package handler

import (
	"context"
	"goodsman/db"
	"goodsman/model"
	"goodsman/response"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func GetGoodsMsg(c *gin.Context) {

	goodsID := c.DefaultQuery("goods_id", "err")
	if goodsID == "err" {
		logrus.Error("can't parse goods_id")
		response.Error(c, response.PARAMS_ERROR)
		return
	}
	var goods model.Goods
	ctx := context.TODO()
	filter := bson.D{{"goods_id", goodsID}}
	err := db.MongoDB.GoodsColl.FindOne(ctx, filter).Decode(&goods)
	if err != nil {
		logrus.Error("error happened in database & ", err.Error())
		response.Error(c, response.DATABASE_ERROR)
		return
	}
	resp := model.GetGoodsMsgResp{
		Goods_auth: goods.Goods_auth,
		Number:     goods.Number,
		State:      goods.State,
		Owner:      goods.Owner,
		Goods_msg:  goods.Goods_msg,
	}
	response.Success(c, resp)
}
