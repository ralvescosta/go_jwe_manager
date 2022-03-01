package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"jwemanager/pkg/app/interfaces"
	"jwemanager/pkg/infra/logger"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Encrypt(t *testing.T) {
	t.Run("should execute encrypt correctly", func(t *testing.T) {
		sut := makeCryptoSutRtn()

		result, err := sut.crypto.Encrypt(&sut.privateKeyMocked.PublicKey, sut.dataMocked)

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func Test_Decrypt(t *testing.T) {
	t.Run("should decrypt correctly", func(t *testing.T) {
		sut := makeCryptoSutRtn()
		encrypted, _ := sut.crypto.Encrypt(&sut.privateKeyMocked.PublicKey, sut.dataMocked)

		result, err := sut.crypto.Decrypt(sut.privateKeyMocked, encrypted)

		assert.NoError(t, err)
		assert.Equal(t, sut.dataMocked, result)
	})
}

type cryptoSutRtn struct {
	crypto           interfaces.ICrypto
	logger           *logger.LoggerSpy
	privateKeyMocked *rsa.PrivateKey
	dataMocked       map[string]interface{}
}

func makeCryptoSutRtn() cryptoSutRtn {
	logger := logger.NewLoggerSpy()
	crypto := NewCrypto(logger)

	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	data := make(map[string]interface{})

	return cryptoSutRtn{crypto, logger, privateKey, data}
}
