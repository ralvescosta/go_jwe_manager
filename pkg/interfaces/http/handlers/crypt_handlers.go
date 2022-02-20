package handlers

import (
	"jwemanager/pkg/app/interfaces"
	httpServer "jwemanager/pkg/infra/http_server"
	"jwemanager/pkg/interfaces/http/factories"
)

type ICryptHandler interface {
	Encrypt(httpRequest httpServer.HttpRequest) httpServer.HttpResponse
	Decrypt(httpRequest httpServer.HttpRequest) httpServer.HttpResponse
}

type cryptHandler struct {
	logger         interfaces.ILogger
	validator      interfaces.IValidator
	httpResFactory factories.HttpResponseFactory
}

func (pst cryptHandler) Encrypt(httpRequest httpServer.HttpRequest) httpServer.HttpResponse {
	return pst.httpResFactory.Ok(nil, nil)
}

func (pst cryptHandler) Decrypt(httpRequest httpServer.HttpRequest) httpServer.HttpResponse {
	return pst.httpResFactory.Ok(nil, nil)
}

func NewCryptHandler(
	logger interfaces.ILogger,
	validator interfaces.IValidator,
	httpResFactory factories.HttpResponseFactory,
) ICryptHandler {
	return cryptHandler{logger, validator, httpResFactory}
}
