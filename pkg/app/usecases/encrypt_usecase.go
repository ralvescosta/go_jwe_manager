package usecases

import (
	"context"

	"jwemanager/pkg/app/interfaces"
	"jwemanager/pkg/domain/usecases"
	valueObjects "jwemanager/pkg/domain/value_objects"
)

type encryptUseCase struct {
	logger  interfaces.ILogger
	keyRepo interfaces.IKeyRepository
	crypto  interfaces.ICrypto
}

func (pst encryptUseCase) Encrypt(ctx context.Context, data valueObjects.EncryptValueObject) (valueObjects.EncryptedValueObject, error) {
	key, err := pst.keyRepo.GetKeyByID(ctx, data.UserID, data.KeyID)
	if err != nil {
		return valueObjects.EncryptedValueObject{}, err
	}

	encrypted, err := pst.crypto.Encrypt(key.PubKey, data.Data)
	if err != nil {
		return valueObjects.EncryptedValueObject{}, err
	}

	return valueObjects.EncryptedValueObject{
		EncryptedData: string(encrypted),
	}, nil
}

func NewEncryptUseCase(
	logger interfaces.ILogger,
	keyRepo interfaces.IKeyRepository,
	crypto interfaces.ICrypto,
) usecases.IEncryptUseCase {
	return encryptUseCase{logger, keyRepo, crypto}
}
