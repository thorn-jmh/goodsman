package midware

import (
	"goodsman/response"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetAuthInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		empID := c.DefaultQuery("employee_id", "nil")
		if empID == "nil" {
			logrus.Error("params error: cant find employee_id")
			response.Error(c, response.PARAMS_ERROR)
			return
		}
		//与飞书小程序交互

		empAuth := 1
		c.Set("employee_auth", empAuth)
	}
}
