package _rabbitmq

import (
	"github.com/streadway/amqp"
	"khanhnguyen234/api-service-2/common"
)

var rabbitmq *amqp.Connection

func ConnectRabbitMQ() (*amqp.Connection, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	rabbitmq = conn
	return conn, err
}

func GetRabbitMQ()  *amqp.Connection {
	return rabbitmq
}

func GetChannel()  (*amqp.Channel, error) {
	ch, err := rabbitmq.Channel()
	common.LogErrorService(err, "FAILED: GetChannel")
	return ch, err
}