package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Routing
func RouteEndpoints() {
	// Routing functions
	// ...
}

// Run the server
func Run() {
	// Initialise router
	router := mux.NewRouter()

	// Begin routing endpoints
	RouteEndpoints()

	// Now serve
	log.Fatal(http.ListenAndServe(":5000", router))
}
