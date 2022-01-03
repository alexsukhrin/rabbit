package rabbit

import (
	"fmt"
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
	User, Password, Host, VHost, Port, ConnectionUrl string
	Connection                                       *amqp.Connection
	Chan                                             *amqp.Channel
}

func (rabbit *Rabbit) BuilderConnectionUrl() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/%s",
		rabbit.User,
		rabbit.Password,
		rabbit.Host,
		rabbit.Port,
		rabbit.VHost,
	)
}

func (rabbit *Rabbit) Connect() *amqp.Connection {
	conn, err := amqp.Dial(rabbit.ConnectionUrl)

	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ")
	}

	return conn
}

func (rabbit *Rabbit) Channel() *amqp.Channel {
	ch, err := rabbit.Connection.Channel()

	if err != nil {
		log.Fatal("Failed to open a channel")
	}

	return ch
}

func (rabbit *Rabbit) Consume(p *ConsumeParams) <-chan amqp.Delivery {
	events, err := rabbit.Chan.Consume(
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

func (rabbit *Rabbit) Publish(p *PublishParams) error {
	params := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(p.Body),
	}
	return rabbit.Chan.Publish(
		p.Exchange,
		p.Routing,
		p.Mandatory,
		p.Immediate,
		params,
	)
}
