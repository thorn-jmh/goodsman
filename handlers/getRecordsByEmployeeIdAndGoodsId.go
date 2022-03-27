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

func GetRecordsByEmployeeIdAndGoodsId(c *gin.Context) {
	empID := c.Query("employee_id")
	if empID == "" {
		logrus.Error("can't find employee_id")
		response.Error(c, response.PARAMS_ERROR)
		return
	}
	goodsID := c.Query("goods_id")
	if goodsID == "" {
		logrus.Error("can't find goods_id")
		response.Error(c, response.PARAMS_ERROR)
		return
	}

	records, err := getRecordsByEmpIdAndGoodsId(empID, goodsID)
	if err != nil {
		logrus.Error("error happened in database & ", err.Error())
		response.Error(c, response.DATABASE_ERROR)
		return
	}

	response.Success(c, records)
}

func getRecordsByEmpIdAndGoodsId(empID string, goodsID string) ([]model.Records, error) {
	var records []model.Records
	ctx := context.TODO()
	filter := bson.D{{"employee_id", empID}, {"goods_id", goodsID}}
	cursor, err := db.MongoDB.RecordsColl.Find(ctx, filter)
	cursor.All(ctx, &records)

	return records, err
}
