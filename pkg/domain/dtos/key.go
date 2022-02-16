package dtos

import "time"

type Key struct {
	Id        string
	PubKey    string
	PriKey    string
	CreatedAt time.Time
	ExpiredAt time.Time
}
