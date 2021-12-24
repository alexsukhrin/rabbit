package rabbit

import (
	"github.com/streadway/amqp"
	"log"
)

func Connect(rabbitUrl string) *amqp.Connection {
	conn, err := amqp.Dial(rabbitUrl)

	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ")
	}

	return conn
}
