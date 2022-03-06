package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"jwemanager/pkg/app/errors"
	"jwemanager/pkg/app/usecases"
	valueObjects "jwemanager/pkg/domain/value_objects"
	httpServer "jwemanager/pkg/infra/http_server"
	"jwemanager/pkg/infra/logger"
	"jwemanager/pkg/infra/validator"
	"jwemanager/pkg/interfaces/http/factories"
	vm "jwemanager/pkg/interfaces/http/view_models"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
)

func Test_Encrypt(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := makeCryptoHandlerSut()

		sut.validator.On("ValidateStruct", sut.model).Return([]valueObjects.ValidateResult(nil))
		sut.encryptUseCase.On("Encrypt", sut.req.Ctx, sut.model.ToValueObject()).Return(valueObjects.EncryptedValueObject{}, nil)

		res := sut.handler.Encrypt(sut.req)

		assert.Equal(t, http.StatusOK, res.StatusCode)
		sut.validator.AssertExpectations(t)
		sut.encryptUseCase.AssertExpectations(t)
	})

	t.Run("should return badRequest status if some error in body parser", func(t *testing.T) {
		sut := makeCryptoHandlerSut()

		sut.req.Body = nil

		res := sut.handler.Encrypt(sut.req)

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})
	t.Run("should return badRequest if some validate error occur", func(t *testing.T) {
		sut := makeCryptoHandlerSut()

		sut.validator.On("ValidateStruct", sut.model).Return([]valueObjects.ValidateResult{{IsValid: true, Field: "some", Message: "some"}})
		sut.logger.On("Error", "some", []zapcore.Field(nil))

		res := sut.handler.Encrypt(sut.req)

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		sut.validator.AssertExpectations(t)
	})

	t.Run("should return internalServerError if usecase return internalError", func(t *testing.T) {
		sut := makeCryptoHandlerSut()

		sut.validator.On("ValidateStruct", sut.model).Return([]valueObjects.ValidateResult(nil))
		sut.encryptUseCase.On("Encrypt", sut.req.Ctx, sut.model.ToValueObject()).Return(valueObjects.EncryptedValueObject{}, errors.NewInternalError("some error"))

		res := sut.handler.Encrypt(sut.req)

		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
		sut.validator.AssertExpectations(t)
		sut.encryptUseCase.AssertExpectations(t)
	})

}

func Test_Decrypt(t *testing.T) {

}

type cryptoHandlerSutRtn struct {
	handler        ICryptHandler
	logger         *logger.LoggerSpy
	validator      *validator.ValidatorSpy
	httpResFactory factories.HttpResponseFactory
	encryptUseCase *usecases.EncryptUseCaseSpy
	decryptUseCase *usecases.DecryptUseCaseSpy
	model          vm.EncryptViewModel
	req            httpServer.HttpRequest
}

func makeCryptoHandlerSut() cryptoHandlerSutRtn {
	logger := logger.NewLoggerSpy()
	validator := validator.NewValidatorSpy()
	httpResFactory := factories.NewHttpResponseFactory()
	encryptUseCase := usecases.NewEncryptUseCaseSpy()
	decryptUseCase := usecases.NewDecryptUseCaseSpy()

	handler := NewCryptHandler(logger, validator, httpResFactory, encryptUseCase, decryptUseCase)

	model := vm.EncryptViewModel{
		UserID: "45a710fd-1bb4-4171-9571-224c241838e6",
		KeyID:  "fcf7ee3e-9512-4416-8e1f-da12a667832a",
		Data:   make(map[string]interface{}),
	}
	body, _ := json.Marshal(model)
	req := httpServer.HttpRequest{
		Ctx:  context.Background(),
		Body: body,
	}

	return cryptoHandlerSutRtn{handler, logger, validator, httpResFactory, encryptUseCase, decryptUseCase, model, req}
}
