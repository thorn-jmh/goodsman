package handler

import (
	"context"
	"goodsman/db"
	"goodsman/model"
	"goodsman/response"
	"goodsman/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func AddNewGoods(c *gin.Context) {
	var req model.AddNewGoodsRequest

	//TODO: Auth Check?
	if err := c.Bind(&req); err != nil {
		logrus.Error(err)
		response.Error(c, response.PARAMS_ERROR)
		return
	}

	uid, err := CreateNewGoods(req)
	if err != nil {
		logrus.Error("error happened in database & ", err.Error())
		response.Error(c, response.DATABASE_ERROR)
		return
	}

	resp := model.AddNewGoodsResp{GoodsId: uid}
	response.Success(c, resp)
}

func CreateNewGoods(req model.AddNewGoodsRequest) (string, error) {
	newUID, _ := utils.GetUID()

	newGoods := &model.Goods{
		Goods_id:   newUID,
		Goods_auth: req.GoodsAuth,
		Number:     req.Number,
		State:      req.State,
		Owner:      req.Owner,
		Goods_msg:  req.GoodsMsg,
	}

	ctx := context.TODO()
	createResult, err := db.MongoDB.GoodsColl.InsertOne(ctx, newGoods)

	if err != nil {
		return newUID, err
	}
	logrus.Info(createResult)
	return newUID, nil
}
