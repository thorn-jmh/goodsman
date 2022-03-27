package handler

import (
	"bytes"
	"encoding/json"
	"goodsman/feishu"
	"goodsman/model"
	"goodsman/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//Feishu_API:https://open.feishu.cn/open-apis/mina/v2/tokenLoginValidate
func GetUserId(c *gin.Context) {
	url := "https://open.feishu.cn/open-apis/mina/v2/tokenLoginValidate"
	getIDreq := model.GetUserIDRequest{}

	if err := c.BindJSON(&getIDreq); err != nil {
		logrus.Error("failed to binding params & ", err.Error())
		response.Error(c, response.PARAMS_ERROR)
		return
	}

	apptoken, err := feishu.AppTokenManager.GetAccessToken()
	if err != nil {
		logrus.Error("failed to get access_token & ", err.Error())
		response.Error(c, response.FEISHU_ERROR)
		return
	}

	content, _ := json.Marshal(getIDreq)
	req, _ := http.NewRequest("POST", url, bytes.NewReader(content))
	resp, err := feishu.CommonClient.Do(req, apptoken)
	if err != nil {
		logrus.Error("feishu response error & ", err.Error())
		response.Error(c, response.FEISHU_ERROR)
		return
	}

	getIDresp := model.GetUserIDResp{}
	if err = json.Unmarshal(resp, &getIDresp); err != nil {
		logrus.Error("failed to unmarshal feishu response body & ", err.Error())
		response.Error(c, response.FEISHU_ERROR)
		return
	}

	response.Success(c, getIDresp)

}
