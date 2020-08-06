package _rabbitmq

type Queue struct {
	ExchangeName string
	ExchangeType string
	QueueName    string
	RoutingKey   string
}
