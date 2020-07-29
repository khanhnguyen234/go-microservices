package redis

import (
	"github.com/gin-gonic/gin"
)

func RedisNoAuth(router *gin.RouterGroup) {
	router.GET("/count", RedisCount)
}
