package viewmodels

import "jwemanager/pkg/domain/dtos"

type CreateKeyViewModel struct {
	Id string `json:"id" validate:"required, uuidv4"`
}

func (pst CreateKeyViewModel) ToDto() dtos.Key {
	return dtos.Key{
		Id: pst.Id,
	}
}
