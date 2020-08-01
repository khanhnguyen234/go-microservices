package _rabbitmq

import (
	"github.com/streadway/amqp"
	"khanhnguyen234/api-service-2/common"
	"os"
)

var rabbitmq *amqp.Connection

func ConnectRabbitMQ() (*amqp.Connection, error) {
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_DIAL"))
	common.LogStatus(err, "Connect RabbitMQ")
	rabbitmq = conn
	return conn, err
}

func GetRabbitMQ()  *amqp.Connection {
	return rabbitmq
}

func GetChannel()  (*amqp.Channel, error) {
	ch, err := rabbitmq.Channel()
	common.LogError(err, "FAILED: GetChannel")
	return ch, err
}