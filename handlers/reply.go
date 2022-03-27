package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"goodsman/config"
	"goodsman/db"
	"goodsman/feishu"
	"goodsman/model"
	"goodsman/response"
	"io/ioutil"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

var KeyWord = config.App.KeyWord

func ReplyCheck(c *gin.Context) {
	eventreq := model.CommonEvent{}
	body, _ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal(body, &eventreq)
	if err != nil {
		logrus.Error("failed to unmarshal feishu response body & ", err.Error())
		response.Error(c, response.FEISHU_ERROR)
		return
	}
	userID := eventreq.Event.Sender.Sender_id.UserID

	if eventreq.Header.EventType == feishu.ReplyEvent {
		text := model.EventContent{}
		err = json.Unmarshal(body, &text)
		if err != nil {
			logrus.Error("failed to unmarshal feishu response body & ", err.Error())
			response.Error(c, response.FEISHU_ERROR)
			return
		}
		txt := text.Event.Message.Content

		if strings.Contains(txt, KeyWord) {
			Reply(c, userID)
		} else {
			Hello(c, userID)
		}

	} else {
		Hello(c, userID)
	}
}

func Reply(c *gin.Context, empID string) {
	messages := make([]string, 0)

	ctx := context.TODO()
	filter := bson.D{{"owner", empID}}
	cursor, err := db.MongoDB.GoodsColl.Find(ctx, filter)
	if err != nil {
		logrus.Error("error happened in database & ", err.Error())
		response.Error(c, response.DATABASE_ERROR)
		return
	}
	records := []model.Goodsmsg{}
	err = cursor.All(ctx, &records)
	if err != nil {
		logrus.Error("error happened in database & ", err.Error())
		response.Error(c, response.DATABASE_ERROR)
		return
	}
	messages = append(messages,
		fmt.Sprintf("您目前借了 %d 件物品", len(records)))

	for i := 0; i < len(records); i++ {
		messages = append(messages,
			fmt.Sprintf("%d. Name:%s Type:%s", i+1, records[i].Name, records[i].Type))
	}

	formMsg := &feishu.TextMsg{}
	formMsg.Content = formMsg.NewMsg(messages).(string)
	err = feishu.SendMessage(empID, "text", formMsg)
	if err != nil {
		logrus.Error("error happened when sending message & ", err.Error())
		response.Error(c, response.FEISHU_ERROR)
		return
	}
	response.Success(c, "succeed to send a reply")
}

func Hello(c *gin.Context, empID string) {
	messages := make([]string, 0)
	messages = append(messages, "欢迎使用物资借用bot！")
	messages = append(messages, "回复\"借用物品\"查询待归还物品")

	formMsg := &feishu.TextMsg{}
	formMsg.Content = formMsg.NewMsg(messages).(string)
	err := feishu.SendMessage(empID, "text", formMsg)
	if err != nil {
		logrus.Error("error happened when sending message & ", err.Error())
		response.Error(c, response.FEISHU_ERROR)
		return
	}
	response.Success(c, "succeed to send a hello")
}
