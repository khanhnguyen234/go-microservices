package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"khanhnguyen234/api-service-1/common"
	"khanhnguyen234/api-service-1/rabbitmq"
	"khanhnguyen234/api-service-1/apis/products"
	"khanhnguyen234/api-service-1/apis/redis"
	"khanhnguyen234/api-service-1/apis/freeship"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&products.ProductModel{})
}

func main() {
	db := common.InitPostgreSQL()
	Migrate(db)

	common.InitRedis()
	common.InitElasticsearch()
	rabbitmq.ConnectRabbitMQ()

	route := gin.Default()
	route.GET("/query", query)
	route.GET("/param/:name/:id", param)

	noAuth := route.Group("/no-auth")
	products.ProductNoAuthRegister(noAuth.Group("/products"))
	redis.RedisNoAuth(noAuth.Group("/redis"))
	freeship.FreeshipNoAuth(noAuth.Group("/freeship"))
	
	route.Run(":7001")
}

type PersonQuery struct {
	Name string `form:"name"`
	ID string `form:"id"`
}

func query(c *gin.Context) {
	var person PersonQuery
	if c.ShouldBindQuery(&person) == nil {
		fmt.Println("====== Only Bind By Query String ======")
		fmt.Println(person.Name)
		fmt.Println(person.ID)
	}
	c.JSON(200, gin.H{"name": person.Name, "id": person.ID})
}


type PersonParam struct {
	Name string `uri:"name"`
	ID string `uri:"id"`
}

func param(c *gin.Context)  {
	var person PersonParam
	if err := c.ShouldBindUri(&person); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	c.JSON(200, gin.H{"name": person.Name, "id": person.ID})
}