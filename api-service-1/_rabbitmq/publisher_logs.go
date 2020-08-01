package _rabbitmq

import (
	"log"
	"khanhnguyen234/api-service-1/common"
	"github.com/streadway/amqp"
)

func LogsPublisher(body string) {
	ch, err := GetChannel()

	queueLogs := Queue{
		ExchangeName: "logs",
		ExchangeType: "fanout",
		QueueName: "",
		RoutingKey: "",
	}

	err = ch.ExchangeDeclare(
		queueLogs.ExchangeName,   // name
		queueLogs.ExchangeType, // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	common.LogStatus(err, "Failed to declare a queue")
	
	err = ch.Publish(
		"logs", // exchange
		"",     // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	log.Printf("[LogsPublisher] %s", body)
}