package keyGenerator

import (
	"crypto/rsa"
	"errors"
	"io"
	"testing"

	"jwemanager/pkg/app/interfaces"
	"jwemanager/pkg/infra/logger"

	"github.com/stretchr/testify/assert"
)

func Test_KeyGen(t *testing.T) {
	t.Run("should execute GenereteKey correctly", func(t *testing.T) {
		sut := makeKeyGenSutRtn()

		genRSAKey = func(random io.Reader, bits int) (*rsa.PrivateKey, error) {
			return &rsa.PrivateKey{}, nil
		}

		result, err := sut.keyGen.GenerateKey()

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should execute GenereteKey and return error when some error occur", func(t *testing.T) {
		sut := makeKeyGenSutRtn()

		genRSAKey = func(random io.Reader, bits int) (*rsa.PrivateKey, error) {
			return nil, errors.New("some error")
		}

		result, err := sut.keyGen.GenerateKey()

		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

type keyGenSutRtn struct {
	keyGen interfaces.IKeyGenerator
	logger *logger.LoggerSpy
}

func makeKeyGenSutRtn() keyGenSutRtn {
	logger := logger.NewLoggerSpy()
	keyGen := NewKeyGenerator(logger)

	return keyGenSutRtn{keyGen, logger}
}
