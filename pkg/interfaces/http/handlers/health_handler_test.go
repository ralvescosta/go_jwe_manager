package handlers

import (
	"net/http"
	"testing"

	httpServer "jwemanager/pkg/infra/http_server"
	"jwemanager/pkg/infra/logger"
	"jwemanager/pkg/interfaces/http/factories"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

func Test_Check(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := makeHealthHandlerSut()

		req := httpServer.HttpRequest{}

		res := sut.handler.Check(req)

		assert.Equal(t, http.StatusOK, res.StatusCode)
	})

	t.Run("should return badRequest if some error occur in redis", func(t *testing.T) {
		sut := makeHealthHandlerSut()
		sut.redisClient.Close()
		req := httpServer.HttpRequest{}

		res := sut.handler.Check(req)

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})
}

type healthHandlerSutRtn struct {
	handler     IHealthHandler
	logger      *logger.LoggerSpy
	httpFactory factories.HttpResponseFactory
	redisClient *redis.Client
}

func makeHealthHandlerSut() healthHandlerSutRtn {
	logger := logger.NewLoggerSpy()
	httpFactor := factories.NewHttpResponseFactory()
	miniRedis, _ := miniredis.Run()
	redisClient := redis.NewClient(&redis.Options{
		Addr: miniRedis.Addr(),
	})

	handler := NewHealthHandler(logger, httpFactor, redisClient)

	return healthHandlerSutRtn{handler, logger, httpFactor, redisClient}
}
