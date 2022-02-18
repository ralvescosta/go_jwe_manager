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
	logger    interfaces.ILogger
	validator interfaces.IValidator
	factories.HttpResponseFactory
	createKeyUseCase usecases.ICreateKeyUseCase
}

func (pst KeysHandler) Create(httpRequest httpServer.HttpRequest) httpServer.HttpResponse {
	vModel := vm.CreateKeyViewModel{}
	if err := json.Unmarshal(httpRequest.Body, &vModel); err != nil {
		return pst.BadRequest("body is required", nil)
	}

	if validationErrs := pst.validator.ValidateStruct(vModel); validationErrs != nil {
		pst.logger.Error(validationErrs[0].Message)
		return pst.BadRequest(validationErrs[0].Message, nil)
	}

	result, err := pst.createKeyUseCase.Execute(httpRequest.Ctx, vModel.ToDto())
	if err != nil {
		return pst.BadRequest("some error occur", nil)
	}

	return pst.Created(vm.NewCreatedKeyViewModel(result), nil)
}

func (pst KeysHandler) FindOne(httpRequest httpServer.HttpRequest) httpServer.HttpResponse {
	return pst.Ok(nil, nil)
}

func NewKeysHandlers(
	logger interfaces.ILogger,
	validator interfaces.IValidator,
	httpFactory factories.HttpResponseFactory,
	createKeyUseCase usecases.ICreateKeyUseCase,
) IKeysHandler {
	return KeysHandler{
		logger,
		validator,
		httpFactory,
		createKeyUseCase,
	}
}
