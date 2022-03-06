package handlers

import (
	httpServer "jwemanager/pkg/infra/http_server"

	"github.com/stretchr/testify/mock"
)

//
type CryptHandlersSpy struct {
	mock.Mock
}

func (pst CryptHandlersSpy) Encrypt(httpRequest httpServer.HttpRequest) httpServer.HttpResponse {
	args := pst.Called(httpRequest)

	return args.Get(0).(httpServer.HttpResponse)
}

func (pst CryptHandlersSpy) Decrypt(httpRequest httpServer.HttpRequest) httpServer.HttpResponse {
	args := pst.Called(httpRequest)

	return args.Get(0).(httpServer.HttpResponse)
}

func NewCryptoHandlersSpy() *CryptHandlersSpy {
	return new(CryptHandlersSpy)
}

//
type HealthHandlersSpy struct {
	mock.Mock
}

func (pst HealthHandlersSpy) Check(httpRequest httpServer.HttpRequest) httpServer.HttpResponse {
	args := pst.Called(httpRequest)

	return args.Get(0).(httpServer.HttpResponse)
}

func NewHealthHandlerSpy() *HealthHandlersSpy {
	return new(HealthHandlersSpy)
}

//
type KeysHandlersSpy struct {
	mock.Mock
}

func (pst KeysHandlersSpy) Create(httpRequest httpServer.HttpRequest) httpServer.HttpResponse {
	args := pst.Called(httpRequest)

	return args.Get(0).(httpServer.HttpResponse)
}

func (pst KeysHandlersSpy) FindOne(httpRequest httpServer.HttpRequest) httpServer.HttpResponse {
	args := pst.Called(httpRequest)

	return args.Get(0).(httpServer.HttpResponse)
}

func NewKeysHandlersSpy() *KeysHandlersSpy {
	return new(KeysHandlersSpy)
}
