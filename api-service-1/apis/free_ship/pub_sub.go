package free_ship

import (
	"khanhnguyen234/api-service-1/_rabbitmq"
)

func PubFreeShip(msg string) {
	e := _rabbitmq.Exchange{
		Exchange: "free_ship",
		Type:     _rabbitmq.ExchangeFanout,
	}
	e.Pub(msg)
}
