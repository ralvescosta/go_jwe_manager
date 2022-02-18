package cmd

import (
	"jwemanager/pkg/infra/environments"
)

func HttpServer() error {
	err := environments.Configure()
	if err != nil {
		return err
	}

	container, err := NewContainer()
	if err != nil {
		return err
	}

	container.httpServer.Setup()

	container.keysRoutes.Register(container.httpServer)

	if err := container.httpServer.Run(); err != nil {
		return err
	}

	return nil
}
