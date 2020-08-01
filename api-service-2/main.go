package main

import (
	"github.com/joho/godotenv"
	"khanhnguyen234/api-service-2/_rabbitmq"
	"khanhnguyen234/api-service-2/_mongo"
	"khanhnguyen234/api-service-2/common"
	"khanhnguyen234/api-service-2/services/free_ship"
)

type Queue struct {
	ExchangeName string
	ExchangeType string
	QueueName string
	RoutingKey string
}

func main() {
	err := godotenv.Load()
	common.LogErrorService(err, "Load Env")

	_mongo.ConnectMongo()
	_rabbitmq.ConnectRabbitMQ()
	free_ship.FreeShipConsummer()

	forever := make(chan bool)
	<-forever
}