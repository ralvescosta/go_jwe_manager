package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"errors"
	"testing"

	"jwemanager/pkg/app/interfaces"
	"jwemanager/pkg/infra/logger"

	appErrors "jwemanager/pkg/app/errors"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwe"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
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
		encrypt = jwe.Encrypt
		marshal = func(v interface{}) ([]byte, error) {
			return nil, errors.New("some error")
		}
		sut.logger.On("Error", "[Crypto::Encrypt] Marshal Error: some error", []zap.Field(nil))

		_, err := sut.crypto.Encrypt(&sut.privateKeyMocked.PublicKey, nil)

		assert.Error(t, err)
		sut.logger.AssertExpectations(t)
	})

	t.Run("should return error if some error occur in encrypt", func(t *testing.T) {
		sut := makeCryptoSutRtn()
		marshal = json.Marshal
		encrypt = func(payload []byte, keyalg jwa.KeyEncryptionAlgorithm, key interface{}, contentalg jwa.ContentEncryptionAlgorithm, compressalg jwa.CompressionAlgorithm, options ...jwe.EncryptOption) ([]byte, error) {
			return nil, errors.New("some error")
		}
		sut.logger.On("Error", "[Crypto::Encrypt] JWE Encrypt Error: some error", []zap.Field(nil))

		_, err := sut.crypto.Encrypt(&sut.privateKeyMocked.PublicKey, nil)

		assert.Error(t, err)
		assert.IsType(t, appErrors.InternalError{}, err)
		sut.logger.AssertExpectations(t)
	})
}

func Test_Decrypt(t *testing.T) {
	t.Run("should decrypt correctly", func(t *testing.T) {
		sut := makeCryptoSutRtn()
		encrypt = jwe.Encrypt
		marshal = json.Marshal
		encrypted, _ := sut.crypto.Encrypt(&sut.privateKeyMocked.PublicKey, sut.dataMocked)
		decrypt = jwe.Decrypt

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
		sut.logger.On("Error", "[Crypto::Encrypt] JWE Decrypt Error: some error", []zap.Field(nil))

		_, err := sut.crypto.Decrypt(sut.privateKeyMocked, encrypted)

		assert.Error(t, err)
		sut.logger.AssertExpectations(t)
	})

	t.Run("should error if some error occur in unmarshal", func(t *testing.T) {
		sut := makeCryptoSutRtn()
		encrypted, _ := sut.crypto.Encrypt(&sut.privateKeyMocked.PublicKey, sut.dataMocked)

		decrypt = func(buf []byte, alg jwa.KeyEncryptionAlgorithm, key interface{}, options ...jwe.DecryptOption) ([]byte, error) {
			return nil, nil
		}
		sut.logger.On("Error", "[Crypto::Encrypt] Decrypted Data Unmarshaler Error: unexpected end of JSON input", []zap.Field(nil))

		_, err := sut.crypto.Decrypt(sut.privateKeyMocked, encrypted)

		assert.Error(t, err)
		sut.logger.AssertExpectations(t)
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
