package rest

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/ismdeep/message-center/app/server/conf"
	"github.com/ismdeep/message-center/pkg/log"
)

var eng *gin.Engine

func init() {
	gin.SetMode(conf.ROOT.Server.Mode)
	eng = gin.New()

	eng.Use(traceLoggerMiddleware())

	eng.Use(gin.BasicAuth(gin.Accounts{
		conf.ROOT.Auth.Username: conf.ROOT.Auth.Password,
	}))

	eng.POST("/api/v1/buckets", BucketCreate)                   // create a bucket
	eng.GET("/api/v1/buckets", BucketList)                      // get bucket list
	eng.POST("/api/v1/buckets/:bucket_id/messages", MessageAdd) // post a message
	eng.GET("/api/v1/buckets/:bucket_id/messages", MessageList) // get messages after id
}

func traceLoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 每个请求生成的请求traceId具有全局唯一性
		ctx.Set(string(log.TraceIDKey), uuid.NewString())
		ctx.Next()
	}
}

func Run() error {
	log.WithContext(context.TODO()).Info("rest server started", zap.String("mode", conf.ROOT.Server.Mode), zap.String("bind", conf.ROOT.Server.Bind))
	return eng.Run(conf.ROOT.Server.Bind)
}
