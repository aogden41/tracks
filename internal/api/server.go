package api

import (
	"github.com/aogden41/tracks/internal/api/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	Router *mux.Router
}

// Routing
func (s Server) Initialise() (router *mux.Router) {
	// Initialise router
	r := mux.NewRouter()

	// Index
	/////////
	r.HandleFunc("/", handlers.Index).Methods("GET")

	// Current tracks
	/////////////////
	// Router
	rCurrentTracks := r.PathPrefix("/current").Subrouter()
	// GET
	rCurrentTracks.HandleFunc("", handlers.GetAllCurrentTracks).Methods("GET")
	rCurrentTracks.HandleFunc("/", handlers.GetAllCurrentTracks).Methods("GET")
	rCurrentTracks.HandleFunc("/{track_id}", handlers.GetCurrentTrack).Methods("GET")
	rCurrentTracks.HandleFunc("/eastbound", handlers.GetCurrentEastboundTracks).Methods("GET")
	rCurrentTracks.HandleFunc("/westbound", handlers.GetCurrentWestboundTracks).Methods("GET")
	rCurrentTracks.HandleFunc("/now", handlers.GetCurrentTracksValidNow).Methods("GET")
	rCurrentTracks.HandleFunc("/later", handlers.GetCurrentTracksValidLater).Methods("GET")
	rCurrentTracks.HandleFunc("/earlier", handlers.GetCurrentTracksValidEarlier).Methods("GET")

	// Cached tracks
	////////////////
	// Router
	rCachedTracks := r.PathPrefix("/cached").Subrouter()
	// GET
	rCachedTracks.HandleFunc("", handlers.GetAllCachedTracks).Methods("GET")
	rCachedTracks.HandleFunc("/", handlers.GetAllCachedTracks).Methods("GET")
	rCachedTracks.HandleFunc("/{track_id}", handlers.GetCachedTrack).Methods("GET")
	rCachedTracks.HandleFunc("/days/{days_old}", handlers.GetCachedTracksByDaysOld).Methods("GET")
	rCachedTracks.HandleFunc("/eastbound", handlers.GetCachedEastboundTracks).Methods("GET")
	rCachedTracks.HandleFunc("/westbound", handlers.GetCachedWestboundTracks).Methods("GET")
	rCachedTracks.HandleFunc("/check/{track_id}", handlers.CheckIsTrackCached).Methods("GET")

	// Event tracks
	///////////////
	// Router
	rEventTracks := r.PathPrefix("/event").Subrouter()
	// GET
	rEventTracks.HandleFunc("", handlers.GetAllEventTracks).Methods("GET")
	rEventTracks.HandleFunc("/", handlers.GetAllEventTracks).Methods("GET")
	rEventTracks.HandleFunc("/{track_id}", handlers.GetEventTrack).Methods("GET")
	// POST
	rEventTracks.HandleFunc("/{track_obj}", handlers.PostEventTrack).Methods("POST")
	// DELETE
	rEventTracks.HandleFunc("/{track_id}", handlers.DeleteEventTrack).Methods("DELETE")

	// Concorde tracks
	//////////////////
	rConcordeTracks := r.PathPrefix("/concorde").Subrouter()
	// GET
	rConcordeTracks.HandleFunc("", handlers.GetAllConcordeTracks).Methods("GET")
	rConcordeTracks.HandleFunc("/", handlers.GetAllConcordeTracks).Methods("GET")
	rConcordeTracks.HandleFunc("/{track_id}", handlers.GetConcordeTrack).Methods("GET")

	// Fixes
	////////
	// Router
	rFixes := r.PathPrefix("/fixes").Subrouter()
	// GET
	rFixes.HandleFunc("", handlers.GetAllFixes).Methods("GET")
	rFixes.HandleFunc("/", handlers.GetAllFixes).Methods("GET")
	rFixes.HandleFunc("/{fix_name}", handlers.GetFix).Methods("GET")
	// POST
	rFixes.HandleFunc("/{fix_obj}", handlers.PostFix).Methods("POST")
	// UPDATE
	rFixes.HandleFunc("/{fix_obj}", handlers.UpdateFix).Methods("UPDATE")
	// DELETE
	rFixes.HandleFunc("/{fix_name}", handlers.DeleteFix).Methods("DELETE")

	// Return the parent router
	return r
}

// Run the server
func (s Server) Run() error {
	// Initialise router
	router := s.Initialise()

	// Now serve
	err := http.ListenAndServe(":5000", router)

	// Error
	return err
}
