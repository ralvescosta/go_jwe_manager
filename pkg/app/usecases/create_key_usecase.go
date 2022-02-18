package usecases

import (
	"context"
	"jwemanager/pkg/app/interfaces"
	"jwemanager/pkg/domain/dtos"
	"jwemanager/pkg/domain/usecases"
	"time"
)

type createKeyUseCase struct {
	logger  interfaces.ILogger
	keyRepo interfaces.IKeyRepository
}

func (pst createKeyUseCase) Execute(ctx context.Context, key dtos.Key) (dtos.Key, error) {
	key.ExpiredAt = time.Now().Add(time.Hour * 720)

	result, err := pst.keyRepo.CreateKey(ctx, key)
	if err != nil {
		return dtos.Key{}, err
	}

	return result, nil
}

func NewCreateKeyUseCase(
	logger interfaces.ILogger,
	keyRepo interfaces.IKeyRepository,
) usecases.ICreateKeyUseCase {
	return createKeyUseCase{logger, keyRepo}
}
