package crypto

import (
	"crypto/rsa"

	"github.com/stretchr/testify/mock"
)

type CryptoSpy struct {
	mock.Mock
}

func (pst CryptoSpy) Encrypt(pubKey *rsa.PublicKey, data map[string]interface{}) ([]byte, error) {
	args := pst.Called(pubKey, data)

	return args.Get(0).([]byte), args.Error(1)
}

func (pst CryptoSpy) Decrypt(pubprivKey *rsa.PrivateKey, data []byte) (map[string]interface{}, error) {
	args := pst.Called(pubprivKey, data)

	return args.Get(0).(map[string]interface{}), args.Error(1)
}

func NewCryptoSpy() *CryptoSpy {
	return new(CryptoSpy)
}
