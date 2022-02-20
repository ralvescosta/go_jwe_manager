package usecases

import (
	"context"
	"jwemanager/pkg/app/interfaces"
	"jwemanager/pkg/domain/usecases"
	valueObjects "jwemanager/pkg/domain/value_objects"
)

type getKeyUseCase struct {
	logger  interfaces.ILogger
	keyRepo interfaces.IKeyRepository
}

func (pst getKeyUseCase) GetKey(ctx context.Context, userID, keyID string) (valueObjects.Key, error) {
	return pst.keyRepo.GetKeyByID(ctx, userID, keyID)
}

func NewGetKeyUseCase(logger interfaces.ILogger, keyRepo interfaces.IKeyRepository) usecases.IGetKeyUseCase {
	return getKeyUseCase{
		logger,
		keyRepo,
	}
}
