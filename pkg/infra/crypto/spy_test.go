package crypto

import (
	"crypto/rsa"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_EncryptSpy(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := NewCryptoSpy()

		pubKey := &rsa.PublicKey{}
		data := make(map[string]interface{})

		sut.On("Encrypt", pubKey, data).Return([]byte(""), nil)

		result, err := sut.Encrypt(pubKey, data)

		assert.NoError(t, err)
		assert.Equal(t, []byte(""), result)
	})
}

func Test_DecryptSpy(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := NewCryptoSpy()

		privKey := &rsa.PrivateKey{}
		data := []byte("")
		expectedResult := make(map[string]interface{})

		sut.On("Decrypt", privKey, data).Return(expectedResult, nil)

		result, err := sut.Decrypt(privKey, data)

		assert.NoError(t, err)
		assert.Equal(t, expectedResult, result)
	})
}
