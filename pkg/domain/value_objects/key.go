package valueObjects

import (
	"crypto/rsa"
	"time"
)

type Key struct {
	ID        string
	UserID    string
	KeyID     string
	PubKey    *rsa.PublicKey
	PriKey    *rsa.PrivateKey
	CreatedAt time.Time
	ExpiredAt time.Time
}
