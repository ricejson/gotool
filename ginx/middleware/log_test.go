package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ricejson/gotool/logx"
	"net/http"
	"testing"
	"time"
)

func TestLogMiddlewareBuilder_Build(t *testing.T) {
	logger := logx.NewSlogLogger()
	server := gin.Default()
	server.Use(NewLogMiddlewareBuilder(func(ctx *gin.Context, log AccessLog) {
		logger.Info("处理HTTP请求与响应", logx.Field{
			Key:   "accessLog",
			Value: log,
		})
	}).AllowReqBody().
		AllowRespBody().
		MaxReqBodySize(1024).
		MaxRespBodySize(1024).Build())
	server.GET("/test", func(ctx *gin.Context) {
		time.Sleep(time.Second * 1)
		ctx.String(http.StatusOK, "hello world")
	})
	server.Run(":8082")
}
