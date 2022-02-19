package repositories

import (
	"fmt"
	valueObjects "jwemanager/pkg/domain/value_objects"
)

func getRedisKeyByKey(key valueObjects.Key) string {
	return fmt.Sprintf("%s:%s", key.UserID, key.KeyID)
}

func getRedisKeyByIDs(userID, keyID string) string {
	return fmt.Sprintf("%s:%s", userID, keyID)
}
