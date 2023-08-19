package rest

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/ismdeep/message-center/app/server/handler"
	"github.com/ismdeep/message-center/internal/model"
)

// MessageAdd push a message
// @Summary push a message
// @Author jianglinwei@uniontech.com
// @Author l.jiang.1024@gmail.com
// @Description push a message
// @Tags Message
// @Param Authorization	header string true "Bearer 31a165ba1be6dec616b1f8f3207b4273"
// @Param pid path string true "XX ID"
// @Param qid query string true "XX ID"
// @Param req body request. true "JSON数据"
// @Success 200 {object} response.
// @Router /api/v1/buckets/:bucket_id/messages [POST]
func MessageAdd(c *gin.Context) {
	bucketID := c.Param("bucket_id")
	var req model.MessageAddReq
	if err := c.ShouldBindJSON(&req); err != nil {
		failed(c, "bad request")
		return
	}

	if _, err := handler.MessageAdd(c, bucketID, req.Title, req.Content, req.Tags); err != nil {
		failed(c, err.Error())
		return
	}

	success(c, nil)
}

// MessageList get message list
// @Summary get message list
// @Author jianglinwei@uniontech.com
// @Author l.jiang.1024@gmail.com
// @Description get message list
// @Tags Message
// @Param Authorization	header string true "Bearer 31a165ba1be6dec616b1f8f3207b4273"
// @Param pid path string true "XX ID"
// @Param qid query string true "XX ID"
// @Param req body request. true "JSON数据"
// @Success 200 {object} response.
// @Router /api/v1/buckets/:bucket_id/messages [GET]
func MessageList(c *gin.Context) {
	bucketID := c.Param("bucket_id")
	from, err := strconv.ParseInt(c.DefaultQuery("from", "0"), 10, 64)
	if err != nil {
		failed(c, "bad parameter of from")
	}

	success(c, gin.H{
		"bucketId": bucketID,
		"from":     from,
	})
}
