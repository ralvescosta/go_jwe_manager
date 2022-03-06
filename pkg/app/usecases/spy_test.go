package usecases

import (
	"context"
	valueObjects "jwemanager/pkg/domain/value_objects"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateKey_Execute(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := NewCreateKeyUseCaseSpy()

		ctx := context.Background()
		key := valueObjects.Key{}
		timeToExpiration := 0

		sut.On("Execute", ctx, key, timeToExpiration).Return(valueObjects.Key{}, nil)

		result, err := sut.Execute(ctx, key, timeToExpiration)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		sut.AssertExpectations(t)
	})
}

func Test_Decrypt_Decrypt(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := NewDecryptUseCaseSpy()

		ctx := context.Background()
		data := valueObjects.DecryptValueObject{}
		decryptedData := valueObjects.DecryptedValueObject{}

		sut.On("Decrypt", ctx, data).Return(decryptedData, nil)

		result, err := sut.Decrypt(ctx, data)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		sut.AssertExpectations(t)
	})
}

func Test_Encrypt_Encrypt(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := NewEncryptUseCaseSpy()

		ctx := context.Background()
		data := valueObjects.EncryptValueObject{}
		encryptedData := valueObjects.EncryptedValueObject{}

		sut.On("Encrypt", ctx, data).Return(encryptedData, nil)

		result, err := sut.Encrypt(ctx, data)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		sut.AssertExpectations(t)
	})
}

func Test_GetKey_GetKey(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := NewGetKeyUseCaseSpy()

		ctx := context.Background()
		userID := "user_id"
		keyID := "key_id"
		key := valueObjects.Key{}

		sut.On("GetKey", ctx, userID, keyID).Return(key, nil)

		result, err := sut.GetKey(ctx, userID, keyID)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		sut.AssertExpectations(t)
	})
}
