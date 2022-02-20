package cmd

import (
	"jwemanager/pkg/infra/environments"
)

func HttpServer() error {
	env := environments.NewEnvironment()
	if err := env.Configure(); err != nil {
		return err
	}

	container, err := NewContainer(env)
	if err != nil {
		return err
	}

	container.httpServer.Default()
	container.keysRoutes.Register(container.httpServer)
	container.httpServer.Setup()

	if err := container.httpServer.Run(); err != nil {
		return err
	}

	return nil
}
