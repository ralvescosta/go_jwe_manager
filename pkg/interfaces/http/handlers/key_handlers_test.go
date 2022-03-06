package handlers

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"jwemanager/pkg/app/errors"
	"jwemanager/pkg/app/usecases"
	valueObjects "jwemanager/pkg/domain/value_objects"
	httpServer "jwemanager/pkg/infra/http_server"
	"jwemanager/pkg/infra/logger"
	"jwemanager/pkg/infra/validator"
	"jwemanager/pkg/interfaces/http/factories"
	vm "jwemanager/pkg/interfaces/http/view_models"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
)

func Test_Create(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := makeKeyHandlerSut()

		sut.validator.On("ValidateStruct", sut.createKeyModel).Return([]valueObjects.ValidateResult(nil))
		sut.createKeyUseCase.On("Execute", sut.createKeyReq.Ctx, sut.createKeyModel.ToValueObject(), sut.createKeyModel.TimeToExpiration).Return(sut.keyMocked, nil)

		res := sut.handler.Create(sut.createKeyReq)

		assert.Equal(t, http.StatusCreated, res.StatusCode)
		sut.validator.AssertExpectations(t)
		sut.createKeyUseCase.AssertExpectations(t)
	})

	t.Run("should return badRequest status if some error in body parser", func(t *testing.T) {
		sut := makeKeyHandlerSut()

		sut.createKeyReq.Body = nil

		res := sut.handler.Create(sut.createKeyReq)

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("should return badRequest status if has some contract err", func(t *testing.T) {
		sut := makeKeyHandlerSut()

		sut.validator.On("ValidateStruct", sut.createKeyModel).Return([]valueObjects.ValidateResult{{IsValid: false, Field: "some", Message: "some"}})
		sut.logger.On("Error", "some", []zapcore.Field(nil))
		res := sut.handler.Create(sut.createKeyReq)

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		sut.validator.AssertExpectations(t)
		sut.logger.AssertExpectations(t)
	})

	t.Run("should return internalServerError status if usecase return internalError", func(t *testing.T) {
		sut := makeKeyHandlerSut()

		sut.validator.On("ValidateStruct", sut.createKeyModel).Return([]valueObjects.ValidateResult(nil))
		sut.createKeyUseCase.On("Execute", sut.createKeyReq.Ctx, sut.createKeyModel.ToValueObject(), sut.createKeyModel.TimeToExpiration).Return(sut.keyMocked, errors.NewInternalError("some error"))

		res := sut.handler.Create(sut.createKeyReq)

		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
		sut.validator.AssertExpectations(t)
		sut.createKeyUseCase.AssertExpectations(t)
	})
}

func Test_FindOne(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := makeKeyHandlerSut()

		sut.getKeyUseCase.On("GetKey", sut.getKeyReq.Ctx, sut.userID, sut.keyID).Return(sut.keyMocked, nil)

		res := sut.handler.FindOne(sut.getKeyReq)

		assert.Equal(t, http.StatusOK, res.StatusCode)
		sut.getKeyUseCase.AssertExpectations(t)
	})

	t.Run("should return badRequest status if don't have userID", func(t *testing.T) {
		sut := makeKeyHandlerSut()

		sut.getKeyReq.Params = map[string]string{"key_id": ""}

		res := sut.handler.FindOne(sut.getKeyReq)

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("should return badRequest status if don't have keyID", func(t *testing.T) {
		sut := makeKeyHandlerSut()

		sut.getKeyReq.Params = map[string]string{"user_id": ""}

		res := sut.handler.FindOne(sut.getKeyReq)

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("should return internalServerError status if usecase return internalError", func(t *testing.T) {
		sut := makeKeyHandlerSut()

		sut.getKeyUseCase.On("GetKey", sut.getKeyReq.Ctx, sut.userID, sut.keyID).Return(sut.keyMocked, errors.NewInternalError("some error"))

		res := sut.handler.FindOne(sut.getKeyReq)

		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
		sut.getKeyUseCase.AssertExpectations(t)
	})
}

type keyHandlerSutRtn struct {
	handler          IKeysHandler
	logger           *logger.LoggerSpy
	validator        *validator.ValidatorSpy
	httpFactory      factories.HttpResponseFactory
	createKeyUseCase *usecases.CreateKeyUseCaseSpy
	getKeyUseCase    *usecases.GetKeyUseCaseSpy
	createKeyModel   vm.CreateKeyViewModel
	createKeyReq     httpServer.HttpRequest
	keyMocked        valueObjects.Key
	getKeyReq        httpServer.HttpRequest
	userID           string
	keyID            string
}

func makeKeyHandlerSut() keyHandlerSutRtn {
	logger := logger.NewLoggerSpy()
	validator := validator.NewValidatorSpy()
	httpFactory := factories.NewHttpResponseFactory()
	createKeyUseCase := usecases.NewCreateKeyUseCaseSpy()
	getKeyUseCase := usecases.NewGetKeyUseCaseSpy()

	handler := NewKeysHandlers(logger, validator, httpFactory, createKeyUseCase, getKeyUseCase)

	createKeyModel := vm.CreateKeyViewModel{
		UserID:           "45a710fd-1bb4-4171-9571-224c241838e6",
		TimeToExpiration: 0,
	}
	createKeyBody, _ := json.Marshal(createKeyModel)
	createKeyReq := httpServer.HttpRequest{
		Ctx:  context.Background(),
		Body: createKeyBody,
	}

	priv, _ := rsa.GenerateKey(rand.Reader, 2048)

	keyMocked := valueObjects.Key{
		PriKey: priv,
		PubKey: &priv.PublicKey,
	}

	userID := "9f232d03-0ef7-4284-b64d-4eb756b235b9"
	keyID := "2a0f6aa7-8a97-452d-a754-9b4971b14966"
	getKeyReq := httpServer.HttpRequest{
		Params: map[string]string{
			"user_id": userID,
			"key_id":  keyID,
		},
	}

	return keyHandlerSutRtn{handler, logger, validator, httpFactory, createKeyUseCase, getKeyUseCase, createKeyModel, createKeyReq, keyMocked, getKeyReq, userID, keyID}
}
