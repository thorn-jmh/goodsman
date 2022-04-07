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

	code := c.DefaultQuery("code", "nil")
	if code == "nil" {
		logrus.Error("failed to parse 'code'")
		response.Error(c, response.PARAMS_ERROR)
		return
	}
	getIDreq.Code = code
	apptoken, err := feishu.AppTokenManager.GetAccessToken()
	if err != nil {
		logrus.Error("failed to get access_token & ", err.Error())
		response.Error(c, response.FEISHU_ERROR)
		return
	}

	content, _ := json.Marshal(getIDreq)
	req, _ := http.NewRequest("POST", url, bytes.NewReader(content))
	resp, err := feishu.CommonClient.Do(req, apptoken)

	if err != nil && err.Error() == "app access token auth failed" {
		// if err != nil {
		apptoken, err = feishu.TenantTokenManager.GetNewAccessToken()
		if err != nil {
			logrus.Error("failed to get access_token & ", err.Error())
			response.Error(c, response.FEISHU_ERROR)
			return
		}
		content, _ = json.Marshal(getIDreq)
		req, _ = http.NewRequest("POST", url, bytes.NewReader(content))
		resp, err = feishu.CommonClient.Do(req, apptoken)
	}

	if err != nil {
		logrus.Error("feishu response error & ", err.Error())
		response.Error(c, response.FEISHU_ERROR)
		return
	}

	fsGetIDresp := model.FSUserIDResp{}
	if err = json.Unmarshal(resp, &fsGetIDresp); err != nil {
		logrus.Error("failed to unmarshal feishu response body & ", err.Error())
		response.Error(c, response.FEISHU_ERROR)
		return
	}

	getIDresp := model.GetUserIDResp{
		EmployeeID:   fsGetIDresp.Data.EmployeeID,
		AccessToken:  fsGetIDresp.Data.AccessToken,
		ExpiresIn:    fsGetIDresp.Data.ExpiresIn,
		RefreshToken: fsGetIDresp.Data.RefreshToken,
	}
	response.Success(c, getIDresp)

}
