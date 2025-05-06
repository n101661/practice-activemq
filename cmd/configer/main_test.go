package main

import (
	"testing"

	"github.com/n101661/practice-activemq/internal/config"
	"github.com/stretchr/testify/assert"
)

func Test_setDefault(t *testing.T) {
	var cfg config.Config
	setDefault(&cfg)
	assert.Equal(t, &config.Config{
		AMQP: &config.AMQP{
			Host: "",
			Port: 0,
			Provider: &config.AMQPProvider{
				Auth: &config.AMQPAuth{
					Username: "",
					Password: "",
				},
				AddressName: "",
				QueueName:   "",
			},
			Consumer: &config.AMQPConsumer{
				Auth: &config.AMQPAuth{
					Username: "",
					Password: "",
				},
				AddressName: "",
				QueueName:   "",
			},
		},
	}, &cfg)
}
