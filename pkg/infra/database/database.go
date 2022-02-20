package database

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"jwemanager/pkg/app/errors"
	"jwemanager/pkg/app/interfaces"

	"github.com/go-redis/redis/v8"
)

func Connection(logger interfaces.ILogger, shotdown chan bool) (*redis.Client, error) {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	password := os.Getenv("REDIS_PASSWORD")

	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		return nil, errors.NewInternalError(err.Error())
	}

	secondsToSleep, err := strconv.Atoi(os.Getenv("REDIS_SECONDS_TO_PING"))
	if err != nil {
		return nil, errors.NewInternalError(err.Error())
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       db,
	})

	_, err = rdb.Ping(context.Background()).Result()
	if err != nil {
		logger.Error(fmt.Sprintf("[Database::Connection] - Redis Connection failure : %s", err.Error()))
		return nil, errors.NewInternalError(err.Error())
	}

	go signalShotdown(rdb, logger, secondsToSleep, shotdown)

	return rdb, nil
}

func signalShotdown(rdb *redis.Client, logger interfaces.ILogger, secondsToSleep int, shotdown chan bool) {
	time.Sleep(time.Duration(secondsToSleep) * time.Second)
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		logger.Error(fmt.Sprintf("[Database::Connection] - Redis Connection failure : %s", err.Error()))
		shotdown <- true
	}
}
