package rabbit

import "github.com/streadway/amqp"

type MessageBroker struct {
	Channel *amqp.Channel
}
