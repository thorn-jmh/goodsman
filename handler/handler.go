package handler

import (
	"goodsman/response"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	response.Success(c, "pong!")
}
