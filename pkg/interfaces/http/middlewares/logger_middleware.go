package middlewares

import (
	"fmt"
	"io/ioutil"
	"jwemanager/pkg/app/interfaces"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap/zapcore"
)

type LoggerMiddleware struct {
	logger interfaces.ILogger
}

func (pst LoggerMiddleware) Execute(ctx *gin.Context) {
	startTime := time.Now()
	ctx.Next()
	endTime := time.Now()
	latencyTimeInMileseconds := float64(endTime.Sub(startTime).Nanoseconds() / 1000)

	requestBody, _ := ioutil.ReadAll(ctx.Request.Body)
	responseBody, _ := ioutil.ReadAll(ctx.Request.Response.Body)

	pst.logger.Info("[HTTP Request]",
		zapcore.Field{
			Key:    "method",
			Type:   zapcore.StringType,
			String: ctx.Request.Method,
		},
		zapcore.Field{
			Key:     "statusCode",
			Type:    zapcore.Int64Type,
			Integer: int64(ctx.Writer.Status()),
		},
		zapcore.Field{
			Key:    "latencyTime",
			Type:   zapcore.StringType,
			String: fmt.Sprintf("%.2f us", latencyTimeInMileseconds),
		},
		zapcore.Field{
			Key:    "clientIP",
			Type:   zapcore.StringType,
			String: ctx.ClientIP(),
		},
		zapcore.Field{
			Key:    "uri",
			Type:   zapcore.StringType,
			String: ctx.Request.RequestURI,
		},
		zapcore.Field{
			Key:    "headers",
			Type:   zapcore.StringType,
			String: headerToString(ctx.Request.Header),
		},
		zapcore.Field{
			Key:    "request",
			Type:   zapcore.StringType,
			String: string(requestBody),
		},
		zapcore.Field{
			Key:    "response",
			Type:   zapcore.StringType,
			String: string(responseBody),
		},
	)
}

func headerToString(header http.Header) string {
	h := ""
	for k, v := range header {
		h += k + ":" + strings.Join(v, ",") + ";"
	}

	return h
}
