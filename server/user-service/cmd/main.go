package main

import (
	"log"
	"user-service/internal"
)

func main() {
	application := internal.NewApp()

	if err := application.Run(); err != nil {
		log.Fatalf("Failed to start application: %v", err)
	}
}
