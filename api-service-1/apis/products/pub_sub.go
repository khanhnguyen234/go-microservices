package products

import (
	"khanhnguyen234/api-service-1/_rabbitmq"
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
