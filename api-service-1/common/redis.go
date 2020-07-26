package common

import (
  "github.com/go-redis/redis/v7"
)

var RDB *redis.Client

func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // use default Addr
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	RDB = rdb
	return rdb
}

func Redis() *redis.Client {
	return RDB
}

func SetRedis(key string, value string) error {
	if err := RDB.Set(key, value, 0).Err(); err != nil {
		return err
	}
	return nil
}

func GetRedis(key string) (string, bool){
	value, err := RDB.Get(key).Result()

	if err != nil {
		return "", true
	}
	
	return value, false
}