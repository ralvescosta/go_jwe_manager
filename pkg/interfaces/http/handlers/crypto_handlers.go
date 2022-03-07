package handlers

import (
	"encoding/json"

	"jwemanager/pkg/app/interfaces"
	"jwemanager/pkg/domain/usecases"
	httpServer "jwemanager/pkg/infra/http_server"
	"jwemanager/pkg/interfaces/http/factories"
	vm "jwemanager/pkg/interfaces/http/view_models"
)

type ICryptoHandler interface {
	Encrypt(httpRequest httpServer.HttpRequest) httpServer.HttpResponse
	Decrypt(httpRequest httpServer.HttpRequest) httpServer.HttpResponse
}

type cryptoHandler struct {
	logger         interfaces.ILogger
	validator      interfaces.IValidator
	httpResFactory factories.HttpResponseFactory
	encryptUseCase usecases.IEncryptUseCase
	decryptUseCase usecases.IDecryptUseCase
}

func (pst cryptoHandler) Encrypt(httpRequest httpServer.HttpRequest) httpServer.HttpResponse {
	vModel := vm.EncryptViewModel{}
	if err := json.Unmarshal(httpRequest.Body, &vModel); err != nil {
		return pst.httpResFactory.BadRequest("body is required", nil)
	}

	if validationErrs := pst.validator.ValidateStruct(vModel); validationErrs != nil {
		pst.logger.Error(validationErrs[0].Message)
		return pst.httpResFactory.BadRequest(validationErrs[0].Message, nil)
	}

	result, err := pst.encryptUseCase.Encrypt(httpRequest.Ctx, vModel.ToValueObject())
	if err != nil {
		return pst.httpResFactory.ErrorResponseMapper(err, nil)
	}

	return pst.httpResFactory.Ok(vm.ToEncryptedViewModel(result), nil)
}

func (pst cryptoHandler) Decrypt(httpRequest httpServer.HttpRequest) httpServer.HttpResponse {
	vModel := vm.DecryptViewModel{}
	if err := json.Unmarshal(httpRequest.Body, &vModel); err != nil {
		return pst.httpResFactory.BadRequest("body is required", nil)
	}

	if validationErrs := pst.validator.ValidateStruct(vModel); validationErrs != nil {
		pst.logger.Error(validationErrs[0].Message)
		return pst.httpResFactory.BadRequest(validationErrs[0].Message, nil)
	}

	result, err := pst.decryptUseCase.Decrypt(httpRequest.Ctx, vModel.ToValueObject())
	if err != nil {
		return pst.httpResFactory.ErrorResponseMapper(err, nil)
	}

	return pst.httpResFactory.Ok(vm.ToDecryptedViewModel(result), nil)
}

func NewCryptoHandler(
	logger interfaces.ILogger,
	validator interfaces.IValidator,
	httpResFactory factories.HttpResponseFactory,
	encryptUseCase usecases.IEncryptUseCase,
	decryptUseCase usecases.IDecryptUseCase,
) ICryptoHandler {
	return cryptoHandler{logger, validator, httpResFactory, encryptUseCase, decryptUseCase}
}
