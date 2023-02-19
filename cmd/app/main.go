package main

import (
	"client-api/internal/app"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"os"
)

func main() {
	cfgFName := os.Getenv("CONFIG")
	if cfgFName == "" {
		cfgFName = "config.yml"
	}

	f, err := os.Open(cfgFName)
	if err != nil {
		log.Fatalf("open config file: %s", err.Error())
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			return
		}
	}(f)

	cfgContent, err := io.ReadAll(f)
	if err != nil {
		log.Printf("read config file: %s\n", err.Error())
	}

	var cfg app.Config

	err = yaml.Unmarshal(cfgContent, &cfg)
	if err != nil {
		log.Printf("unmarshal config file: %s", err.Error())
	}

	application := app.New(&cfg)

	err = application.Run()
	if err != nil {
		log.Printf("run app: %s", err.Error())
	}

}
