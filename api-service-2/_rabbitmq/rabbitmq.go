package _rabbitmq

import (
	"github.com/streadway/amqp"
	"khanhnguyen234/api-service-2/common"
	"log"
	"os"
)

var rabbitmq *amqp.Connection

func ConnectRabbitMQ() (*amqp.Connection, error) {
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_DIAL"))
	common.LogStatus(err, "Connect RabbitMQ")
	rabbitmq = conn
	return conn, err
}

func GetRabbitMQ() *amqp.Connection {
	return rabbitmq
}

func GetChannel() (*amqp.Channel, error) {
	ch, err := rabbitmq.Channel()
	common.LogStatus(err, "Get Channel")
	return ch, err
}

func (s Exchange) Pub(body string) {
	ch, err := GetChannel()

	err = ch.ExchangeDeclare(
		s.Exchange, // name
		s.Type,     // type
		true,       // durable
		false,      // auto-deleted
		false,      // internal
		false,      // no-wait
		s.Headers,  // arguments
	)
	if err != nil {
		log.Println("[ExchangeDeclare]", s, err)
	}

	err = ch.Publish(
		s.Exchange,   // exchange
		s.RoutingKey, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	if err != nil {
		log.Println("[Publish]", s, body)
	}
}

func (s Exchange) Sub() <-chan amqp.Delivery {
	ch, err := GetChannel()

	err = ch.ExchangeDeclare(
		s.Exchange, // name
		s.Type,     // type
		true,       // durable
		false,      // auto-deleted
		false,      // internal
		false,      // no-wait
		s.Headers,  // arguments
	)
	if err != nil {
		log.Println("[ExchangeDeclare]", s, err)
	}

	q, err := ch.QueueDeclare(
		s.Queue, // queue
		false,   // durable
		false,   // delete when unused
		true,    // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Println("[QueueDeclare]", s, err)
	}

	err = ch.QueueBind(
		s.Queue,      // queue name
		s.RoutingKey, // routing key
		s.Exchange,   // exchange
		false,
		nil,
	)
	if err != nil {
		log.Println("[QueueBind]", s, err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Println("[Consume]", s, err)
	}

	return msgs
}
