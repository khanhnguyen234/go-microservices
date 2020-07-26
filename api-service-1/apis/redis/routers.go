package redis

import (
	"github.com/gin-gonic/gin"
)

func RedisNoAuthCount(router *gin.RouterGroup) {
	router.GET("/count", RedisCount)
}
