package email

import (
	"encoding/json"
	"khanhnguyen234/job/_mongo"
	"khanhnguyen234/job/_rabbitmq"
	"khanhnguyen234/job/common"
	"log"
	"time"
)

func SubProductCreate() {
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
	common.LogError(err, "ExchangeDeclare")

	q, err := ch.QueueDeclare(
		queue.QueueName, // name
		false,           // durable
		false,           // delete when unused
		true,            // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	common.LogError(err, "QueueDeclare")

	err = ch.QueueBind(
		queue.QueueName,    // queue name
		queue.RoutingKey,   // routing key
		queue.ExchangeName, // exchange
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
			log.Println("[Received Sub]", queue)
			var data map[string]interface{}
			json.Unmarshal([]byte(d.Body), &data)

			time.Sleep(2 * time.Second)
			InsertToMongo(data)

			log.Println("[Sub]", queue, data)
		}
	}()

	log.Printf("[Sub Product Created] Waiting for messages.")
}

func InsertToMongo(data map[string]interface{}) {
	db, err := _mongo.ConnectMongo()
	c := db.Collection("product")

	ctx := common.GetContext()
	_, err = c.InsertOne(ctx, data)
	common.LogStatus(err, "Sub Product Created")
}
