package main

import (
	"khanhnguyen234/api-service-2/rabbitmq"
)

type Queue struct {
	ExchangeName string
	ExchangeType string
	QueueName string
	RoutingKey string
}

func main() {
	rabbitmq.ConnectRabbitMQ()
	rabbitmq.LogsConsummer()

	forever := make(chan bool)
	<-forever
}