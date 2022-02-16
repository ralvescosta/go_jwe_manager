package validator

import (
	"fmt"

	"jwemanager/pkg/app/interfaces"
	"jwemanager/pkg/domain/dtos"

	"github.com/go-playground/validator/v10"
)

type vAlidator struct{}

func (vAlidator) ValidateStruct(m interface{}) []dtos.ValidateResult {
	v := validator.New()
	err := v.Struct(m)

	if err == nil {
		return nil
	}

	validationErrors := err.(validator.ValidationErrors)

	var validatedErros []dtos.ValidateResult
	for _, validationErr := range validationErrors {
		validatedErros = append(validatedErros, dtos.ValidateResult{
			IsValid: false,
			Field:   validationErr.Field(),
			Message: fmt.Sprintf("%s is invalid", validationErr.Field()),
		})
	}

	return validatedErros
}

func NewValidator() interfaces.IValidator {
	return vAlidator{}
}
