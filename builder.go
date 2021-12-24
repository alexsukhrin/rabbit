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
