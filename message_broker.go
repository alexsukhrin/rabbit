package rabbit

import (
	"github.com/streadway/amqp"
	"log"
)

type ConsumeParams struct {
	Queue, ConsumerName                 string
	AutoAck, Exclusive, NoLocal, NoWait bool
	Args                                map[string]interface{}
}

type PublishParams struct {
	Exchange, Routing    string
	Mandatory, Immediate bool
	Body                 []byte
}

type MessageBroker interface {
	Consume(*ConsumeParams) <-chan amqp.Delivery
	Publish(*PublishParams) error
}

type Rabbit struct {
	Channel *amqp.Channel
}

func GetConnect(connUrl string) *amqp.Connection {
	conn, err := amqp.Dial(connUrl)

	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ")
	}

	return conn
}

func GetChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()

	if err != nil {
		log.Fatal("Failed to open a channel")
	}

	return ch
}

func (r *Rabbit) Consume(p *ConsumeParams) <-chan amqp.Delivery {
	events, err := r.Channel.Consume(
		p.Queue,
		p.ConsumerName,
		p.AutoAck,
		p.Exclusive,
		p.NoLocal,
		p.NoWait,
		p.Args,
	)

	if err != nil {
		log.Fatal("Failed to register a consumer")
	}

	return events
}

func (r *Rabbit) Publish(p *PublishParams) error {
	params := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(p.Body),
	}
	return r.Channel.Publish(
		p.Exchange,
		p.Routing,
		p.Mandatory,
		p.Immediate,
		params,
	)
}
