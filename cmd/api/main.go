package main

import (
	"log"

	"github.com/kyawzinkoko-dev/online-learning-platform/internal/bootstrap"
)

func main() {
	app, err := bootstrap.Build()
	if err != nil {
		log.Fatalf("Failed to build application: %v", err)
	}

	if err := app.Run(); err != nil {
		log.Fatalf("Failed to run application: %v", err)
	}
}
