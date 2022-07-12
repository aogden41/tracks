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

// @title NAT Tracks API
// @version 2.0
// @description Parses the track message at https://www.notams.faa.gov/common/nat.html and stores a rolling 30-day track cache. Also handles events and concorde tracks.

// @contact.name Andrew Ogden
// @contact.url https://ganderoceanic.ca
// @contact.email a.ogden@vatcan.ca

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host tracks.ganderoceanic.ca
// @BasePath /docs
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
