package viewmodels

import (
	"jwemanager/pkg/domain/dtos"
	"time"
)

type CreateKeyViewModel struct {
	UserID string `json:"user_id" validate:"required,uuid4"`
	KeyID  string `json:"key_id" validate:"required,uuid4"`
}

type CreatedKeyViewModel struct {
	UserID    string `json:"user_id"`
	KeyID     string `json:"key_id"`
	PubKey    string `json:"pub_key"`
	ExpiredAt string `json:"expired_at"`
}

func (pst CreateKeyViewModel) ToDto() dtos.Key {
	return dtos.Key{
		UserID: pst.UserID,
		KeyID:  pst.KeyID,
	}
}

func NewCreatedKeyViewModel(key dtos.Key) CreatedKeyViewModel {
	return CreatedKeyViewModel{
		UserID:    key.UserID,
		KeyID:     key.KeyID,
		PubKey:    key.PubKey,
		ExpiredAt: key.ExpiredAt.Format(time.RFC3339),
	}
}
