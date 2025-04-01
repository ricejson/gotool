package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

// LogMiddlewareBuilder gin扩展日志中间件
type LogMiddlewareBuilder struct {
	logFunc         func(ctx *gin.Context, log AccessLog)
	allowReqBody    bool // 是否打印请求体
	allowRespBody   bool // 是否打印响应体
	maxReqBodySize  int  // 最大允许打印的请求体大小
	maxRespBodySize int  // 最大允许打印的响应体大小
}

func NewLogMiddlewareBuilder(logFunc func(ctx *gin.Context, log AccessLog)) *LogMiddlewareBuilder {
	return &LogMiddlewareBuilder{
		logFunc: logFunc,
	}
}

func (l *LogMiddlewareBuilder) AllowReqBody() *LogMiddlewareBuilder {
	l.allowReqBody = true
	return l
}

func (l *LogMiddlewareBuilder) AllowRespBody() *LogMiddlewareBuilder {
	l.allowRespBody = true
	return l
}

func (l *LogMiddlewareBuilder) MaxReqBodySize(size int) *LogMiddlewareBuilder {
	l.maxReqBodySize = size
	return l
}

func (l *LogMiddlewareBuilder) MaxRespBodySize(size int) *LogMiddlewareBuilder {
	l.maxRespBodySize = size
	return l
}

// Build 在这里需要考虑如何打印请求与响应
func (l *LogMiddlewareBuilder) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var accessLog AccessLog
		// 1. 获取请求方法
		method := ctx.Request.Method
		accessLog.method = method
		// 2. 获取请求URL
		url := ctx.Request.URL.String()
		accessLog.url = url
		// 3. 判断是否需要打印请求体
		if l.allowReqBody && ctx.Request.Body != nil {
			data, _ := ctx.GetRawData()
			// 读取的数据要放回去
			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(data))
			// 如果请求体超过最大允许打印字节数就截取
			if l.maxReqBodySize > 0 && len(data) > l.maxReqBodySize {
				accessLog.reqBody = string(data[:l.maxReqBodySize])
			} else {
				accessLog.reqBody = string(data)
			}
		}
		// 4. 判断是否需要打印响应体
		if l.allowRespBody {
			ctx.Writer = &ResponseWriter{
				ResponseWriter: ctx.Writer,
				al:             &accessLog,
			}
		}
		now := time.Now()
		defer func() {
			accessLog.duration = time.Since(now)
			if l.allowRespBody && l.maxRespBodySize > 0 {
				if len([]byte(accessLog.respBody)) > l.maxRespBodySize {
					accessLog.respBody = string([]byte(accessLog.respBody)[:l.maxRespBodySize])
				}
			}
			l.logFunc(ctx, accessLog)
		}()
		ctx.Next()
	}
}

type ResponseWriter struct {
	gin.ResponseWriter
	al *AccessLog
}

func (r *ResponseWriter) Write(i []byte) (int, error) {
	r.al.respBody = string(i)
	return r.ResponseWriter.Write(i)
}

func (r *ResponseWriter) WriteHeader(statusCode int) {
	r.al.statusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

// AccessLog 日志格式打印对象
type AccessLog struct {
	url        string        // 访问URL
	method     string        // 请求方法
	reqBody    string        // 请求体
	statusCode int           // HTTP响应状态码
	respBody   string        // 响应体
	duration   time.Duration // 请求执行时间
}
