package response

import (
	"github.com/gin-gonic/gin"
)

func response(c *gin.Context, data interface{}, code int, errType int) {
	c.JSON(code, gin.H{
		"code": errType,
		"msg":  msgFlags[errType],
		"data": data,
	})
}

//Succeed，return data
func Success(c *gin.Context, data interface{}) {
	response(c, data, httpStatus[0], 0)
}

//An error happened,
//You can add new errType in /response/define.go
func Error(c *gin.Context, errType int) {

	response(c, nil, httpStatus[errType], errType)
}
