package cmd

import (
	"jwemanager/pkg/app/interfaces"
	"jwemanager/pkg/app/usecases"
	"jwemanager/pkg/infra/database"
	guidGenerator "jwemanager/pkg/infra/guid_generator"
	httpServer "jwemanager/pkg/infra/http_server"
	keyGenerator "jwemanager/pkg/infra/key_generator"
	"jwemanager/pkg/infra/logger"
	"jwemanager/pkg/infra/repositories"
	"jwemanager/pkg/infra/validator"
	"jwemanager/pkg/interfaces/http/factories"
	"jwemanager/pkg/interfaces/http/handlers"
	"jwemanager/pkg/interfaces/http/presenters"
)

type webApiContainer struct {
	logger     interfaces.ILogger
	httpServer httpServer.IHttpServer

	keysRoutes presenters.IKeysRoutes
}

func NewContainer(env interfaces.IEnvironments) (webApiContainer, error) {
	logger := logger.NewLogger()
	httpServer := httpServer.NewHttpServer(env, logger)

	rdb, err := database.Connection(logger)
	if err != nil {
		return webApiContainer{}, err
	}

	httpResponseFactory := factories.NewHttpResponseFactory()
	vValidator := validator.NewValidator()
	guidGen := guidGenerator.NewGuidGenerator()
	keyGen := keyGenerator.NewKeyGenerator()
	keyRepository := repositories.NewKeyRepository(logger, guidGen, rdb)

	createKeyUseCase := usecases.NewCreateKeyUseCase(logger, keyRepository, keyGen)
	keyHandlers := handlers.NewKeysHandlers(logger, vValidator, httpResponseFactory, createKeyUseCase)
	keysRoutes := presenters.NewKeysRoutes(logger, keyHandlers)

	return webApiContainer{
		logger,
		httpServer,

		keysRoutes,
	}, nil
}
