package _rabbitmq

type Exchange struct {
	Exchange   string
	Type       string
	RoutingKey string
	Queue      string
	Headers    map[string]interface{}
}

const ExchangeDirect = "direct"
const ExchangeFanout = "fanout"
const ExchangeTopic = "topic"
const ExchangeHeaders = "headers"

type PubDirectExchange struct {
	ExchangeName string
	RoutingKey   string
}

type SubDirectExchange struct {
	ExchangeName string
	RoutingKey   string
	QueueName    string
}

type PubFanoutExchange struct {
	ExchangeName string
}

type SubFanoutExchange struct {
	ExchangeName string
	QueueName    string
}

type PubTopicExchange struct {
	ExchangeName string
	RoutingKey   string `string.matchAll`
}

type SubTopicExchange struct {
	ExchangeName string
	RoutingKey   string `string.#`
	QueueName    string
}

type PubHeadersExchange struct {
	ExchangeName string
	Headers      map[string]interface{}
}

type SubHeadersExchange struct {
	ExchangeName string
	Headers      map[string]interface{}
	QueueName    string
}
