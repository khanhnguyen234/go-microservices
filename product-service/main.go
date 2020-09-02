package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"khanhnguyen234/product-service/_elastic"
	"khanhnguyen234/product-service/_mongo"
	"khanhnguyen234/product-service/_rabbitmq"
	"khanhnguyen234/product-service/_redis"
	"khanhnguyen234/product-service/apis/product"
	"khanhnguyen234/product-service/common"
	"os"
)

func main() {
	err := godotenv.Load()
	common.LogStatus(err, "Load Env")

	_mongo.ConnectMongo()
	_redis.ConnectRedis()
	_elastic.ConnectElastic()
	_rabbitmq.ConnectRabbitMQ()

	//gin.SetMode(gin.ReleaseMode)
	route := gin.Default()

	product.ProductRouters(route.Group("/product"))
	product.CreateProductSub()

	route.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
