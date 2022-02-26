package keyGenerator

import (
	"crypto/rsa"

	"github.com/stretchr/testify/mock"
)

type KeyGeneratorSpy struct {
	mock.Mock
}

func (pst KeyGeneratorSpy) GenerateKey() (*rsa.PrivateKey, error) {
	args := pst.Called()

	return args.Get(0).(*rsa.PrivateKey), args.Error(1)
}

func NewKeyGeneratorSpy() *KeyGeneratorSpy {
	return new(KeyGeneratorSpy)
}
