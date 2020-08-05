package auth

import (
	"github.com/gin-gonic/gin"
	"khanhnguyen234/api-gateway/common"
	"net/http"
)

func AuthRouters(router *gin.RouterGroup) {
	router.POST("/sign-up", SignUpRouter)
	router.POST("/sign-in", SignInRouter)
	router.GET("/context", AuthContextRouter)
}

func SignUpRouter(c *gin.Context) {
	var validator SignUpValidator

	if err := c.ShouldBindJSON(&validator.Request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	auth := SignUpController(validator)
	c.JSON(http.StatusCreated, gin.H{"auth": auth})
}

func SignInRouter(c *gin.Context) {
	var validator SignUpValidator

	if err := c.ShouldBindJSON(&validator.Request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	auth, err := SignInController(validator)

	c.JSON(http.StatusOK, gin.H{"auth": auth, "err": err})
}

func AuthContextRouter(c *gin.Context) {
	auth, err := AuthContextController(c.Request)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"auth": nil, "err": common.ErrorToString(err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"auth": auth, "err": nil})
}
