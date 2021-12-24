package rabbit

import (
	"testing"
)

func TestBu(t *testing.T) {
	user := "rabbit"
	password := "123123"
	host := "example.com"
	vhost := "testing"
	port := "1711"

	config := NewRabbitConfig(user, password, host, vhost, port)
	rabbitUrl := BuilderRabbitUrl(*config)
	url := "amqp://rabbit:123123@example.com:1711/testing"

	if rabbitUrl != url {
		t.Errorf("Builder %s must be %s", rabbitUrl, url)
	}
}
