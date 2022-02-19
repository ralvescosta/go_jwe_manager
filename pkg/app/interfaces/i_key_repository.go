package interfaces

import (
	"context"
	valueObjects "jwemanager/pkg/domain/value_objects"
)

type IKeyRepository interface {
	CreateKey(context.Context, valueObjects.Key) (valueObjects.Key, error)
}
