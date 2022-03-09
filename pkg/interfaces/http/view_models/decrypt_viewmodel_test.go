package viewmodels

import (
	"testing"

	valueObjects "jwemanager/pkg/domain/value_objects"

	"github.com/stretchr/testify/assert"
)

func Test_DecryptViewModel_ToValueObject(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := DecryptViewModel{
			UserID:        "user_id",
			KeyID:         "key_id",
			EncryptedData: "encrypted_data",
		}

		vo := sut.ToValueObject()

		assert.IsType(t, valueObjects.DecryptValueObject{}, vo)
		assert.Equal(t, "user_id", vo.UserID)
		assert.Equal(t, "key_id", vo.KeyID)
		assert.Equal(t, []byte("encrypted_data"), vo.EncryptedData)
	})
}

func Test_ToDecryptedViewModel(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		decryptedMocked := valueObjects.DecryptedValueObject{
			Data: make(map[string]interface{}),
		}

		sut := NewDecryptedViewModel(decryptedMocked)

		assert.IsType(t, DecryptedViewModel{}, sut)
	})
}
