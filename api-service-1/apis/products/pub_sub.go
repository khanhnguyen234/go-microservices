package products

import (
	"github.com/khanhnguyen234/go-microservices/_rabbitmq"
)

func PubProductCreated(msg string) {
	e := _rabbitmq.Exchange{
		Exchange: "product_headers",
		Type:     _rabbitmq.ExchangeHeaders,
		Headers: map[string]interface{}{
			"format": "pdf",
			"type":   "report",
		},
	}
	e.Pub(msg)
}
