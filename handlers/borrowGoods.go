package handler

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"

	"goodsman/db"
	"goodsman/model"
	"goodsman/response"
)

func BorrowGoods(c *gin.Context) {
	var req model.BorrowGoodsRequest

	if err := c.Bind(&req); err != nil {
		logrus.Error(err)
		response.Error(c, response.PARAMS_ERROR)
		return
	}

	approved, err := BorrowingAuthVerification(c, req.GoodsId)
	if err != nil {
		logrus.Error("error happened in database & ", err.Error())
		response.Error(c, response.DATABASE_ERROR)
		return
	}
	if !approved {
		logrus.Info("Insufficient employee auth")
		response.Error(c, response.AUTH_ERROR)
		return
	}

	goodsState, restGoodsNum, err := GoodsStockVerification(req.GoodsId, req.GoodsNum)
	if err != nil {
		logrus.Error("error happened in database & ", err.Error())
		response.Error(c, response.DATABASE_ERROR)
		return
	}
	if goodsState != 1 {
		logrus.Info("Goods is abnormal & state ", goodsState)
		response.Error(c, response.STOCK_ERROR)
		return
	}
	if restGoodsNum < 0 {
		logrus.Info("Insufficient quantity of goods")
		response.Error(c, response.STOCK_ERROR)
		return
	}

	err = UpdateBorrowGoods(req.GoodsId, restGoodsNum, req.EmployeeId)
	if err != nil {
		logrus.Error("error happened in database & ", err.Error())
		response.Error(c, response.DATABASE_ERROR)
		return
	}

	response.Success(c, response.SUCCESS)
}

func UpdateBorrowGoods(goodsId string, restGoodsNum int, employeeId string) error {

	var goods model.Goods
	ctx := context.TODO()
	filter := bson.M{"goods_id": goodsId}
	err := db.MongoDB.GoodsColl.FindOne(ctx, filter).Decode(&goods)
	if err != nil {
		return err
	}

	// TODO: remove
	ctx = context.TODO()
	filter = bson.M{"goods_id": goodsId}
	update := bson.M{"$set": bson.M{"number": restGoodsNum}}

	if goods.Goods_msg.Consumables == 1 {
		update = bson.M{"$set": bson.M{"number": restGoodsNum}}
	} else if goods.Goods_msg.Consumables == 0 {
		update = bson.M{"$set": bson.M{"number": 0, "owner": employeeId, "state": 0}}
	}

	updateResult, err := db.MongoDB.GoodsColl.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	logrus.Info(updateResult)
	return nil
}
