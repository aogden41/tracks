package server

import (
	"log"
	"net/http"

	"github.com/aogden41/tracks/internal/api"
	"github.com/gorilla/mux"
)

// Routing
func RouteEndpoints(router *mux.Router) {
	// Routes
	router.HandleFunc("/", api.Index).Methods("GET")
	router.HandleFunc("/data", api.Get).Methods("GET")
	router.HandleFunc("/concorde", api.GetConcorde).Methods("GET")
	router.HandleFunc("/event", api.GetEvent).Methods("GET")

}

// Run the server
func Run() {
	go func() {
		// Initialise router
		router := mux.NewRouter()

		// Begin routing endpoints
		RouteEndpoints(router)

		// Now serve
		log.Fatal(http.ListenAndServe(":5000", router))
	}()
	select {} // Prevent immediate exit
}
