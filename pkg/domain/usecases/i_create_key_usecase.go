package usecases

import (
	"context"

	valueObjects "jwemanager/pkg/domain/value_objects"
)

type ICreateKeyUseCase interface {
	Execute(ctx context.Context, key valueObjects.Key) (valueObjects.Key, error)
}
