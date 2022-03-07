package presenters

import (
	httpServer "jwemanager/pkg/infra/http_server"
	"jwemanager/pkg/infra/logger"
	"jwemanager/pkg/interfaces/http/handlers"
	"testing"
)

func Test_Health_Register(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := makeHealthSut()

		sut.handler.On("Check").Return(httpServer.HttpResponse{})
		sut.server.On("RegisterRoute", "GET", "/api/v1/health").Return(nil)

		sut.routes.Register(sut.server)

		sut.server.AssertExpectations(t)
	})
}

type healthSutRtn struct {
	routes  IHealthRoutes
	handler *handlers.HealthHandlersSpy
	logger  *logger.LoggerSpy
	server  *httpServer.HTTPServerSpy
}

func makeHealthSut() healthSutRtn {
	logger := logger.NewLoggerSpy()
	handler := handlers.NewHealthHandlerSpy()

	routes := NewHealthRoutes(logger, handler)
	server := httpServer.NewHTTPServerSpy()

	return healthSutRtn{routes, handler, logger, server}
}
