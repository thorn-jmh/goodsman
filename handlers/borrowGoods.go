package handler

import (
	"context"
	"time"

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

	approved, err := BorrowingAuthVerification(c, req.EmployeeId, req.GoodsId)
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

	// TODO: remove delnum or restnum
	err = UpdateBorrowGoods(req.GoodsId, restGoodsNum, req.EmployeeId, -req.GoodsNum)
	if err != nil {
		logrus.Error("error happened in database & ", err.Error())
		response.Error(c, response.DATABASE_ERROR)
		return
	}

	response.Success(c, response.SUCCESS)
}

func UpdateBorrowGoods(goodsId string, restGoodsNum int, employeeId string, delNum int) error {

	var goods model.Goods
	ctx := context.TODO()
	filter := bson.M{"goods_id": goodsId}
	err := db.MongoDB.GoodsColl.FindOne(ctx, filter).Decode(&goods)
	if err != nil {
		return err
	}

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

	record := &model.Records{
		Employee_id: employeeId,
		Goods_id:    goodsId,
		Date:        time.Now().Unix(),
		State:       0,
		Del_num:     delNum, // It should be negative
	}
	borrowRecord, err := db.MongoDB.RecordsColl.InsertOne(ctx, record)

	logrus.Info(updateResult)
	logrus.Info(borrowRecord)
	return nil
}

func BorrowingAuthVerification(c *gin.Context, empId, goodsId string) (bool, error) {
	employeeAuth, err := queryAuth(empId)
	if err != nil {
		return false, err
	}

	var goods model.Goods
	ctx := context.TODO()
	filter := bson.D{{"goods_id", goodsId}}
	err = db.MongoDB.GoodsColl.FindOne(ctx, filter).Decode(&goods)

	if err != nil {
		return false, err
	}

	goodsAuthority := goods.Goods_auth

	if employeeAuth >= goodsAuthority {
		return true, nil
	} else {
		return false, nil
	}

}
