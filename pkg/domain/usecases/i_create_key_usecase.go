package usecases

import (
	"context"

	valueObjects "jwemanager/pkg/domain/value_objects"
)

type ICreateKeyUseCase interface {
	Execute(ctx context.Context, key valueObjects.Key, timeToExpiration int) (valueObjects.Key, error)
}
