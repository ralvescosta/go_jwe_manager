package repositories

import (
	"fmt"
	"jwemanager/pkg/domain/dtos"
)

func getRedisKeyByKey(key dtos.Key) string {
	return fmt.Sprintf("%s:%s", key.UserID, key.KeyID)
}

func getRedisKeyByIDs(userID, keyID string) string {
	return fmt.Sprintf("%s:%s", userID, keyID)
}
