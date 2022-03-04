package keyGenerator

import (
	"crypto/rsa"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GenerateKey(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := NewKeyGeneratorSpy()

		sut.On("GenerateKey").Return(&rsa.PrivateKey{}, nil)

		result, err := sut.GenerateKey()

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}
