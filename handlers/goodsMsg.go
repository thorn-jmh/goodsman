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

	goodsID := c.DefaultQuery("goods_id", "nil")
	if goodsID == "nil" {
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
		response.Error(c, response.NO_CERTAIN_GOODS)
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

func GoodsStockVerification(goodsId string, goodsNum int) (int, int, error) {
	var goods model.Goods
	ctx := context.TODO()
	filter := bson.D{{"goods_id", goodsId}}
	err := db.MongoDB.GoodsColl.FindOne(ctx, filter).Decode(&goods)
	if err != nil {
		return -1, 0, err
	}
	return goods.State, goods.Number - goodsNum, nil
}
