package usecases

import (
	"context"
	"crypto/rsa"
	"testing"
	"time"

	"jwemanager/pkg/app/errors"
	"jwemanager/pkg/domain/usecases"
	valueObjects "jwemanager/pkg/domain/value_objects"
	guidGenerator "jwemanager/pkg/infra/guid_generator"
	keyGenerator "jwemanager/pkg/infra/key_generator"
	"jwemanager/pkg/infra/logger"
	"jwemanager/pkg/infra/repositories"

	"github.com/stretchr/testify/assert"
)

func Test_CreateKeyUseCase(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := makeCreateKeySut()

		timeToExpiration := 0

		sut.guidGenerator.On("V4").Return("some_guid")
		sut.keyGenerator.On("GenerateKey").Return(sut.privateKeyMocked, nil)
		sut.repository.On("CreateKey", sut.ctx, sut.keyMocked, timeToExpiration).Return(sut.keyMocked, nil)

		_, err := sut.useCase.Execute(sut.ctx, sut.keyMocked, timeToExpiration)

		assert.NoError(t, err)
		sut.guidGenerator.AssertExpectations(t)
		sut.keyGenerator.AssertExpectations(t)
		sut.repository.AssertExpectations(t)
	})

	t.Run("should return error if keyGenerator return error", func(t *testing.T) {
		sut := makeCreateKeySut()

		timeToExpiration := 0

		sut.guidGenerator.On("V4").Return("some_guid")
		sut.keyGenerator.On("GenerateKey").Return(sut.privateKeyMocked, errors.NewInternalError("some error"))

		_, err := sut.useCase.Execute(sut.ctx, sut.keyMocked, timeToExpiration)

		assert.Error(t, err)
		sut.guidGenerator.AssertExpectations(t)
		sut.keyGenerator.AssertExpectations(t)
	})

	t.Run("should return error if repository return error", func(t *testing.T) {
		sut := makeCreateKeySut()

		timeToExpiration := 0

		sut.guidGenerator.On("V4").Return("some_guid")
		sut.keyGenerator.On("GenerateKey").Return(sut.privateKeyMocked, nil)
		sut.repository.On("CreateKey", sut.ctx, sut.keyMocked, timeToExpiration).Return(sut.keyMocked, errors.NewInternalError("some error"))

		_, err := sut.useCase.Execute(sut.ctx, sut.keyMocked, timeToExpiration)

		assert.Error(t, err)
		sut.guidGenerator.AssertExpectations(t)
		sut.keyGenerator.AssertExpectations(t)
		sut.repository.AssertExpectations(t)
	})
}

type createKeySutRtn struct {
	useCase          usecases.ICreateKeyUseCase
	logger           *logger.LoggerSpy
	repository       *repositories.KeyRepositorySpy
	guidGenerator    *guidGenerator.GuidGeneratorSpy
	keyGenerator     *keyGenerator.KeyGeneratorSpy
	ctx              context.Context
	privateKeyMocked *rsa.PrivateKey
	keyMocked        valueObjects.Key
}

func makeCreateKeySut() createKeySutRtn {
	logger := logger.NewLoggerSpy()
	repository := repositories.NewKeyRepositorySpy()
	guidGenerator := guidGenerator.NewGuidGeneratorSpy()
	keyGenerator := keyGenerator.NewKeyGeneratorSpy()

	ctx := context.Background()
	privateKey := &rsa.PrivateKey{}
	key := valueObjects.Key{
		KeyID:     "some_guid",
		PriKey:    privateKey,
		PubKey:    &privateKey.PublicKey,
		ExpiredAt: time.Now().Add(time.Hour * 720),
	}

	useCase := NewCreateKeyUseCase(
		logger,
		repository,
		guidGenerator,
		keyGenerator,
	)

	return createKeySutRtn{
		useCase,
		logger,
		repository,
		guidGenerator,
		keyGenerator,
		ctx,
		privateKey,
		key,
	}
}
