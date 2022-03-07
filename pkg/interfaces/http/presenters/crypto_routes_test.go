package presenters

import (
	httpServer "jwemanager/pkg/infra/http_server"
	"jwemanager/pkg/infra/logger"
	"jwemanager/pkg/interfaces/http/handlers"
	"testing"
)

func Test_Crypto_Register(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := makeCryptoSut()

		sut.handler.On("Encrypt").Return(httpServer.HttpResponse{})
		sut.handler.On("Decrypt").Return(httpServer.HttpResponse{})
		sut.server.On("RegisterRoute", "POST", "/api/v1/encrypt").Return(nil)
		sut.server.On("RegisterRoute", "POST", "/api/v1/decrypt").Return(nil)

		sut.routes.Register(sut.server)

		sut.server.AssertExpectations(t)
	})
}

type cryptoSutRtn struct {
	routes  ICryptoRoutes
	logger  *logger.LoggerSpy
	handler *handlers.CryptoHandlersSpy
	server  *httpServer.HTTPServerSpy
}

func makeCryptoSut() cryptoSutRtn {
	logger := logger.NewLoggerSpy()
	handler := handlers.NewCryptoHandlersSpy()
	routes := NewCryptRoutes(logger, handler)
	server := httpServer.NewHTTPServerSpy()

	return cryptoSutRtn{routes, logger, handler, server}
}
