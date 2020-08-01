package main

import (
	"khanhnguyen234/api-service-2/_rabbitmq"
)

type Queue struct {
	ExchangeName string
	ExchangeType string
	QueueName string
	RoutingKey string
}

func main() {
	_rabbitmq.ConnectRabbitMQ()
	_rabbitmq.LogsConsummer()

	forever := make(chan bool)
	<-forever
}