package rabbit

import (
	"github.com/streadway/amqp"
	"log"
)

type ConsumeParams struct {
	Queue, ConsumerName                 string
	AutoAck, Exclusive, NoLocal, NoWait bool
	Args                                *map[string]interface{}
}

type PublishParams struct {
	Exchange, Routing    string
	Mandatory, Immediate bool
	Params               *amqp.Publishing
}

type MessageBroker interface {
	Consume(*ConsumeParams)
	Publish(*PublishParams) error
	Connect(rabbitUrl string) *amqp.Connection
	Channel(conn *amqp.Connection) *amqp.Channel
}

type Rabbit struct {
	getConnection *amqp.Connection
	getChannel *amqp.Channel
}

func (r Rabbit) Channel(conn *amqp.Connection) *amqp.Channel {
	ch, err := r.getConnection.Channel()

	if err != nil {
		log.Fatal("Failed to open a channel")
	}

	return ch
}

func (r Rabbit) Connect(rabbitUrl string) *amqp.Connection {
	conn, err := amqp.Dial(rabbitUrl)

	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ")
	}

	return conn
}

func (r Rabbit) Consume(p *ConsumeParams) (<-chan amqp.Delivery, error) {
	return r.getChannel.Consume(
		p.Queue,
		p.ConsumerName,
		p.AutoAck,
		p.Exclusive,
		p.NoLocal,
		p.NoWait,
		*p.Args,
	)
}

func (r Rabbit) Publish(p *PublishParams) error {
	return r.getChannel.Publish(
		p.Exchange,
		p.Routing,
		p.Mandatory,
		p.Immediate,
		*p.Params,
	)
}