package main

import (
	"github.com/aogden41/tracks/internal/api"
	"github.com/joho/godotenv"
	"log"
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
		s := api.Server{}
		err := s.Run()
		if err != nil {
			// Error
			log.Fatal(err)
		}
	}()
	select {} // Graceful shutdown
}
