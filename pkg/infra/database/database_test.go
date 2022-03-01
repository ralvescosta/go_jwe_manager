package database

import (
	"jwemanager/pkg/infra/logger"
	"os"
	"strings"
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

func Test_Connection(t *testing.T) {
	t.Run("should connect to the redis correctly", func(t *testing.T) {
		sut := makeDatabaseSut()

		client, err := Connection(sut.logger, sut.shotdown)

		assert.IsType(t, &redis.Client{}, client)
		assert.NoError(t, err)
	})

	t.Run("should return error when convert db port env", func(t *testing.T) {
		sut := makeDatabaseSut()
		os.Setenv("REDIS_DB", "")
		client, err := Connection(sut.logger, sut.shotdown)

		assert.Error(t, err)
		assert.Nil(t, client)
	})

	t.Run("should return error when convert seconds to pint env", func(t *testing.T) {
		sut := makeDatabaseSut()
		os.Setenv("REDIS_SECONDS_TO_PING", "")
		client, err := Connection(sut.logger, sut.shotdown)

		assert.Error(t, err)
		assert.Nil(t, client)
	})
}

type databaseSutRtn struct {
	redis    *miniredis.Miniredis
	logger   *logger.LoggerSpy
	shotdown chan bool
}

func makeDatabaseSut() databaseSutRtn {
	miniRedis, _ := miniredis.Run()

	spited := strings.Split(miniRedis.Addr(), ":")
	os.Setenv("REDIS_HOST", spited[0])
	os.Setenv("REDIS_PORT", spited[1])
	os.Setenv("REDIS_DB", "0")
	os.Setenv("REDIS_SECONDS_TO_PING", "10")

	logger := logger.NewLoggerSpy()

	shotdown := make(chan bool)

	return databaseSutRtn{miniRedis, logger, shotdown}
}
