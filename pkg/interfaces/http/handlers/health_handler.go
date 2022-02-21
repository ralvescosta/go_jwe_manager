package handlers

import (
	"context"
	"jwemanager/pkg/app/interfaces"
	httpServer "jwemanager/pkg/infra/http_server"
	"jwemanager/pkg/interfaces/http/factories"
	vm "jwemanager/pkg/interfaces/http/view_models"
	"net/http"

	"github.com/go-redis/redis/v8"
)

type IHealthHandler interface {
	Check(httpRequest httpServer.HttpRequest) httpServer.HttpResponse
}

type healthHandler struct {
	logger         interfaces.ILogger
	httpResFactory factories.HttpResponseFactory
	rdb            *redis.Client
}

func (pst healthHandler) Check(httpRequest httpServer.HttpRequest) httpServer.HttpResponse {
	_, err := pst.rdb.Ping(context.Background()).Result()

	if err != nil {
		return pst.httpResFactory.GenericResponse(http.StatusBadRequest, vm.HealthViewModel{
			Status: "No Healthy",
			Integrations: []vm.IntegrationsViewModel{
				{Name: "Redis", Status: "No Healthy"},
			},
		}, nil)
	}

	return pst.httpResFactory.Ok(vm.HealthViewModel{
		Status: "Healthy",
		Integrations: []vm.IntegrationsViewModel{
			{Name: "Redis", Status: "Healthy"},
		},
	}, nil)
}

func NewHealthHandler(
	logger interfaces.ILogger,
	httpResFactory factories.HttpResponseFactory,
	rdb *redis.Client,
) IHealthHandler {
	return healthHandler{logger, httpResFactory, rdb}
}
