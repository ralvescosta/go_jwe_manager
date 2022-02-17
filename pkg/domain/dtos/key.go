package dtos

import "time"

type Key struct {
	ID        string
	UserID    string
	KeyID     string
	PubKey    string
	PriKey    string
	CreatedAt time.Time
	ExpiredAt time.Time
}
