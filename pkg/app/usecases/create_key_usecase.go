package usecases

import (
	"jwemanager/pkg/domain/dtos"
	"jwemanager/pkg/domain/usecases"
)

type createKeyUseCase struct{}

func (createKeyUseCase) Execute(key dtos.Key) (dtos.Key, error) {
	return dtos.Key{}, nil
}

func NewCreateKeyUseCase() usecases.ICreateKeyUseCase {
	return createKeyUseCase{}
}
