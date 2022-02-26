package interfaces

import (
	"context"

	valueObjects "jwemanager/pkg/domain/value_objects"
)

type IKeyRepository interface {
	CreateKey(ctx context.Context, key valueObjects.Key) (valueObjects.Key, error)
	GetKeyByID(ctx context.Context, userID, keyID string) (valueObjects.Key, error)
}
