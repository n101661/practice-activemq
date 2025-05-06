package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	amqp "github.com/Azure/go-amqp"
	"github.com/n101661/practice-activemq/amqp/utils"
	"github.com/n101661/practice-activemq/internal/config"
)

var (
	path = flag.String("path", "./config.yaml", "The path of the config file.")
)

func main() {
	log.Println("start")
	defer log.Println("finish")

	flag.Parse()

	cfg, err := loadConfigAndValidate(*path)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := amqp.Dial(context.Background(), fmt.Sprintf("amqp://%s:%d", cfg.Host, cfg.Port), &amqp.ConnOptions{
		SASLType: amqp.SASLTypePlain(cfg.Consumer.Auth.Username, cfg.Consumer.Auth.Password),
	})
	if err != nil {
		log.Fatalf("failed to connect to MQ using AMQP: %v", err)
	}
	defer conn.Close()

	session, err := conn.NewSession(context.Background(), nil)
	if err != nil {
		log.Fatalf("failed to create a new session: %v", err)
	}
	defer session.Close(context.Background())

	receiver, err := session.NewReceiver(context.Background(), utils.FQQN(cfg.Consumer.AddressName, cfg.Consumer.QueueName), nil)
	if err != nil {
		log.Fatalf("failed to create a new receiver: %v", err)
	}
	defer receiver.Close(context.Background())

	ctx := context.Background()
	for {
		msg, err := receiver.Receive(ctx, nil)
		if err != nil {
			log.Fatalf("failed to receive message: %v", err)
		}
		for _, data := range msg.Data {
			log.Printf("got `%s`\n", string(data))
		}
		if err = receiver.AcceptMessage(ctx, msg); err != nil {
			log.Fatalf("failed to ack message: %v", err)
		}
	}
}

func loadConfigAndValidate(path string) (*config.AMQP, error) {
	cfg, err := config.Load(path)
	if err != nil {
		return nil, err
	}

	if cfg == nil || cfg.AMQP == nil {
		return nil, fmt.Errorf("missing amqp config")
	}
	if cfg.AMQP.Consumer == nil {
		return nil, fmt.Errorf("missing amqp.consumer config")
	}
	if cfg.AMQP.Consumer.Auth == nil {
		return nil, fmt.Errorf("missing amqp.consumer.auth config")
	}

	return cfg.AMQP, nil
}
