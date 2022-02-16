package usecases

import "jwemanager/pkg/domain/dtos"

type ICreateKeyUseCase interface {
	Execute(key dtos.Key) (dtos.Key, error)
}
