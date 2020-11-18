package main

import (
	"fmt"
	"github.com/khanhnguyen234/go-microservices/_common"
	"github.com/khanhnguyen234/go-microservices/_elastic"
	"github.com/khanhnguyen234/go-microservices/_mongo"
	"github.com/khanhnguyen234/go-microservices/_rabbitmq"
	"github.com/khanhnguyen234/go-microservices/_redis"
	"khanhnguyen234/product-service/apis/product"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	err := godotenv.Load()
	_common.LogStatus(err, "Load Env")

	_mongo.ConnectMongo()
	_redis.ConnectRedis()
	_elastic.ConnectElastic()
	_rabbitmq.ConnectRabbitMQ()

	//gin.SetMode(gin.ReleaseMode)
	route := gin.Default()
	route.Use(CORS())

	product.ProductRouters(route.Group("/product"))
	product.CreateProductSub()

	route.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	route.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
