package usecases

import (
	"context"
	valueObjects "jwemanager/pkg/domain/value_objects"
)

type IGetKeyUseCase interface {
	GetKey(ctx context.Context, userID, keyID string) (valueObjects.Key, error)
}
