package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"khanhnguyen234/seller-service/_elastic"
	"khanhnguyen234/seller-service/_mongo"
	"khanhnguyen234/seller-service/_rabbitmq"
	"khanhnguyen234/seller-service/_redis"
	"khanhnguyen234/seller-service/apis/shop"
	"khanhnguyen234/seller-service/common"
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

	shop.ShopRouters(route.Group("/shop"))

	route.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
