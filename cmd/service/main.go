package main

import "github.com/aogden41/tracks/internal/api"

func main() {
	go func() {
		// Run the server
		api.Run()
	}()
	select {} // Graceful shutdown
}
