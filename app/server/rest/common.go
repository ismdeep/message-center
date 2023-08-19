package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func failed(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  msg,
	})
	c.Abort()
}

func success(c *gin.Context, data interface{}) {
	if data == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": data,
	})
	c.Abort()
}
