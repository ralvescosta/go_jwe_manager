package cmd

import (
	"jwemanager/pkg/app/interfaces"
	"jwemanager/pkg/app/usecases"
	"jwemanager/pkg/infra/crypto"
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

	keysRoutes   presenters.IKeysRoutes
	cryptRoutes  presenters.ICryptRoutes
	healthRoutes presenters.IHealthRoutes
}

func NewContainer(env interfaces.IEnvironments) (webApiContainer, error) {
	var shutdown = make(chan bool)

	logger, err := logger.NewLogger()
	if err != nil {
		return webApiContainer{}, err
	}
	httpServer := httpServer.NewHttpServer(env, logger, shutdown)

	rdb, err := database.Connection(logger, shutdown)
	if err != nil {
		return webApiContainer{}, err
	}

	httpResponseFactory := factories.NewHttpResponseFactory()
	vValidator := validator.NewValidator()
	guidGen := guidGenerator.NewGuidGenerator()
	keyGen := keyGenerator.NewKeyGenerator(logger)
	keyRepository := repositories.NewKeyRepository(logger, guidGen, rdb)
	crypto := crypto.NewCrypto(logger)

	createKeyUseCase := usecases.NewCreateKeyUseCase(logger, keyRepository, guidGen, keyGen)
	getKeyUseCase := usecases.NewGetKeyUseCase(logger, keyRepository)
	keyHandlers := handlers.NewKeysHandlers(logger, vValidator, httpResponseFactory, createKeyUseCase, getKeyUseCase)
	keysRoutes := presenters.NewKeysRoutes(logger, keyHandlers)

	encryptUseCase := usecases.NewEncryptUseCase(logger, keyRepository, crypto)
	decryptUseCase := usecases.NewDecryptUseCase(logger, keyRepository, crypto)
	cryptHandlers := handlers.NewCryptHandler(logger, vValidator, httpResponseFactory, encryptUseCase, decryptUseCase)
	cryptRoutes := presenters.NewCryptRoutes(logger, cryptHandlers)

	healthHandlers := handlers.NewHealthHandler(logger, httpResponseFactory, rdb)
	healthRoutes := presenters.NewHealthRoutes(logger, healthHandlers)

	return webApiContainer{
		logger,
		httpServer,

		keysRoutes,
		cryptRoutes,
		healthRoutes,
	}, nil
}
