package repositories

import (
	"context"
	"fmt"
	"jwemanager/pkg/app/interfaces"
	valueObjects "jwemanager/pkg/domain/value_objects"
	"jwemanager/pkg/infra/database/models"
	"time"

	"github.com/go-redis/redis/v8"
)

type keyRepository struct {
	logger  interfaces.ILogger
	guidGen interfaces.IGuidGenerator
	rdb     *redis.Client
}

func (pst keyRepository) CreateKey(ctx context.Context, key valueObjects.Key) (valueObjects.Key, error) {
	key.ID = pst.guidGen.V4()
	key.CreatedAt = time.Now()

	if err := pst.rdb.Set(ctx, getRedisKeyByKey(key), models.ToKeyModel(key), 0).Err(); err != nil {
		pst.logger.Error(fmt.Sprintf("[KeyRepository::CreateKey] - Error: %s", err.Error()))
		return valueObjects.Key{}, err
	}

	return key, nil
}

func (pst keyRepository) GetKeyByID(ctx context.Context, userID, keyID string) (valueObjects.Key, error) {
	result := pst.rdb.Get(ctx, getRedisKeyByIDs(userID, keyID))
	if err := result.Err(); err != nil {
		pst.logger.Error(fmt.Sprintf("[KeyRepository::GetKeyByID] - Error: %s", err.Error()))
		return valueObjects.Key{}, err
	}

	var key valueObjects.Key
	if err := result.Scan(&key); err != nil {
		pst.logger.Error(fmt.Sprintf("[KeyRepository::GetKeyByID] - Error: %s", err.Error()))
		return valueObjects.Key{}, err
	}

	return key, nil
}

func NewKeyRepository(logger interfaces.ILogger, guidGen interfaces.IGuidGenerator, rdb *redis.Client) interfaces.IKeyRepository {
	return keyRepository{logger, guidGen, rdb}
}
