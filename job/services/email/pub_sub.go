package email

import (
	"encoding/json"
	"github.com/khanhnguyen234/go-microservices/_common"
	"github.com/khanhnguyen234/go-microservices/_mongo"
	"github.com/khanhnguyen234/go-microservices/_rabbitmq"
	"log"
	"time"
)

func SubProductCreate() {
	e := _rabbitmq.Exchange{
		Exchange: "product_headers",
		Type:     "headers",
		Headers: map[string]interface{}{
			"format": "pdf",
			"type":   "report",
		},
		Queue: "product",
	}

	msgs := e.Sub()

	go func() {
		for d := range msgs {
			log.Println("[Received Sub]", e)
			var data map[string]interface{}
			json.Unmarshal([]byte(d.Body), &data)

			time.Sleep(2 * time.Second)
			InsertToMongo(data)

			log.Println("[Sub]", e, data)
		}
	}()

	log.Printf("[Sub Product Created] Waiting for messages.")
}

func InsertToMongo(data map[string]interface{}) {
	db := _mongo.ConnectMongo()
	c := db.Collection("product")

	ctx := _common.GetContext()
	_, err := c.InsertOne(ctx, data)
	_common.LogStatus(err, "Sub Product Created")
}
