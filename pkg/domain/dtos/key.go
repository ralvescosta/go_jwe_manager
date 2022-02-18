package dtos

import (
	"encoding/json"
	"time"
)

type Key struct {
	ID        string
	UserID    string
	KeyID     string
	PubKey    string
	PriKey    string
	CreatedAt time.Time
	ExpiredAt time.Time
}

func (pst Key) MarshalBinary() ([]byte, error) {
	return json.Marshal(pst)
}
