package cmd

import (
	"jwemanager/pkg/app/interfaces"
	"jwemanager/pkg/app/usecases"
	httpServer "jwemanager/pkg/infra/http_server"
	"jwemanager/pkg/infra/logger"
	"jwemanager/pkg/interfaces/http/handlers"
	"jwemanager/pkg/interfaces/http/presenters"
)

type webApiContainer struct {
	logger     interfaces.ILogger
	httpServer httpServer.IHttpServer

	keysRoutes presenters.IKeysRoutes
}

func NewContainer() webApiContainer {
	logger := logger.NewLogger()
	httpServer := httpServer.NewHttpServer(logger)

	createKeyUseCase := usecases.NewCreateKeyUseCase()
	keyHandlers := handlers.NewKeysHandlers(logger, struct{}{}, createKeyUseCase)
	keysRoutes := presenters.NewKeysRoutes(logger, keyHandlers)

	return webApiContainer{
		logger,
		httpServer,

		keysRoutes,
	}
}
