package products

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"khanhnguyen234/api-service-1/_rabbitmq"
	"khanhnguyen234/api-service-1/common"
	"log"
)

func PubProductCreated(product ProductModel) {
	ch, err := _rabbitmq.GetChannel()

	queue := _rabbitmq.Queue{
		ExchangeName: "product",
		ExchangeType: "fanout",
		QueueName:    "created",
		RoutingKey:   "RoutingKey",
	}

	err = ch.ExchangeDeclare(
		queue.ExchangeName, // name
		queue.ExchangeType, // type
		true,               // durable
		false,              // auto-deleted
		false,              // internal
		false,              // no-wait
		nil,                // arguments
	)
	common.LogStatus(err, "ExchangeDeclare product_created")

	jsonBody, _ := json.Marshal(product)
	body := string(jsonBody)

	err = ch.Publish(
		queue.ExchangeName, // exchange
		queue.RoutingKey,   // routing key
		false,              // mandatory
		false,              // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	log.Println("[Pub]", queue, body)
}
