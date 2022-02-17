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
	logger interfaces.ILogger
	factories.HttpResponseFactory
	createKeyUseCase usecases.ICreateKeyUseCase
}

func (pst KeysHandler) Create(httpRequest httpServer.HttpRequest) httpServer.HttpResponse {
	vModel := vm.CreateKeyViewModel{}
	if err := json.Unmarshal(httpRequest.Body, &vModel); err != nil {
		return pst.BadRequest("body is required", nil)
	}

	result, err := pst.createKeyUseCase.Execute(vModel.ToDto())
	if err != nil {
		return pst.BadRequest("some error occur", nil)
	}

	return pst.Created(result, nil)
}

func (pst KeysHandler) FindOne(httpRequest httpServer.HttpRequest) httpServer.HttpResponse {
	return pst.Ok(nil, nil)
}

func NewKeysHandlers(
	logger interfaces.ILogger,
	httpFactory factories.HttpResponseFactory,
	createKeyUseCase usecases.ICreateKeyUseCase,
) IKeysHandler {
	return KeysHandler{
		logger,
		httpFactory,
		createKeyUseCase,
	}
}
