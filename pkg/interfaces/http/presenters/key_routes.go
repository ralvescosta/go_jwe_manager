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
	logger   interfaces.ILogger
	handlers handlers.IKeysHandler
}

func (pst KeysRoutes) Register(httpServer httpServer.IHttpServer) {
	httpServer.RegistreRoute("POST", "/api/v1/keys", adapters.HandlerAdapt(pst.handlers.Create, pst.logger))
	httpServer.RegistreRoute("GET", "/api/v1/keys/:user_id/:key_id", adapters.HandlerAdapt(pst.handlers.FindOne, pst.logger))
}

func NewKeysRoutes(logger interfaces.ILogger, handlers handlers.IKeysHandler) IKeysRoutes {
	return KeysRoutes{
		logger:   logger,
		handlers: handlers,
	}
}
