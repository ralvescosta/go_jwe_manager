package viewmodels

import (
	"testing"

	valueObjects "jwemanager/pkg/domain/value_objects"

	"github.com/stretchr/testify/assert"
)

func Test_EncryptViewModel_ToValueObject(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := EncryptViewModel{
			UserID: "user_id",
			KeyID:  "key_id",
			Data:   make(map[string]interface{}),
		}

		vo := sut.ToValueObject()

		assert.IsType(t, valueObjects.EncryptValueObject{}, vo)
		assert.Equal(t, "user_id", vo.UserID)
		assert.Equal(t, "key_id", vo.KeyID)
	})
}

func Test_NewEncryptViewModel(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		encryptedMocked := valueObjects.EncryptedValueObject{
			EncryptedData: "encrypted_data",
		}

		sut := NewEncryptedViewModel(encryptedMocked)

		assert.IsType(t, EncryptedViewModel{}, sut)
		assert.Equal(t, "encrypted_data", sut.EncryptedData)
	})
}
