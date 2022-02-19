package models

import (
	"crypto/x509"
	"encoding/json"
	"time"

	valueObjects "jwemanager/pkg/domain/value_objects"
)

type KeyModel struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	KeyID     string    `json:"key_id"`
	PubKey    []byte    `json:"pub_key"`
	PriKey    []byte    `json:"priv_key"`
	CreatedAt time.Time `json:"created_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func ToKeyModel(vo valueObjects.Key) KeyModel {
	return KeyModel{
		ID:        vo.ID,
		UserID:    vo.UserID,
		KeyID:     vo.KeyID,
		PubKey:    x509.MarshalPKCS1PublicKey(vo.PubKey),
		PriKey:    x509.MarshalPKCS1PrivateKey(vo.PriKey),
		CreatedAt: vo.CreatedAt,
		ExpiredAt: vo.ExpiredAt,
	}
}

func (pst KeyModel) MarshalBinary() ([]byte, error) {
	return json.Marshal(pst)
}
