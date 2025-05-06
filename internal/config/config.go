package config

import (
	"fmt"
	"os"

	yaml "github.com/goccy/go-yaml"
)

type Config struct {
	AMQP *AMQP `yaml:"amqp"`
}

type AMQP struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`

	Provider *AMQPProvider `yaml:"provider"`
	Consumer *AMQPConsumer `yaml:"consumer"`
}

type AMQPProvider = amqpIO

type AMQPConsumer = amqpIO

type amqpIO struct {
	Auth *AMQPAuth `yaml:"auth"`
	// AddressName is the name of the address where messages are sent.
	// If the value is empty, it will use a random name as address name.
	AddressName string `yaml:"addressName"`
	// QueueName is a specified name of the queue of the address.
	// The empty value means unspecified name.
	QueueName string `yaml:"queueName"`
}

type AMQPAuth struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("failed to read %s file: %v", path, err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %v", err)
	}

	return &cfg, nil
}
