package main

import (
	"log"

	"github.com/millirud/go-service-boilerplate/config"
	"github.com/millirud/go-service-boilerplate/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
