package handler

import (
	"context"
	"fmt"
	"goodsman/db"
	"goodsman/feishu"
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

	rec := model.Goods{}
	db.MongoDB.GoodsColl.FindOne(ctx, filter).Decode(&rec)
	rec.Number = delNum
	if err = changeNotify(&rec); err != nil {
		logrus.Error("failed to send notification & ", err.Error())
	}

	return nil
}

//变更物品提醒
func changeNotify(newgoods *model.Goods) error {
	userID := ManagerID
	messages := make([]string, 0)
	messages = append(messages, "物品变更提醒:")
	messages = append(messages,
		fmt.Sprintf("Name: %s Type: %s", newgoods.Goods_msg.Name, newgoods.Goods_msg.Type))
	messages = append(messages,
		fmt.Sprintf("ChangeNum: %d State: %d", newgoods.Number, newgoods.State))

	formMsg := &feishu.TextMsg{}
	formMsg.Content = formMsg.NewMsg(messages).(string)
	err := feishu.SendMessage(userID, "text", formMsg)
	if err != nil {
		return err
	}
	logrus.Info("Notification has been sent to manager")
	return nil
}
