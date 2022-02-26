package usecases

import (
	"context"

	"jwemanager/pkg/app/interfaces"
	"jwemanager/pkg/domain/usecases"
	valueObjects "jwemanager/pkg/domain/value_objects"
)

type createKeyUseCase struct {
	logger  interfaces.ILogger
	keyRepo interfaces.IKeyRepository
	guidGen interfaces.IGuidGenerator
	keyGen  interfaces.IKeyGenerator
}

func (pst createKeyUseCase) Execute(ctx context.Context, key valueObjects.Key, timeToExpiration int) (valueObjects.Key, error) {
	key.KeyID = pst.guidGen.V4()

	rsaKeys, err := pst.keyGen.GenerateKey()
	if err != nil {
		return valueObjects.Key{}, err
	}

	key.PriKey = rsaKeys
	key.PubKey = &rsaKeys.PublicKey

	result, err := pst.keyRepo.CreateKey(ctx, key, timeToExpiration)
	if err != nil {
		return valueObjects.Key{}, err
	}

	return result, nil
}

func NewCreateKeyUseCase(
	logger interfaces.ILogger,
	keyRepo interfaces.IKeyRepository,
	guidGen interfaces.IGuidGenerator,
	keyGen interfaces.IKeyGenerator,
) usecases.ICreateKeyUseCase {
	return createKeyUseCase{logger, keyRepo, guidGen, keyGen}
}
