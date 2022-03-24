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

func ChangeGoodsState(c *gin.Context) {
	var req model.ChangeGoodsStateRequest

	if err := c.Bind(&req); err != nil {
		logrus.Error(err)
		response.Error(c, response.PARAMS_ERROR)
		return
	}

	err := UpdateChangeGoodsState(req.GoodsId, req.GoodsState.NewState, req.GoodsState.DelNum)
	if err != nil {
		logrus.Error("error happened in database & ", err.Error())
		response.Error(c, response.DATABASE_ERROR)
		return
	}

	response.Success(c, response.SUCCESS)
}

func UpdateChangeGoodsState(goodsId string, goodsState int, delNum int) error {
	var goods model.Goods
	ctx := context.TODO()
	filter := bson.M{"goods_id": goodsId}
	err := db.MongoDB.GoodsColl.FindOne(ctx, filter).Decode(&goods)
	if err != nil {
		return err
	}

	ctx = context.TODO()
	filter = bson.M{"goods_id": goodsId}
	update := bson.M{"$set": bson.M{"number": goods.Number + delNum, "state": goodsState}}
	updateResult, err := db.MongoDB.GoodsColl.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}
	logrus.Info(updateResult)
	return nil
}
