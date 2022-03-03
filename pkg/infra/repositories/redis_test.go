package repositories

import (
	"fmt"
	"testing"

	valueObjects "jwemanager/pkg/domain/value_objects"

	"github.com/stretchr/testify/assert"
)

func Test_GetRedisKeyByKey(t *testing.T) {
	t.Run("execute correctly", func(t *testing.T) {
		key := valueObjects.Key{
			UserID: "some_user",
			KeyID:  "some_key",
		}

		assert.Equal(t, fmt.Sprintf("%s:%s", key.UserID, key.KeyID), getRedisKeyByKey(key))
	})
}

func Test_GetRedisKeyByIDs(t *testing.T) {
	t.Run("execute correctly", func(t *testing.T) {
		userID := "some_user"
		keyID := "some_key"

		assert.Equal(t, fmt.Sprintf("%s:%s", userID, keyID), getRedisKeyByIDs(userID, keyID))
	})
}
