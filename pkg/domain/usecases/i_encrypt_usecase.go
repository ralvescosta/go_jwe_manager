package usecases

import (
	"context"
	valueObjects "jwemanager/pkg/domain/value_objects"
)

type IEncryptUseCase interface {
	Encrypt(ctx context.Context, data valueObjects.EncryptValueObject) (valueObjects.EncryptedValueObject, error)
}
