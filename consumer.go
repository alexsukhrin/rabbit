package rabbit

import "github.com/streadway/amqp"

type ConsumeParams struct {
	Queue, ConsumerName                 string
	AutoAck, Exclusive, NoLocal, NoWait bool
	Args                                *map[string]interface{}
}

type Consumer interface {
	Consume(*ConsumeParams)
}

func (broker MessageBroker) Consume(p *ConsumeParams) (<-chan amqp.Delivery, error) {
	return broker.Channel.Consume(
		p.Queue,
		p.ConsumerName,
		p.AutoAck,
		p.Exclusive,
		p.NoLocal,
		p.NoWait,
		*p.Args,
	)
}
