package viewmodels

import "jwemanager/pkg/domain/dtos"

type CreateKeyViewModel struct {
	UserID string `json:"user_id" validate:"required, uuidv4"`
	KeyID  string `json:"key_id" validate:"required, uuidv4"`
}

func (pst CreateKeyViewModel) ToDto() dtos.Key {
	return dtos.Key{
		UserID: pst.UserID,
		KeyID:  pst.KeyID,
	}
}
