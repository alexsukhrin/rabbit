package rabbit

type RabbitConfig struct {
	User, Password, Host, VHost, Port string
}

func NewRabbitConfig(user, password, host, vhost, port string) *RabbitConfig {
	return &RabbitConfig{
		User:     user,
		Password: password,
		Host:     host,
		VHost:    vhost,
		Port:     port,
	}
}
