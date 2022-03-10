package presenters

import (
	"jwemanager/pkg/app/interfaces"
	"jwemanager/pkg/infra/adapters"
	httpServer "jwemanager/pkg/infra/http_server"
	"jwemanager/pkg/interfaces/http/handlers"
)

type IKeysRoutes interface {
	Register(httpServer httpServer.IHTTPServer)
}

type KeysRoutes struct {
	logger   interfaces.ILogger
	handlers handlers.IKeysHandler
}

func (pst KeysRoutes) Register(httpServer httpServer.IHTTPServer) {
	httpServer.RegisterRoute("POST", "/api/v1/keys", adapters.HandlerAdapt(pst.handlers.Create, pst.logger))
	httpServer.RegisterRoute("GET", "/api/v1/keys/:user_id/:key_id", adapters.HandlerAdapt(pst.handlers.FindOne, pst.logger))
}

func NewKeysRoutes(logger interfaces.ILogger, handlers handlers.IKeysHandler) IKeysRoutes {
	return KeysRoutes{
		logger:   logger,
		handlers: handlers,
	}
}
