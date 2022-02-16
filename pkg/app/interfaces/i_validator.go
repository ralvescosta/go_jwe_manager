package interfaces

import (
	"jwemanager/pkg/domain/dtos"
)

type IValidator interface {
	ValidateStruct(data interface{}) []dtos.ValidateResult
}
