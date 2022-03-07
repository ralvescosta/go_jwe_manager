package presenters

import (
	"jwemanager/pkg/app/interfaces"
	"jwemanager/pkg/infra/adapters"
	httpServer "jwemanager/pkg/infra/http_server"
	"jwemanager/pkg/interfaces/http/handlers"
)

type ICryptoRoutes interface {
	Register(httpServer httpServer.IHttpServer)
}

type cryptRoutes struct {
	logger   interfaces.ILogger
	handlers handlers.ICryptoHandler
}

func (pst cryptRoutes) Register(httpServer httpServer.IHttpServer) {
	httpServer.RegisterRoute("POST", "/api/v1/encrypt", adapters.HandlerAdapt(pst.handlers.Encrypt, pst.logger))
	httpServer.RegisterRoute("POST", "/api/v1/decrypt", adapters.HandlerAdapt(pst.handlers.Decrypt, pst.logger))
}

func NewCryptRoutes(logger interfaces.ILogger, handlers handlers.ICryptoHandler) ICryptoRoutes {
	return cryptRoutes{logger, handlers}
}
