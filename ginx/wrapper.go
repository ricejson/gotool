package ginx

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Result 统一返回对象
type Result[T any] struct {
	Code int    `json:"code"`
	Data T      `json:"data"`
	Msg  string `json:"msg"`
}

// WrapReq 封装Req序列化过程
func WrapReq[Req any](ctx *gin.Context, bizFunc func(ctx *gin.Context, req Req) Result[any]) {
	// JSON序列化
	var req Req
	if err := ctx.Bind(&req); err != nil {
		return
	}
	res := bizFunc(ctx, req)
	ctx.JSON(http.StatusOK, res)
}
