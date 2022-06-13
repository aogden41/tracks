package main

import (
	"log"

	"github.com/aogden41/tracks/internal/api"
	"github.com/joho/godotenv"
)

func init() {
	// Load ENV
	if err := godotenv.Load("../../.env"); err != nil {
		log.Print(".env file not found.")
	}
}

func main() {
	go func() {
		// Run the server
		api.Run()
	}()
	select {} // Graceful shutdown
}
