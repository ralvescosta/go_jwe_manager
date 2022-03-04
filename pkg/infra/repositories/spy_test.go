package repositories

import (
	"context"
	"testing"

	valueObjects "jwemanager/pkg/domain/value_objects"

	"github.com/stretchr/testify/assert"
)

func Test_CreateKey(t *testing.T) {
	t.Run("execute correctly", func(t *testing.T) {
		sut := NewKeyRepositorySpy()
		ctx := context.Background()
		key := valueObjects.Key{}
		timeToExpiration := 0

		sut.On("CreateKey", ctx, key, timeToExpiration).Return(valueObjects.Key{}, nil)

		result, err := sut.CreateKey(ctx, key, timeToExpiration)

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func Test_GetKeyByID(t *testing.T) {
	t.Run("execute correctly", func(t *testing.T) {
		sut := NewKeyRepositorySpy()
		ctx := context.Background()
		userID := "some_user"
		keyID := "some_key"

		sut.On("GetKeyByID", ctx, userID, keyID).Return(valueObjects.Key{}, nil)

		result, err := sut.GetKeyByID(ctx, userID, keyID)

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}
