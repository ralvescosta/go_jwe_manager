package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"jwemanager/pkg/app/interfaces"
	"jwemanager/pkg/infra/logger"
	"testing"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwe"
	"github.com/stretchr/testify/assert"
)

func Test_Encrypt(t *testing.T) {
	t.Run("should execute encrypt correctly", func(t *testing.T) {
		sut := makeCryptoSutRtn()

		result, err := sut.crypto.Encrypt(&sut.privateKeyMocked.PublicKey, sut.dataMocked)

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should return error if some error occur in marshal", func(t *testing.T) {
		sut := makeCryptoSutRtn()

		marshal = func(v interface{}) ([]byte, error) {
			return nil, errors.New("some error")
		}
		_, err := sut.crypto.Encrypt(&sut.privateKeyMocked.PublicKey, nil)

		assert.Error(t, err)
	})

	t.Run("should return error if some error occur in encrypt", func(t *testing.T) {
		sut := makeCryptoSutRtn()

		encrypt = func(payload []byte, keyalg jwa.KeyEncryptionAlgorithm, key interface{}, contentalg jwa.ContentEncryptionAlgorithm, compressalg jwa.CompressionAlgorithm, options ...jwe.EncryptOption) ([]byte, error) {
			return nil, errors.New("some error")
		}
		_, err := sut.crypto.Encrypt(&sut.privateKeyMocked.PublicKey, nil)

		assert.Error(t, err)
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

	t.Run("should error if some error occur in decrypt", func(t *testing.T) {
		sut := makeCryptoSutRtn()
		encrypted, _ := sut.crypto.Encrypt(&sut.privateKeyMocked.PublicKey, sut.dataMocked)

		decrypt = func(buf []byte, alg jwa.KeyEncryptionAlgorithm, key interface{}, options ...jwe.DecryptOption) ([]byte, error) {
			return nil, errors.New("some error")
		}
		_, err := sut.crypto.Decrypt(sut.privateKeyMocked, encrypted)

		assert.Error(t, err)
	})

	t.Run("should error if some error occur in unmarshal", func(t *testing.T) {
		sut := makeCryptoSutRtn()
		encrypted, _ := sut.crypto.Encrypt(&sut.privateKeyMocked.PublicKey, sut.dataMocked)

		decrypt = func(buf []byte, alg jwa.KeyEncryptionAlgorithm, key interface{}, options ...jwe.DecryptOption) ([]byte, error) {
			return nil, nil
		}
		_, err := sut.crypto.Decrypt(sut.privateKeyMocked, encrypted)

		assert.Error(t, err)
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
