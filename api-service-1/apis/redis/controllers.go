package redis

import (
	"github.com/gin-gonic/gin"
	"khanhnguyen234/api-service-1/common"
	"strconv"
	"time"
)


func RedisCount (c *gin.Context) {
	value := Increase()
	c.JSON(200, gin.H{"result": value})
}

func Increase() int {
	key := "COUNT_REQUEST"
	var value int 

	redis := common.Redis()
	stringPrev, err := redis.Get(key).Result()
	intPrev, err := strconv.Atoi(stringPrev)

	if err != nil {
		intPrev = 0
	}

	value = intPrev + 1
	// redis.Set(key, value, 0)
	redis.Set(key, value, 5 * time.Minute)

	return value
}
