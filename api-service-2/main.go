package main

import (
	"github.com/joho/godotenv"
	"github.com/khanhnguyen234/go-microservices/_common"
	"github.com/khanhnguyen234/go-microservices/_mongo"
	"github.com/khanhnguyen234/go-microservices/_rabbitmq"
	"khanhnguyen234/api-service-2/services/free_ship"
)

func main() {
	err := godotenv.Load()
	_common.LogStatus(err, "Load Env")

	_mongo.ConnectMongo()
	_, err = _rabbitmq.ConnectRabbitMQ()

	if err != nil {
		return
	}

	free_ship.FreeShipConsumer()

	forever := make(chan bool)
	<-forever
}
