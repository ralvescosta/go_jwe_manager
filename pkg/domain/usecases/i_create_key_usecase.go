package usecases

import (
	"context"
	"jwemanager/pkg/domain/dtos"
)

type ICreateKeyUseCase interface {
	Execute(ctx context.Context, key dtos.Key) (dtos.Key, error)
}
