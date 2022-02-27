package repositories

import (
	"context"
	"testing"

	"jwemanager/pkg/app/interfaces"
	valueObjects "jwemanager/pkg/domain/value_objects"
	guidGenerator "jwemanager/pkg/infra/guid_generator"
	"jwemanager/pkg/infra/logger"

	"github.com/alicebob/miniredis/v2"
	"github.com/elliotchance/redismock/v8"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

func Test_KeyRepository(t *testing.T) {
	t.Run("should CreateKey execute correctly", func(t *testing.T) {
		sut := makeKeyRepositorySut(t)

		ctx := context.Background()
		key := valueObjects.Key{}
		timeToExpiration := 0

		sut.guidGen.On("V4").Return("some_guid")
		sut.redisClient.On("Set", ":", key).Return(redis.NewStatusResult("", nil))

		result, err := sut.repo.CreateKey(ctx, key, timeToExpiration)

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

type keyRepositorySutRtn struct {
	repo        interfaces.IKeyRepository
	logger      *logger.LoggerSpy
	guidGen     *guidGenerator.GuidGeneratorSpy
	miniRedis   *miniredis.Miniredis
	redisClient *redismock.ClientMock
}

func makeKeyRepositorySut(t *testing.T) keyRepositorySutRtn {
	logger := logger.NewLoggerSpy()
	guidGen := guidGenerator.NewGUidGeneratorSpy()

	miniRedis := miniredis.RunT(t)
	redisClient := redismock.NewNiceMock(redis.NewClient(&redis.Options{
		Addr: miniRedis.Addr(),
	}))

	repo := NewKeyRepository(logger, guidGen, redisClient)
	return keyRepositorySutRtn{repo, logger, guidGen, miniRedis, redisClient}
}
