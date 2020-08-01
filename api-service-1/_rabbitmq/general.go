package _rabbitmq

import (
	"github.com/streadway/amqp"
	"khanhnguyen234/api-service-1/common"
	"os"
)

var rabbitmq *amqp.Connection

func ConnectRabbitMQ() (*amqp.Connection, error) {
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_DIAL"))
	common.LogErrorService(err, "Connect RabbitMQ")
	rabbitmq = conn
	common.LogSuccess("Connect RabbitMQ")
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