package presenters

import (
	"jwemanager/pkg/app/interfaces"
	"jwemanager/pkg/infra/adapters"
	httpServer "jwemanager/pkg/infra/http_server"
	"jwemanager/pkg/interfaces/http/handlers"
)

type IHealthRoutes interface {
	Register(httpServer httpServer.IHttpServer)
}

type healthRoutes struct {
	logger   interfaces.ILogger
	handlers handlers.IHealthHandler
}

func (pst healthRoutes) Register(httpServer httpServer.IHttpServer) {
	httpServer.RegistreRoute("GET", "/api/v1/health", adapters.HandlerAdapt(pst.handlers.Check, pst.logger))
}

func NewHealthRoutes(logger interfaces.ILogger, healthHandler handlers.IHealthHandler) IHealthRoutes {
	return healthRoutes{logger, healthHandler}
}
