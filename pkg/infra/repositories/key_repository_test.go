package repositories

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"jwemanager/pkg/app/interfaces"
	valueObjects "jwemanager/pkg/domain/value_objects"
	"jwemanager/pkg/infra/database/models"
	guidGenerator "jwemanager/pkg/infra/guid_generator"
	"jwemanager/pkg/infra/logger"

	"github.com/alicebob/miniredis/v2"
	"github.com/elliotchance/redismock/v8"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func Test_KeyRepository_CreateKey(t *testing.T) {
	t.Run("should CreateKey execute correctly", func(t *testing.T) {
		sut := makeKeyRepositorySut()

		sut.guidGen.On("V4").Return("some_guid")
		sut.redisClient.On(
			"Set", sut.ctx,
			getRedisKeyByKey(sut.keyMocked),
			models.ToKeyModel(sut.keyMocked),
			time.Duration(0),
		).Return(redis.NewStatusResult("", nil))

		result, err := sut.repo.CreateKey(sut.ctx, sut.keyMocked, 0)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		sut.guidGen.AssertExpectations(t)
		sut.redisClient.AssertExpectations(t)
	})

	t.Run("should CreateKey returns error", func(t *testing.T) {
		sut := makeKeyRepositorySut()

		sut.guidGen.On("V4").Return("some_guid")
		sut.redisClient.On(
			"Set", sut.ctx,
			getRedisKeyByKey(sut.keyMocked),
			models.ToKeyModel(sut.keyMocked),
			time.Duration(0),
		).Return(redis.NewStatusResult("", errors.New("some error")))
		sut.logger.On("Error", "[KeyRepository::CreateKey] - Error: some error", []zap.Field(nil))

		result, err := sut.repo.CreateKey(sut.ctx, sut.keyMocked, 0)

		assert.Error(t, err)
		assert.NotNil(t, result)
		sut.guidGen.AssertExpectations(t)
		sut.redisClient.AssertExpectations(t)
		sut.logger.AssertExpectations(t)
	})
}

func Test_KeyRepository_GetById(t *testing.T) {
	t.Run("should GetKeyByID execute correctly", func(t *testing.T) {
		sut := makeKeyRepositorySut()

		marshal, _ := json.Marshal(models.ToKeyModel(sut.keyMocked))
		strCmd := redis.NewStringCmd(sut.ctx)
		strCmd.SetVal(string(marshal))

		sut.redisClient.On(
			"Get",
			sut.ctx,
			getRedisKeyByIDs(sut.keyMocked.UserID, sut.keyMocked.KeyID),
		).Return(strCmd)

		result, err := sut.repo.GetKeyByID(sut.ctx, sut.keyMocked.UserID, sut.keyMocked.KeyID)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		sut.redisClient.AssertExpectations(t)
		sut.logger.AssertExpectations(t)
	})

	t.Run("should GetKeyByID return error if some error occur in redis", func(t *testing.T) {
		sut := makeKeyRepositorySut()

		strCmd := redis.NewStringCmd(sut.ctx)
		strCmd.SetErr(errors.New(""))

		sut.redisClient.On(
			"Get",
			sut.ctx,
			getRedisKeyByIDs(sut.keyMocked.UserID, sut.keyMocked.KeyID),
		).Return(strCmd)
		sut.logger.On("Error", "[KeyRepository::GetKeyByID] - Error: ", []zap.Field(nil))

		result, err := sut.repo.GetKeyByID(sut.ctx, sut.keyMocked.UserID, sut.keyMocked.KeyID)

		assert.Error(t, err)
		assert.NotNil(t, result)
		sut.redisClient.AssertExpectations(t)
		sut.logger.AssertExpectations(t)
	})

	t.Run("should GetKeyByID return error if the value received from redis is invalid", func(t *testing.T) {
		sut := makeKeyRepositorySut()

		strCmd := redis.NewStringCmd(sut.ctx)
		strCmd.SetVal("")

		sut.redisClient.On(
			"Get",
			sut.ctx,
			getRedisKeyByIDs(sut.keyMocked.UserID, sut.keyMocked.KeyID),
		).Return(strCmd)
		sut.logger.On("Error", "[KeyRepository::GetKeyByID] - Error: unexpected end of JSON input", []zap.Field(nil))

		result, err := sut.repo.GetKeyByID(sut.ctx, sut.keyMocked.UserID, sut.keyMocked.KeyID)

		assert.Error(t, err)
		assert.NotNil(t, result)
		sut.redisClient.AssertExpectations(t)
		sut.logger.AssertExpectations(t)
	})

	t.Run("should GetKeyByID return error if the value received from redis do not have rsa keys", func(t *testing.T) {
		sut := makeKeyRepositorySut()

		sut.keyMocked.PriKey = nil
		sut.keyMocked.PubKey = nil
		marshal, _ := json.Marshal(sut.keyMocked)
		strCmd := redis.NewStringCmd(sut.ctx)
		strCmd.SetVal(string(marshal))

		sut.redisClient.On(
			"Get",
			sut.ctx,
			getRedisKeyByIDs(sut.keyMocked.UserID, sut.keyMocked.KeyID),
		).Return(strCmd)
		sut.logger.On("Error", "[KeyRepository::GetKeyByID] - Error: asn1: syntax error: sequence truncated", []zap.Field(nil))

		result, err := sut.repo.GetKeyByID(sut.ctx, sut.keyMocked.UserID, sut.keyMocked.KeyID)

		assert.Error(t, err)
		assert.NotNil(t, result)
		sut.redisClient.AssertExpectations(t)
		sut.logger.AssertExpectations(t)
	})
}

type keyRepositorySutRtn struct {
	repo             interfaces.IKeyRepository
	logger           *logger.LoggerSpy
	guidGen          *guidGenerator.GuidGeneratorSpy
	miniRedis        *miniredis.Miniredis
	redisClient      *redismock.ClientMock
	ctx              context.Context
	privateKeyMocked *rsa.PrivateKey
	keyMocked        valueObjects.Key
	timeMocked       time.Time
}

func makeKeyRepositorySut() keyRepositorySutRtn {
	logger := logger.NewLoggerSpy()
	guidGen := guidGenerator.NewGUidGeneratorSpy()

	miniRedis, _ := miniredis.Run()
	redisClient := redismock.NewNiceMock(redis.NewClient(&redis.Options{
		Addr: miniRedis.Addr(),
	}))

	timeMocked := time.Now()
	now = func() time.Time {
		return timeMocked
	}

	ctx := context.Background()
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	key := valueObjects.Key{
		ID:        "some_guid",
		UserID:    "some_guid",
		KeyID:     "some_guid",
		PriKey:    privateKey,
		PubKey:    &privateKey.PublicKey,
		CreatedAt: timeMocked,
		ExpiredAt: timeMocked,
	}

	repo := NewKeyRepository(logger, guidGen, redisClient)
	return keyRepositorySutRtn{repo, logger, guidGen, miniRedis, redisClient, ctx, privateKey, key, timeMocked}
}
