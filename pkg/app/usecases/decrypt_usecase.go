package usecases

import (
	"context"

	"jwemanager/pkg/app/interfaces"
	"jwemanager/pkg/domain/usecases"
	valueObjects "jwemanager/pkg/domain/value_objects"
)

type decrypt struct {
	logger  interfaces.ILogger
	keyRepo interfaces.IKeyRepository
	crypto  interfaces.ICrypto
}

func (pst decrypt) Decrypt(ctx context.Context, data valueObjects.DecryptValueObject) (valueObjects.DecryptedValueObject, error) {
	key, err := pst.keyRepo.GetKeyByID(ctx, data.UserID, data.KeyID)
	if err != nil {
		return valueObjects.DecryptedValueObject{}, err
	}

	result, err := pst.crypto.Decrypt(key.PriKey, []byte(data.EncryptedData))
	if err != nil {
		return valueObjects.DecryptedValueObject{}, err
	}

	return valueObjects.DecryptedValueObject{
		Data: result,
	}, nil
}

func NewDecryptUseCase(
	logger interfaces.ILogger,
	keyRepo interfaces.IKeyRepository,
	crypto interfaces.ICrypto,
) usecases.IDecryptUseCase {
	return decrypt{logger, keyRepo, crypto}
}
