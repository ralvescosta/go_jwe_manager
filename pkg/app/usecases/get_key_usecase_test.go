package usecases

import (
	"context"
	"jwemanager/pkg/app/errors"
	"jwemanager/pkg/domain/usecases"
	valueObjects "jwemanager/pkg/domain/value_objects"
	"jwemanager/pkg/infra/logger"
	"jwemanager/pkg/infra/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetByKey(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := makeGetKeySut()

		sut.keyRepo.On("GetKeyByID", sut.ctx, sut.userID, sut.keyID).Return(sut.keyMocked, nil)

		result, err := sut.useCase.GetKey(sut.ctx, sut.userID, sut.keyID)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		sut.keyRepo.AssertExpectations(t)
	})

	t.Run("should return error if some error occur in keyRepo", func(t *testing.T) {
		sut := makeGetKeySut()

		sut.keyRepo.On("GetKeyByID", sut.ctx, sut.userID, sut.keyID).Return(sut.keyMocked, errors.NewInternalError("some error"))

		_, err := sut.useCase.GetKey(sut.ctx, sut.userID, sut.keyID)

		assert.Error(t, err)
		sut.keyRepo.AssertExpectations(t)
	})
}

type getKeySutRtn struct {
	useCase   usecases.IGetKeyUseCase
	logger    *logger.LoggerSpy
	keyRepo   *repositories.KeyRepositorySpy
	ctx       context.Context
	userID    string
	keyID     string
	keyMocked valueObjects.Key
}

func makeGetKeySut() getKeySutRtn {
	logger := logger.NewLoggerSpy()
	keyRepo := repositories.NewKeyRepositorySpy()
	useCase := NewGetKeyUseCase(logger, keyRepo)

	ctx := context.Background()
	userID := "user_id"
	keyID := "key_id"
	keyMocked := valueObjects.Key{}

	return getKeySutRtn{useCase, logger, keyRepo, ctx, userID, keyID, keyMocked}
}
