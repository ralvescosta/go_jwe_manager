package repositories

import (
	"context"
	"fmt"
	"jwemanager/pkg/app/interfaces"
	"jwemanager/pkg/domain/dtos"

	"github.com/go-redis/redis/v8"
)

type keyRepository struct {
	logger interfaces.ILogger
	rdb    *redis.Client
}

func (pst keyRepository) CreateKey(ctx context.Context, key dtos.Key) (dtos.Key, error) {
	key.ID = "ai pai."

	if err := pst.rdb.Set(ctx, getRedisKeyByKey(key), key, 0).Err(); err != nil {
		pst.logger.Error(fmt.Sprintf("[KeyRepository::CreateKey] - Error: %s", err.Error()))
		return dtos.Key{}, err
	}

	return key, nil
}

func (pst keyRepository) GetKeyByID(ctx context.Context, userID, keyID string) (dtos.Key, error) {
	result := pst.rdb.Get(ctx, getRedisKeyByIDs(userID, keyID))
	if err := result.Err(); err != nil {
		pst.logger.Error(fmt.Sprintf("[KeyRepository::GetKeyByID] - Error: %s", err.Error()))
		return dtos.Key{}, err
	}

	var key dtos.Key
	if err := result.Scan(&key); err != nil {
		pst.logger.Error(fmt.Sprintf("[KeyRepository::GetKeyByID] - Error: %s", err.Error()))
		return dtos.Key{}, err
	}

	return key, nil
}

func NewKeyRepository(logger interfaces.ILogger, rdb *redis.Client) interfaces.IKeyRepository {
	return keyRepository{logger, rdb}
}
