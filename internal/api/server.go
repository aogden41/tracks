package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Routing
func RouteEndpoints(router *mux.Router) {
	// Index/others
	router.HandleFunc("/", Index).Methods("GET")

	// Normal tracks
	router.HandleFunc("/data", GetAllTracks).Methods("GET")
	router.HandleFunc("/data/{track_id}", GetOneTrack).Methods("GET")

	// Concorde tracks
	router.HandleFunc("/concorde", GetAllConcordeTracks).Methods("GET")
	router.HandleFunc("/concorde/{track_id}", GetOneTrack).Methods("GET")

	// Event tracks
	router.HandleFunc("/event", GetAllEventTracks).Methods("GET")
	router.HandleFunc("/concorde/{track_id}", GetOneTrack).Methods("GET")

	// Fixes
	router.HandleFunc("/fixes", GetAllFixes).Methods("GET")
	router.HandleFunc("/fixes/{fix_name}", GetOneFix).Methods("GET")
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
