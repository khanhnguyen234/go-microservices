package free_ship

import (
	"log"
	"encoding/json"
	"khanhnguyen234/api-service-2/common"
	"khanhnguyen234/api-service-2/_rabbitmq"
)

func FreeShipConsummer() {
	ch, err := _rabbitmq.GetChannel()

	queueLogs := _rabbitmq.Queue{
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
	common.LogErrorService(err, "Failed to declare a queue")

	q, err := ch.QueueDeclare(
		queueLogs.QueueName,    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	common.LogErrorService(err, "Failed to register a consumer")

	err = ch.QueueBind(
		queueLogs.QueueName, // queue name
		queueLogs.RoutingKey,     // routing key
		queueLogs.ExchangeName, // exchange
		false,
		nil,
	)
	common.LogErrorService(err, "Failed to register a consumer")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	common.LogErrorService(err, "Failed to register a consumer")

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