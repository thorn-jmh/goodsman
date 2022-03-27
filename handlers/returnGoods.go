package handler

import (
	"context"
	"goodsman/db"
	"goodsman/model"
	"goodsman/response"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ReturnAllGoods(c *gin.Context) {
	var req model.ReturnAllGoodsRequest

	if err := c.Bind(&req); err != nil {
		logrus.Error(err)
		response.Error(c, response.PARAMS_ERROR)
		return
	}

	records, err := getRecordsByEmpIdAndGoodsId(req.EmployeeId, req.GoodsId)
	record := records[0]

	err = UpdateChangeGoodsState(req.GoodsId, 1, -record.Del_num)
	if err != nil {
		logrus.Error("error happened in database & ", err.Error())
		response.Error(c, response.DATABASE_ERROR)
		return
	}

	ctx := context.TODO()

	newRecord := &model.Records{
		Employee_id: req.EmployeeId,
		Goods_id:    req.GoodsId,
		Date:        time.Now().Unix(),
		State:       1,
		Del_num:     -record.Del_num, // It should be positive
	}
	returnRecord, err := db.MongoDB.RecordsColl.InsertOne(ctx, newRecord)
	_ = returnRecord
	response.Success(c, response.SUCCESS)
}

func ReturnGoods(c *gin.Context) {
	var req model.ReturnGoodsRequest

	if err := c.Bind(&req); err != nil {
		logrus.Error(err)
		response.Error(c, response.PARAMS_ERROR)
		return
	}

	err := UpdateChangeGoodsState(req.GoodsId, 1, req.DelNum)
	if err != nil {
		logrus.Error("error happened in database & ", err.Error())
		response.Error(c, response.DATABASE_ERROR)
		return
	}

	ctx := context.TODO()

	newRecord := &model.Records{
		Employee_id: req.EmployeeId,
		Goods_id:    req.GoodsId,
		Date:        time.Now().Unix(),
		State:       1,
		Del_num:     req.DelNum, // It should be positive
	}
	returnRecord, err := db.MongoDB.RecordsColl.InsertOne(ctx, newRecord)
	_ = returnRecord
	response.Success(c, response.SUCCESS)
}
