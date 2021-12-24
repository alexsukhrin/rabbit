package rabbit

import (
	"github.com/streadway/amqp"
	"log"
)

func Channel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()

	if err != nil {
		log.Fatal("Failed to open a channel")
	}

	return ch
}
