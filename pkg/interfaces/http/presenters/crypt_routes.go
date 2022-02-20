package presenters

import (
	"jwemanager/pkg/app/interfaces"
	"jwemanager/pkg/infra/adapters"
	httpServer "jwemanager/pkg/infra/http_server"
	"jwemanager/pkg/interfaces/http/handlers"
)

type ICryptRoutes interface {
	Register(httpServer httpServer.IHttpServer)
}

type cryptRoutes struct {
	logger   interfaces.ILogger
	handlers handlers.ICryptHandler
}

func (pst cryptRoutes) Register(httpServer httpServer.IHttpServer) {
	httpServer.RegistreRoute("POST", "/api/v1/encrypt", adapters.HandlerAdapt(pst.handlers.Encrypt, pst.logger))
	httpServer.RegistreRoute("POST", "/api/v1/decrypt", adapters.HandlerAdapt(pst.handlers.Decrypt, pst.logger))
}

func NewCryptRoutes(logger interfaces.ILogger, handlers handlers.ICryptHandler) ICryptRoutes {
	return cryptRoutes{logger, handlers}
}
