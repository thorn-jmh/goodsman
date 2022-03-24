package handler

import (
	"context"
	"fmt"
	"goodsman/db"
	"goodsman/model"
	"goodsman/response"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

var MAX_MONEY = 100.0

func GetUserAuth(c *gin.Context) {
	resp := model.UserAuthResp{}
	empID := c.Query("employee_id")
	resp.Authority = c.GetInt("employee_auth")

	year, month, day := time.Now().Local().Date()
	date, _ := time.Parse("2006-01-02 15:04:05",
		fmt.Sprintf("%d-%02d-%02d 00:00:00", year, month, day))

	ctx := context.TODO()
	matchState := bson.D{{"employee_id", empID}, {"state", 0}, {"date", bson.D{{"$gte", date.Unix()}}}}
	cursor, err := db.MongoDB.RecordsColl.Find(ctx, matchState)
	if err != nil {
		logrus.Error("error happened in aggregation & ", err.Error())
		response.Error(c, response.DATABASE_ERROR)
		return
	}
	var rec []model.Records
	err = cursor.All(ctx, &rec)
	if err != nil {
		logrus.Error("error happened in database & ", err.Error())
		response.Error(c, response.DATABASE_ERROR)
		return
	}
	for i := 0; i < len(rec); i++ {
		goods := model.Goods{}
		filter := bson.D{{"goods_id", rec[i].Goods_id}}
		err = db.MongoDB.GoodsColl.FindOne(ctx, filter).Decode(&goods)
		if err != nil {
			logrus.Error("error happened in database & ", err.Error())
			response.Error(c, response.DATABASE_ERROR)
			return
		}
		resp.Money += goods.Goods_msg.Cost
	}
	resp.Money = float64(MAX_MONEY) - resp.Money
	response.Success(c, resp)
}
