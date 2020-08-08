package free_ship

import (
	"encoding/json"
	"khanhnguyen234/api-service-2/_rabbitmq"
	"log"
)

func FreeShipConsumer() {
	e := _rabbitmq.Exchange{
		Exchange: "free_ship",
		Type:     _rabbitmq.ExchangeFanout,
		Queue:    "insert_to_mongo",
	}

	msgs := e.Sub()

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
