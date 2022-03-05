package usecases

import (
	"context"
	"crypto/rsa"
	"testing"

	"jwemanager/pkg/app/errors"
	"jwemanager/pkg/domain/usecases"
	valueObjects "jwemanager/pkg/domain/value_objects"
	"jwemanager/pkg/infra/crypto"
	"jwemanager/pkg/infra/logger"
	"jwemanager/pkg/infra/repositories"

	"github.com/stretchr/testify/assert"
)

func Test_Decrypt(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := makeDecryptSut()

		sut.keyRepo.On("GetKeyByID", sut.ctx, sut.decryptDataMocked.UserID, sut.decryptDataMocked.KeyID).Return(sut.keyMocked, nil)
		sut.crypto.On("Decrypt", sut.keyMocked.PriKey, sut.decryptDataMocked.EncryptedData).Return(make(map[string]interface{}), nil)

		result, err := sut.useCase.Decrypt(sut.ctx, sut.decryptDataMocked)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		sut.keyRepo.AssertExpectations(t)
		sut.crypto.AssertExpectations(t)
	})

	t.Run("should return error if some error occur in keyRepo", func(t *testing.T) {
		sut := makeDecryptSut()

		sut.keyRepo.On("GetKeyByID", sut.ctx, sut.decryptDataMocked.UserID, sut.decryptDataMocked.KeyID).Return(sut.keyMocked, errors.NewInternalError("some error"))

		_, err := sut.useCase.Decrypt(sut.ctx, sut.decryptDataMocked)
		assert.Error(t, err)
		sut.keyRepo.AssertExpectations(t)
	})

	t.Run("should return error if some error occur in crypto", func(t *testing.T) {
		sut := makeDecryptSut()

		sut.keyRepo.On("GetKeyByID", sut.ctx, sut.decryptDataMocked.UserID, sut.decryptDataMocked.KeyID).Return(sut.keyMocked, nil)
		sut.crypto.On("Decrypt", sut.keyMocked.PriKey, sut.decryptDataMocked.EncryptedData).Return(make(map[string]interface{}), errors.NewInternalError("some error"))

		_, err := sut.useCase.Decrypt(sut.ctx, sut.decryptDataMocked)
		assert.Error(t, err)
		sut.keyRepo.AssertExpectations(t)
		sut.crypto.AssertExpectations(t)
	})
}

type decryptSutRtn struct {
	useCase           usecases.IDecryptUseCase
	logger            *logger.LoggerSpy
	keyRepo           *repositories.KeyRepositorySpy
	crypto            *crypto.CryptoSpy
	ctx               context.Context
	decryptDataMocked valueObjects.DecryptValueObject
	keyMocked         valueObjects.Key
}

func makeDecryptSut() decryptSutRtn {
	logger := logger.NewLoggerSpy()
	keyRepo := repositories.NewKeyRepositorySpy()
	crypto := crypto.NewCryptoSpy()

	useCase := NewDecryptUseCase(logger, keyRepo, crypto)

	ctx := context.Background()
	decryptDataMocked := valueObjects.DecryptValueObject{}
	privKey := &rsa.PrivateKey{}
	keyMocked := valueObjects.Key{
		PriKey: privKey,
		PubKey: &privKey.PublicKey,
	}

	return decryptSutRtn{useCase, logger, keyRepo, crypto, ctx, decryptDataMocked, keyMocked}
}
