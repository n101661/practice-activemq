package utils

import (
	"crypto/tls"
	"log"

	"github.com/n101661/practice-activemq/internal/config"
	"github.com/n101661/practice-activemq/internal/ssl"
)

func ParseSSLConfig(cfg *config.SSLConfig) (scheme string, config *tls.Config, err error) {
	if cfg.Enable {
		ca, err := ssl.NewCertPool(cfg.AppendedCert)
		if err != nil {
			return "", nil, err
		}

		return "amqps", &tls.Config{
			RootCAs: ca,
		}, nil
	}
	log.Println("[warn] use insecure protocol")
	return "amqp", nil, nil
}

func FQQN(address, queue string) string {
	if queue == "" {
		return address
	}
	return address + "::" + queue
}
