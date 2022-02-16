package presenters

import (
	"jwemanager/pkg/app/interfaces"
	"jwemanager/pkg/infra/adapters"
	httpServer "jwemanager/pkg/infra/http_server"
	"jwemanager/pkg/interfaces/http/handlers"
)

type IKeysRoutes interface {
	Register(httpServer httpServer.IHttpServer)
}

type KeysRoutes struct {
	handlers handlers.IKeysHandler
	logger   interfaces.ILogger
}

func (pst KeysRoutes) Register(httpServer httpServer.IHttpServer) {
	httpServer.RegistreRoute("POST", "/api/v1/keys", adapters.HandlerAdapt(pst.handlers.Create, pst.logger))
	httpServer.RegistreRoute("GET", "/api/v1/keys/:id", adapters.HandlerAdapt(pst.handlers.FindOne, pst.logger))
}

func NewKeysRoutes(handlers handlers.IKeysHandler, logger interfaces.ILogger) IKeysRoutes {
	return KeysRoutes{
		handlers: handlers,
		logger:   logger,
	}
}