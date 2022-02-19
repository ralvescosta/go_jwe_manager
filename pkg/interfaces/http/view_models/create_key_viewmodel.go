package viewmodels

import (
	"crypto/x509"
	"encoding/base64"
	valueObjects "jwemanager/pkg/domain/value_objects"
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

func (pst CreateKeyViewModel) ToDto() valueObjects.Key {
	return valueObjects.Key{
		UserID: pst.UserID,
		KeyID:  pst.KeyID,
	}
}

func NewCreatedKeyViewModel(key valueObjects.Key) CreatedKeyViewModel {
	return CreatedKeyViewModel{
		UserID:    key.UserID,
		KeyID:     key.KeyID,
		PubKey:    base64.RawStdEncoding.EncodeToString(x509.MarshalPKCS1PublicKey(key.PubKey)),
		ExpiredAt: key.ExpiredAt.Format(time.RFC3339),
	}
}
