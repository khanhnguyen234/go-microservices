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
	db, err := _mongo.ConnectMongo()
	c := db.Collection("product")

	ctx := common.GetContext()
	_, err = c.InsertOne(ctx, data)
	common.LogStatus(err, "Sub Product Created")
}
