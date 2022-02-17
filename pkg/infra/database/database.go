package database

import (
	"context"
	"fmt"
	"jwemanager/pkg/app/interfaces"

	"github.com/go-redis/redis/v8"
)

func Connection(logger interfaces.ILogger) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	result := rdb.Ping(context.Background())
	_, err := result.Result()
	if err != nil {
		logger.Error(fmt.Sprintf("[Database::Connection] - Redis Connection failure : %s", err.Error()))
		return nil, err
	}

	return rdb, nil
}
