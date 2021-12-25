package rabbit

import "github.com/streadway/amqp"

type PublishParams struct {
	Exchange, Routing    string
	Mandatory, Immediate bool
	Params               *amqp.Publishing
}

type Publisher interface {
	Publish(*PublishParams) error
}

func (broker MessageBroker) Publish(p *PublishParams) error {
	return broker.Channel.Publish(
		p.Exchange,
		p.Routing,
		p.Mandatory,
		p.Immediate,
		*p.Params,
	)
}
