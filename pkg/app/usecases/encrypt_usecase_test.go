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

func Test_Encrypt(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := makeEncryptSutRtn()

		sut.keyRepo.On("GetKeyByID", sut.ctx, sut.encryptMocked.UserID, sut.encryptMocked.KeyID).Return(sut.keyMocked, nil)
		sut.crypto.On("Encrypt", sut.keyMocked.PubKey, sut.encryptMocked.Data).Return([]byte(""), nil)

		result, err := sut.useCase.Encrypt(sut.ctx, sut.encryptMocked)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		sut.keyRepo.AssertExpectations(t)
		sut.crypto.AssertExpectations(t)
	})

	t.Run("should return err if some error occur in keyRepo", func(t *testing.T) {
		sut := makeEncryptSutRtn()

		sut.keyRepo.On("GetKeyByID", sut.ctx, sut.encryptMocked.UserID, sut.encryptMocked.KeyID).Return(sut.keyMocked, errors.NewInternalError("some error"))

		_, err := sut.useCase.Encrypt(sut.ctx, sut.encryptMocked)

		assert.Error(t, err)
		sut.keyRepo.AssertExpectations(t)
	})

	t.Run("should return err if some error occur during the encrypt", func(t *testing.T) {
		sut := makeEncryptSutRtn()

		sut.keyRepo.On("GetKeyByID", sut.ctx, sut.encryptMocked.UserID, sut.encryptMocked.KeyID).Return(sut.keyMocked, nil)
		sut.crypto.On("Encrypt", sut.keyMocked.PubKey, sut.encryptMocked.Data).Return([]byte(nil), errors.NewInternalError("some error"))

		_, err := sut.useCase.Encrypt(sut.ctx, sut.encryptMocked)

		assert.Error(t, err)
		sut.keyRepo.AssertExpectations(t)
		sut.crypto.AssertExpectations(t)
	})
}

type encryptSutRtn struct {
	useCase         usecases.IEncryptUseCase
	logger          *logger.LoggerSpy
	keyRepo         *repositories.KeyRepositorySpy
	crypto          *crypto.CryptoSpy
	ctx             context.Context
	encryptMocked   valueObjects.EncryptValueObject
	encryptedMocked valueObjects.EncryptedValueObject
	keyMocked       valueObjects.Key
}

func makeEncryptSutRtn() encryptSutRtn {
	logger := logger.NewLoggerSpy()
	keyRepo := repositories.NewKeyRepositorySpy()
	crypto := crypto.NewCryptoSpy()

	useCase := NewEncryptUseCase(logger, keyRepo, crypto)

	ctx := context.Background()
	encryptMocked := valueObjects.EncryptValueObject{}
	encryptedMocked := valueObjects.EncryptedValueObject{}
	privKey := &rsa.PrivateKey{}

	keyMocked := valueObjects.Key{
		PriKey: privKey,
		PubKey: &privKey.PublicKey,
	}

	return encryptSutRtn{useCase, logger, keyRepo, crypto, ctx, encryptMocked, encryptedMocked, keyMocked}
}
