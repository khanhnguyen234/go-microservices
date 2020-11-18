package free_ship

import (
	"github.com/khanhnguyen234/go-microservices/_rabbitmq"
)

func PubFreeShip(msg string) {
	e := _rabbitmq.Exchange{
		Exchange: "free_ship",
		Type:     _rabbitmq.ExchangeFanout,
	}
	e.Pub(msg)
}
