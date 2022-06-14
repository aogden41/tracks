package api

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Routing
func RouteEndpoints(router *mux.Router) {
	// Index
	/////////
	router.HandleFunc("/", Index).Methods("GET")

	// Current tracks
	/////////////////
	// GET
	router.HandleFunc("/current", GetAllCurrentTracks).Methods("GET")
	router.HandleFunc("/current/{track_id}", GetCurrentTrack).Methods("GET")
	router.HandleFunc("/current/eastbound", GetCurrentEastboundTracks).Methods("GET")
	router.HandleFunc("/current/westbound", GetCurrentWestboundTracks).Methods("GET")
	router.HandleFunc("/current/now", GetCurrentTracksValidNow).Methods("GET")
	router.HandleFunc("/current/later", GetCurrentTracksValidLater).Methods("GET")
	router.HandleFunc("/current/earlier", GetCurrentTracksValidEarlier).Methods("GET")

	// Cached tracks
	////////////////
	// GET
	router.HandleFunc("/cached", GetAllCachedTracks).Methods("GET")
	router.HandleFunc("/cached/{track_id}", GetCachedTrack).Methods("GET")
	router.HandleFunc("/cached/days/{days_old}", GetCachedTracksByDaysOld).Methods("GET")
	router.HandleFunc("/cached/eastbound", GetCachedEastboundTracks).Methods("GET")
	router.HandleFunc("/cached/westbound", GetCachedWestboundTracks).Methods("GET")
	router.HandleFunc("/cached/check/{track_id}", CheckIsTrackCached).Methods("GET")

	// Event tracks
	///////////////
	// GET
	router.HandleFunc("/event", GetAllEventTracks).Methods("GET")
	router.HandleFunc("/event/{track_id}", GetEventTrack).Methods("GET")
	// POST
	router.HandleFunc("/event/{track_obj}", PostEventTrack).Methods("POST")
	// DELETE
	router.HandleFunc("/event/{track_id}", DeleteEventTrack).Methods("DELETE")

	// Concorde tracks
	//////////////////
	// GET
	router.HandleFunc("/concorde", GetAllConcordeTracks).Methods("GET")
	router.HandleFunc("/concorde/{track_id}", GetConcordeTrack).Methods("GET")

	// Fixes
	////////
	// GET
	router.HandleFunc("/fixes", GetAllFixes).Methods("GET")
	router.HandleFunc("/fixes/{fix_name}", GetFix).Methods("GET")
	// POST
	router.HandleFunc("/fixes/{fix_obj}", PostFix).Methods("POST")
	// UPDATE
	router.HandleFunc("/fixes/{fix_obj}", UpdateFix).Methods("UPDATE")
	// DELETE
	router.HandleFunc("/fixes/{fix_name}", DeleteFix).Methods("DELETE")
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
