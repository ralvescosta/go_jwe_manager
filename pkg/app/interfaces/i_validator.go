package interfaces

import (
	valueObjects "jwemanager/pkg/domain/value_objects"
)

type IValidator interface {
	ValidateStruct(data interface{}) []valueObjects.ValidateResult
}
