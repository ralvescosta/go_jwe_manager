package usecases

import (
	"context"
	valueObjects "jwemanager/pkg/domain/value_objects"
)

type IDecryptUseCase interface {
	Decrypt(ctx context.Context, data valueObjects.DecryptValueObject) (valueObjects.DecryptedValueObject, error)
}
