package main

import (
	"khanhnguyen234/api-service-2/_rabbitmq"
	"khanhnguyen234/api-service-2/_mongo"
	"khanhnguyen234/api-service-2/services/free_ship"
)

type Queue struct {
	ExchangeName string
	ExchangeType string
	QueueName string
	RoutingKey string
}

func main() {
	_mongo.ConnectMongo()
	_rabbitmq.ConnectRabbitMQ()
	free_ship.FreeShipConsummer()

	forever := make(chan bool)
	<-forever
}