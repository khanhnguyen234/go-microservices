package customer

import (
	"github.com/gin-gonic/gin"
)

func CustomerRouters(router *gin.RouterGroup) {
	router.GET("/order-management", GetOrderManagementRouter)
	router.GET("/order-management-v2", GetOrderManagementRouterV2)
}

func GetOrderManagementRouter(c *gin.Context) {
	result, err := GetOrderManagementController()
	c.JSON(200, gin.H{"result": result, "error": err})
	//GetOrderManagementController()
	//c.JSON(200, gin.H{"result": 1, "error": 2})
}

func GetOrderManagementRouterV2(c *gin.Context) {
	result, err := GetOrderManagementControllerV2()
	c.JSON(200, gin.H{"result": result, "error": err})
	//GetOrderManagementController()
	//c.JSON(200, gin.H{"result": 1, "error": 2})
}
