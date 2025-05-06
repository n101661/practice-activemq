package main

import (
	"flag"
	"log"
	"os"
	"reflect"

	yaml "github.com/goccy/go-yaml"
	"github.com/n101661/practice-activemq/internal/config"
)

var path = flag.String("path", "./config.yaml", "The path of the config file.")

func main() {
	flag.Parse()

	cfg, err := config.Load(*path)
	if err != nil {
		log.Fatal(err)
	}

	setDefault(cfg)

	file, err := os.Create(*path)
	if err != nil {
		log.Fatalf("failed to create config file: %v", err)
	}
	defer file.Close()
	defer file.Sync()

	data, err := yaml.Marshal(cfg)
	if err != nil {
		log.Fatalf("failed to marshal config: %v", err)
	}
	if n, err := file.Write(data); err != nil || n != len(data) {
		log.Fatalf("failed to write config file or the content is truncated: %v", err)
	}
}

func setDefault(cfg *config.Config) {
	rv := reflect.ValueOf(cfg)
	reflectDefault(rv)
}

func reflectDefault(rv reflect.Value) {
	switch kind := rv.Kind(); kind {
	case reflect.Ptr:
		if rv.IsNil() {
			v := reflect.New(rv.Type().Elem())
			reflectDefault(v)
			rv.Set(v)
		} else {
			reflectDefault(rv.Elem())
		}
	case reflect.Struct:
		fieldNum := rv.NumField()
		for i := range fieldNum {
			field := rv.Field(i)
			reflectDefault(field)
		}
	}
}
