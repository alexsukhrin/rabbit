package rabbit

import "fmt"

func BuilderRabbitUrl(config RabbitConfig) string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.VHost,
	)
}

func BuildRabbitUrl(host, port, vhost, user, password string) string {
	rabbitConfig := NewRabbitConfig(user, password, host, vhost, port)
	return BuilderRabbitUrl(*rabbitConfig)
}
