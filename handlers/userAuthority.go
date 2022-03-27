package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"goodsman/config"
	"goodsman/db"
	"goodsman/feishu"
	"goodsman/model"
	"goodsman/response"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

var MAX_MONEY = config.App.MAX_MONEY

//https://open.feishu.cn/open-apis/contact/v3/users/:user_id
func queryAuth(empID string) (int, error) {
	url := "https://open.feishu.cn/open-apis/contact/v3/users/" + empID
	accessToken, err := feishu.TenantTokenManager.GetAccessToken()
	if err != nil {
		return -1, err
	}

	reqbody := `{"user_id_type":"user_id"}`
	req, _ := http.NewRequest("GET", url, strings.NewReader(reqbody))
	body, err := feishu.CommonClient.Do(req, accessToken)
	if err != nil {
		return -1, err
	}
	result := model.FSUserAuth{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return -1, err
	}

	emp_type := 2 - result.Data.User.EmpType
	if emp_type != 0 && emp_type != 1 {
		return -1, errors.New("cant find this employee in feishu")
	}
	return emp_type, nil
}

func GetUserAuth(c *gin.Context) {
	empID := c.DefaultQuery("employee_id", "nil")
	if empID == "nil" {
		logrus.Error()
		response.Error(c, response.PARAMS_ERROR)
		return
	}
	resp := model.UserAuthResp{}

	year, month, day := time.Now().Local().Date()
	date, _ := time.Parse("2006-01-02 15:04:05",
		fmt.Sprintf("%d-%02d-%02d 00:00:00", year, month, day))

	ctx := context.TODO()
	//FIXME: 时区好像不太对？但是我觉得在一台服务器上跑应该是没问题?
	matchState := bson.D{{"employee_id", empID}, {"state", 0}, {"date", bson.D{{"$gte", date.Unix()}}}}
	cursor, err := db.MongoDB.RecordsColl.Find(ctx, matchState)
	if err != nil {
		logrus.Error("error happened in database & ", err.Error())
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

	resp.Authority, err = queryAuth(empID)
	if err != nil {
		logrus.Error("error happened when querying user authority & ", err.Error())
		response.Error(c, response.FEISHU_ERROR)
		return
	}

	response.Success(c, resp)
}
