package interfaces

import (
	"context"
	"jwemanager/pkg/domain/dtos"
)

type IKeyRepository interface {
	CreateKey(context.Context, dtos.Key) (dtos.Key, error)
}
