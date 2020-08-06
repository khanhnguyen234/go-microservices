package free_ship

import (
	"encoding/json"
	"khanhnguyen234/api-service-2/_rabbitmq"
	"khanhnguyen234/api-service-2/common"
	"log"
)

func FreeShipConsumer() {
	ch, err := _rabbitmq.GetChannel()

	queueLogs := _rabbitmq.Queue{
		ExchangeName: "logs",
		ExchangeType: "fanout",
		QueueName:    "",
		RoutingKey:   "",
	}

	err = ch.ExchangeDeclare(
		queueLogs.ExchangeName, // name
		queueLogs.ExchangeType, // type
		true,                   // durable
		false,                  // auto-deleted
		false,                  // internal
		false,                  // no-wait
		nil,                    // arguments
	)
	common.LogError(err, "ExchangeDeclare")

	q, err := ch.QueueDeclare(
		queueLogs.QueueName, // name
		false,               // durable
		false,               // delete when unused
		true,                // exclusive
		false,               // no-wait
		nil,                 // arguments
	)
	common.LogError(err, "QueueDeclare")

	err = ch.QueueBind(
		queueLogs.QueueName,    // queue name
		queueLogs.RoutingKey,   // routing key
		queueLogs.ExchangeName, // exchange
		false,
		nil,
	)
	common.LogError(err, "QueueBind")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	common.LogError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			var data FreeshipCreateConsume
			json.Unmarshal([]byte(d.Body), &data)
			InsertToMongo(data)

			log.Printf("[Free Ship Consumer] %s", d.Body)
		}
	}()

	log.Printf("[Free Ship Consumer] Waiting for messages.")
}
