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
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserAuth(c *gin.Context) {
	//employeeID_midware()
	resp := model.UserAuthResp{}
	empID := c.Query("employee_id")
	resp.Authority = c.GetInt("employee_auth")

	year, month, day := time.Now().Local().Date()
	date, _ := time.Parse("2006-01-02 15:04:05",
		fmt.Sprintf("%d-%02d-%02d 00:00:00", year, month, day))

	ctx := context.TODO()
	matchState1 := bson.D{{"$match", bson.D{{"date", bson.D{{"$gte", date.Unix()}}}}}}
	matchState2 := bson.D{{"$match", bson.D{{"employee_id", empID}}}}
	matchState3 := bson.D{{"$match", bson.D{{"state", 0}}}}
	cursor, err := db.MongoDB.RecordsColl.Aggregate(ctx, mongo.Pipeline{matchState1, matchState2, matchState3})
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
	response.Success(c, resp)
}
