package viewmodels

import (
	"crypto/rand"
	"crypto/rsa"
	"testing"

	valueObjects "jwemanager/pkg/domain/value_objects"

	"github.com/stretchr/testify/assert"
)

func Test_KeyViewModel_ToValueObject(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := CreateKeyViewModel{
			UserID: "user_id",
		}

		vo := sut.ToValueObject()

		assert.Equal(t, "user_id", vo.UserID)
		assert.IsType(t, valueObjects.Key{}, vo)
	})
}

func Test_NewResultViewModel(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		priv, _ := rsa.GenerateKey(rand.Reader, 2048)
		keyMocked := valueObjects.Key{
			UserID: "user_id",
			KeyID:  "key_id",
			PubKey: &priv.PublicKey,
		}

		sut := NewResultKeyViewModel(keyMocked)

		assert.Equal(t, "user_id", sut.UserID)
		assert.Equal(t, "key_id", sut.KeyID)
		assert.IsType(t, ResultKeyViewModel{}, sut)
	})
}
