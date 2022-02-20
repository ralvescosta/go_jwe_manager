package viewmodels

import (
	"crypto/x509"
	"encoding/pem"
	valueObjects "jwemanager/pkg/domain/value_objects"
	"time"
)

type CreateKeyViewModel struct {
	UserID string `json:"user_id" validate:"required,uuid4"`
}

type ResultKeyViewModel struct {
	UserID    string `json:"user_id"`
	KeyID     string `json:"key_id"`
	PubKey    string `json:"pub_key"`
	ExpiredAt string `json:"expired_at"`
}

func (pst CreateKeyViewModel) ToDto() valueObjects.Key {
	return valueObjects.Key{
		UserID: pst.UserID,
	}
}

func NewResultKeyViewModel(key valueObjects.Key) ResultKeyViewModel {
	return ResultKeyViewModel{
		UserID: key.UserID,
		KeyID:  key.KeyID,
		PubKey: string(pem.EncodeToMemory(&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(key.PubKey),
		})),
		ExpiredAt: key.ExpiredAt.Format(time.RFC3339),
	}
}
