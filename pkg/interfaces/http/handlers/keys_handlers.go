package handlers

import (
	"encoding/json"

	"jwemanager/pkg/app/interfaces"
	"jwemanager/pkg/domain/usecases"
	httpServer "jwemanager/pkg/infra/http_server"
	"jwemanager/pkg/interfaces/http/factories"
	vm "jwemanager/pkg/interfaces/http/view_models"
)

type IKeysHandler interface {
	Create(httpRequest httpServer.HttpRequest) httpServer.HttpResponse
	FindOne(httpRequest httpServer.HttpRequest) httpServer.HttpResponse
}

type KeysHandler struct {
	logger           interfaces.ILogger
	validator        interfaces.IValidator
	httpResFactory   factories.HttpResponseFactory
	createKeyUseCase usecases.ICreateKeyUseCase
	getKeyUseCse     usecases.IGetKeyUseCase
}

func (pst KeysHandler) Create(httpRequest httpServer.HttpRequest) httpServer.HttpResponse {
	vModel := vm.CreateKeyViewModel{}
	if err := json.Unmarshal(httpRequest.Body, &vModel); err != nil {
		return pst.httpResFactory.BadRequest("body is required", nil)
	}

	if validationErrs := pst.validator.ValidateStruct(vModel); validationErrs != nil {
		pst.logger.Error(validationErrs[0].Message)
		return pst.httpResFactory.BadRequest(validationErrs[0].Message, nil)
	}

	result, err := pst.createKeyUseCase.Execute(httpRequest.Ctx, vModel.ToValueObject(), vModel.TimeToExpiration)
	if err != nil {
		return pst.httpResFactory.ErrorResponseMapper(err, nil)
	}

	return pst.httpResFactory.Created(vm.NewResultKeyViewModel(result), nil)
}

func (pst KeysHandler) FindOne(httpRequest httpServer.HttpRequest) httpServer.HttpResponse {
	userID, ok := httpRequest.Params["user_id"]
	if !ok {
		return pst.httpResFactory.BadRequest("user_id is required", nil)
	}
	keyID, ok := httpRequest.Params["key_id"]
	if !ok {
		return pst.httpResFactory.BadRequest("key_id is required", nil)
	}

	result, err := pst.getKeyUseCse.GetKey(httpRequest.Ctx, userID, keyID)
	if err != nil {
		return pst.httpResFactory.ErrorResponseMapper(err, nil)
	}

	return pst.httpResFactory.Ok(vm.NewResultKeyViewModel(result), nil)
}

func NewKeysHandlers(
	logger interfaces.ILogger,
	validator interfaces.IValidator,
	httpFactory factories.HttpResponseFactory,
	createKeyUseCase usecases.ICreateKeyUseCase,
	getKeyUseCase usecases.IGetKeyUseCase,
) IKeysHandler {
	return KeysHandler{
		logger,
		validator,
		httpFactory,
		createKeyUseCase,
		getKeyUseCase,
	}
}
