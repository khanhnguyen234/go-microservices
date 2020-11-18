package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/khanhnguyen234/go-microservices/_common"
	"github.com/khanhnguyen234/go-microservices/_elastic"
	"github.com/khanhnguyen234/go-microservices/_mongo"
	"github.com/khanhnguyen234/go-microservices/_rabbitmq"
	"github.com/khanhnguyen234/go-microservices/_redis"
	"khanhnguyen234/seller-service/apis/shop"
	"os"
)

func main() {
	err := godotenv.Load()
	_common.LogStatus(err, "Load Env")

	_mongo.ConnectMongo()
	_redis.ConnectRedis()
	_elastic.ConnectElastic()
	_rabbitmq.ConnectRabbitMQ()

	//gin.SetMode(gin.ReleaseMode)
	route := gin.Default()

	shop.ShopRouters(route.Group("/shop"))

	route.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
