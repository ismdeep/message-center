package rest

import (
	"github.com/gin-gonic/gin"

	"github.com/ismdeep/message-center/app/server/handler"
	"github.com/ismdeep/message-center/internal/model"
)

// BucketCreate create a bucket
// @Summary create a bucket
// @Author l.jiang.1024@gmail.com
// @Description create a bucket
// @Tags Bucket
// @Param req body model.BucketCreateReq true "Bucket Create Request Model"
// @Success 200
// @Router /api/v1/buckets [POST]
func BucketCreate(c *gin.Context) {
	var req model.BucketCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		failed(c, err.Error())
		return
	}

	if err := handler.BucketCreate(c, req.ID, req.Name); err != nil {
		failed(c, err.Error())
		return
	}

	success(c, nil)
}

// BucketList get bucket list
// @Summary get bucket list
// @Author jianglinwei@uniontech.com
// @Author l.jiang.1024@gmail.com
// @Description get bucket list
// @Tags Bucket
// @Param pid path string true "XX ID"
// @Param qid query string true "XX ID"
// @Param req body request. true "JSON数据"
// @Success 200 {object} response.
// @Router /api/v1/buckets [GET]
func BucketList(c *gin.Context) {
	respData, err := handler.BucketList(c)
	if err != nil {
		failed(c, err.Error())
		return
	}

	success(c, respData)
}
