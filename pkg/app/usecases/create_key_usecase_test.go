package usecases

import (
	"context"
	"crypto/rsa"
	"testing"
	"time"

	"jwemanager/pkg/domain/usecases"
	valueObjects "jwemanager/pkg/domain/value_objects"
	guidGenerator "jwemanager/pkg/infra/guid_generator"
	keyGenerator "jwemanager/pkg/infra/key_generator"
	"jwemanager/pkg/infra/logger"
	"jwemanager/pkg/infra/repositories"
)

func Test_CreateKeyUseCase(t *testing.T) {
	t.Run("", func(t *testing.T) {
		sut := makeCreateKeySut()

		ctx := context.Background()
		privateKey := &rsa.PrivateKey{}
		key := valueObjects.Key{
			KeyID:     "some_guid",
			PriKey:    privateKey,
			PubKey:    &privateKey.PublicKey,
			ExpiredAt: time.Now().Add(time.Hour * 720),
		}

		sut.guidGenerator.On("V4").Return("some_guid")
		sut.keyGenerator.On("GenerateKey").Return(privateKey, nil)
		sut.repository.On("CreateKey", ctx, key).Return(key, nil)

		sut.useCase.Execute(ctx, key)
	})
}

type createKeySutRtn struct {
	useCase       usecases.ICreateKeyUseCase
	logger        *logger.LoggerSpy
	repository    *repositories.KeyRepositorySpy
	guidGenerator *guidGenerator.GuidGeneratorSpy
	keyGenerator  *keyGenerator.KeyGeneratorSpy
}

func makeCreateKeySut() createKeySutRtn {
	logger := logger.NewLoggerSpy()
	repository := repositories.NewKeyRepositorySpy()
	guidGenerator := guidGenerator.NewGUidGeneratorSpy()
	keyGenerator := keyGenerator.NewKeyGeneratorSpy()

	useCase := NewCreateKeyUseCase(
		logger,
		repository,
		guidGenerator,
		keyGenerator,
	)

	return createKeySutRtn{useCase, logger, repository, guidGenerator, keyGenerator}
}
