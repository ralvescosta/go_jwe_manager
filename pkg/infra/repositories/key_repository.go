package repositories

import (
	"context"
	"fmt"
	"time"

	"jwemanager/pkg/app/errors"
	"jwemanager/pkg/app/interfaces"
	valueObjects "jwemanager/pkg/domain/value_objects"
	"jwemanager/pkg/infra/database/models"

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
		return valueObjects.Key{}, errors.NewInternalError(err.Error())
	}

	return key, nil
}

func (pst keyRepository) GetKeyByID(ctx context.Context, userID, keyID string) (valueObjects.Key, error) {
	result := pst.rdb.Get(ctx, getRedisKeyByIDs(userID, keyID))
	if err := result.Err(); err != nil {
		pst.logger.Error(fmt.Sprintf("[KeyRepository::GetKeyByID] - Error: %s", err.Error()))
		return valueObjects.Key{}, errors.NewInternalError(err.Error())
	}

	var keyModel models.KeyModel
	if err := result.Scan(&keyModel); err != nil {
		pst.logger.Error(fmt.Sprintf("[KeyRepository::GetKeyByID] - Error: %s", err.Error()))
		return valueObjects.Key{}, errors.NewInternalError(err.Error())
	}

	vo, err := keyModel.ToValueObject()
	if err != nil {
		pst.logger.Error(fmt.Sprintf("[KeyRepository::GetKeyByID] - Error: %s", err.Error()))
		return valueObjects.Key{}, errors.NewInternalError(err.Error())
	}

	return vo, nil
}

func NewKeyRepository(logger interfaces.ILogger, guidGen interfaces.IGuidGenerator, rdb *redis.Client) interfaces.IKeyRepository {
	return keyRepository{logger, guidGen, rdb}
}
