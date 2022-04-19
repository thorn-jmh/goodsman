package handler

import (
	"context"
	"fmt"
	"goodsman/db"
	"goodsman/feishu"
	"goodsman/model"
	"goodsman/response"
	"goodsman/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func AddNewGoods(c *gin.Context) {
	var req model.AddNewGoodsRequest
	if err := c.Bind(&req); err != nil {
		logrus.Error(err)
		response.Error(c, response.PARAMS_ERROR)
		return
	}

	if ok := changeAuthCheck(req.EmployeeId); !ok {
		logrus.Error("permission denied")
		response.Error(c, response.AUTH_ERROR)
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
	return
}

func CreateNewGoods(req model.AddNewGoodsRequest) (string, error) {
	newUID, err := utils.GetUID()
	if err != nil {
		logrus.Error("somthing wrong when generating new uid")
		return "", err
	}
	logrus.Info("Generate new UID : ", newUID)

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

	// newNotify(newGoods)
	return newUID, nil
}

//新增物品提醒
func newNotify(newgoods *model.Goods) {
	ManagerID, err := QueryManagers()
	if err != nil {
		logrus.Error("failed to get managers & ", err.Error())
	}
	for i, userID := range ManagerID {
		messages := make([]string, 0)
		messages = append(messages, "新增物品提醒:")
		messages = append(messages,
			fmt.Sprintf("Name: %s Type: %s", newgoods.Goods_msg.Name, newgoods.Goods_msg.Type))
		messages = append(messages,
			fmt.Sprintf("Num: %d Auth: %d", newgoods.Number, newgoods.Goods_auth))

		formMsg := &feishu.TextMsg{}
		formMsg.Content = formMsg.NewMsg(messages).(string)
		err = feishu.SendMessage(userID, "text", formMsg)
		if err != nil {
			logrus.Error("an error happened when sending message & ", err.Error())
		} else {
			logrus.Info(i+1, "notification has been sent to manager")
		}
	}
	logrus.Info("All notification has been sent to manager")
}
