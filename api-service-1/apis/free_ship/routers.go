package free_ship

import (
	"github.com/gin-gonic/gin"
)

func FreeshipNoAuth(router *gin.RouterGroup) {
	router.GET("/publish", Publish)
}
