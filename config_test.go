package rabbit

import "testing"

func TestRabbitConfig(t *testing.T) {
	user := "rabbit"
	password := "123123"
	host := "example.com"
	vhost := "testing"
	port := "1711"

	config := NewRabbitConfig(user, password, host, vhost, port)

	if config.User != user {
		t.Errorf("User %s must be equal", user)
	}

	if config.Password != password {
		t.Errorf("Password %s must be equal", password)
	}

	if config.Host != host {
		t.Errorf("Host %s must be equal", host)
	}

	if config.VHost != vhost {
		t.Errorf("VHost %s must be equal", vhost)
	}

	if config.Port != port {
		t.Errorf("Port %s must be equal", port)
	}
}
