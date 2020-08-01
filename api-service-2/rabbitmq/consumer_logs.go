package rabbitmq

import (
	"log"
	"khanhnguyen234/api-service-2/common"
)

func LogsConsummer() {
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
			log.Printf("[LogsConsumer] %s", d.Body)
		}
	}()

	log.Printf("[LogsConsumer] Waiting for messages.")
}