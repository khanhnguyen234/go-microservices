package main

import (
	"github.com/joho/godotenv"
	"khanhnguyen234/api-service-2/_mongo"
	"khanhnguyen234/api-service-2/_rabbitmq"
	"khanhnguyen234/api-service-2/common"
	"khanhnguyen234/api-service-2/services/free_ship"
)

func main() {
	err := godotenv.Load()
	common.LogStatus(err, "Load Env")

	_mongo.ConnectMongo()
	_, err = _rabbitmq.ConnectRabbitMQ()

	if err != nil {
		return
	}

	free_ship.FreeShipConsumer()

	forever := make(chan bool)
	<-forever
}
