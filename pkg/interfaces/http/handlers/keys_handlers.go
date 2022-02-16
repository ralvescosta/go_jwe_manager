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
	FindAll(httpRequest httpServer.HttpRequest) httpServer.HttpResponse
	Update(httpRequest httpServer.HttpRequest) httpServer.HttpResponse
	Delete(httpRequest httpServer.HttpRequest) httpServer.HttpResponse
}

type KeysHandler struct {
	logger interfaces.ILogger
	factories.HttpResponseFactory
	usecases usecases.ICreateKeyUseCase
}

func (pst KeysHandler) Create(httpRequest httpServer.HttpRequest) httpServer.HttpResponse {
	vModel := vm.CreateKeyViewModel{}
	if err := json.Unmarshal(httpRequest.Body, &vModel); err != nil {
		return pst.BadRequest("body is required", nil)
	}

	result, err := pst.usecases.Execute(vModel.ToDto())
	if err != nil {
		return pst.BadRequest("some error occur", nil)
	}

	return pst.Created(result, nil)
}

func (pst KeysHandler) FindOne(httpRequest httpServer.HttpRequest) httpServer.HttpResponse {
	return pst.Ok(nil, nil)
}
