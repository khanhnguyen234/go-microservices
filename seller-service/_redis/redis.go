package _redis

import (
	"context"
	"github.com/go-redis/redis/v7"
	"khanhnguyen234/seller-service/common"
	"os"
	"strconv"
)

var ctx = context.Background()
var RDB *redis.Client

func ConnectRedis() *redis.Client {
	db, _ := strconv.Atoi(os.Getenv("REDIS_DATABASE"))

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),      // use default Addr
		Password: os.Getenv("REDIS_PASSWORD"), // no password set
		DB:       db,                          // use default DB
	})

	_, err := rdb.Ping().Result()
	common.LogStatus(err, "Connect Redis")

	RDB = rdb
	return rdb
}

func GetRedis() *redis.Client {
	return RDB
}

func Set(key string, value string) error {
	if err := RDB.Set(key, value, 0).Err(); err != nil {
		return err
	}
	return nil
}

func Get(key string) (string, bool) {
	value, err := RDB.Get(key).Result()

	if err != nil {
		return "", true
	}

	return value, false
}
