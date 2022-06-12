package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Routing
func RouteEndpoints(router *mux.Router) {
	// Routes
	router.HandleFunc("/", Index).Methods("GET")
	router.HandleFunc("/data", Get).Methods("GET")
	router.HandleFunc("/concorde", GetConcorde).Methods("GET")
	router.HandleFunc("/event", GetEvent).Methods("GET")

}

// Run the server
func Run() {
	// Initialise router
	router := mux.NewRouter()

	// Begin routing endpoints
	RouteEndpoints(router)

	// Now serve
	log.Fatal(http.ListenAndServe(":5000", router))
}
