package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

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
		SASLType: amqp.SASLTypePlain(cfg.Provider.Auth.Username, cfg.Provider.Auth.Password),
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

	sender, err := session.NewSender(context.Background(), utils.FQQN(cfg.Provider.AddressName, cfg.Provider.QueueName), nil)
	if err != nil {
		log.Fatalf("failed to create a new sender: %v", err)
	}
	defer sender.Close(context.Background())

	ticker := time.NewTicker(time.Second)
	ctx := context.Background()
	counter := 0
	for {
		<-ticker.C

		text := "ping" + strconv.Itoa(counter)
		msg := amqp.Message{
			Value: text,
		}
		if err := sender.Send(ctx, &msg, nil); err != nil {
			log.Fatalf("failed to send message: %v", err)
		}
		log.Printf("sent a message [%s]\n", text)
		counter++
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
	if cfg.AMQP.Provider == nil {
		return nil, fmt.Errorf("missing amqp.provider config")
	}
	if cfg.AMQP.Provider.Auth == nil {
		return nil, fmt.Errorf("missing amqp.provider.auth config")
	}

	return cfg.AMQP, nil
}
