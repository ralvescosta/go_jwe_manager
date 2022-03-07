package presenters

import (
	httpServer "jwemanager/pkg/infra/http_server"
	"jwemanager/pkg/infra/logger"
	"jwemanager/pkg/interfaces/http/handlers"
	"testing"
)

func Test_Key_Register(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := makeKeySut()

		sut.handler.On("Create").Return(httpServer.HttpResponse{})
		sut.handler.On("FindOne").Return(httpServer.HttpResponse{})
		sut.server.On("RegisterRoute", "POST", "/api/v1/keys").Return(nil)
		sut.server.On("RegisterRoute", "GET", "/api/v1/keys/:user_id/:key_id").Return(nil)

		sut.routes.Register(sut.server)

		sut.server.AssertExpectations(t)
	})
}

type keySutRtn struct {
	routes  IKeysRoutes
	handler *handlers.KeysHandlersSpy
	logger  *logger.LoggerSpy
	server  *httpServer.HTTPServerSpy
}

func makeKeySut() keySutRtn {
	logger := logger.NewLoggerSpy()
	handler := handlers.NewKeysHandlersSpy()
	routes := NewKeysRoutes(logger, handler)
	server := httpServer.NewHTTPServerSpy()

	return keySutRtn{routes, handler, logger, server}
}
