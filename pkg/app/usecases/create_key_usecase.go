package usecases

import (
	"context"
	"jwemanager/pkg/app/interfaces"
	"jwemanager/pkg/domain/usecases"
	valueObjects "jwemanager/pkg/domain/value_objects"
	"time"
)

type createKeyUseCase struct {
	logger  interfaces.ILogger
	keyRepo interfaces.IKeyRepository
	keyGen  interfaces.IKeyGenerator
}

func (pst createKeyUseCase) Execute(ctx context.Context, key valueObjects.Key) (valueObjects.Key, error) {
	key.ExpiredAt = time.Now().Add(time.Hour * 720)
	rsaKeys, err := pst.keyGen.GenerateKey()
	if err != nil {
		return valueObjects.Key{}, err
	}

	key.PriKey = rsaKeys
	key.PubKey = &rsaKeys.PublicKey

	result, err := pst.keyRepo.CreateKey(ctx, key)
	if err != nil {
		return valueObjects.Key{}, err
	}

	return result, nil
}

func NewCreateKeyUseCase(
	logger interfaces.ILogger,
	keyRepo interfaces.IKeyRepository,
	keyGen interfaces.IKeyGenerator,
) usecases.ICreateKeyUseCase {
	return createKeyUseCase{logger, keyRepo, keyGen}
}
