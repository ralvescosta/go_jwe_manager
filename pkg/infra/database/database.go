package database

import (
	"context"
	"fmt"
	"jwemanager/pkg/app/interfaces"
	"time"

	"github.com/go-redis/redis/v8"
)

func Connection(logger interfaces.ILogger, shotdown chan bool) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		logger.Error(fmt.Sprintf("[Database::Connection] - Redis Connection failure : %s", err.Error()))
		return nil, err
	}

	go signalShotdown(rdb, logger, shotdown)

	return rdb, nil
}

func signalShotdown(rdb *redis.Client, logger interfaces.ILogger, shotdown chan bool) {
	time.Sleep(time.Second * 10)
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		logger.Error(fmt.Sprintf("[Database::Connection] - Redis Connection failure : %s", err.Error()))
		shotdown <- true
	}
}
