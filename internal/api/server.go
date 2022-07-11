package api

import (
	"net/http"

	"github.com/aogden41/tracks/internal/api/handlers"
	"github.com/gorilla/mux"
	"github.com/robfig/cron/v3"
)

type Server struct {
	Router        *mux.Router
	StoredMessage string
}

// Index routes
func (s Server) RouteIndex(r *mux.Router) {
	// GET
	r.HandleFunc("", handlers.Index).Methods("GET")
	r.HandleFunc("/", handlers.Index).Methods("GET")
}

// Current routes
func (s Server) RouteCurrent(r *mux.Router) {
	// GET
	r.HandleFunc("", handlers.GetAllCurrentTracks).Methods("GET")
	r.HandleFunc("/", handlers.GetAllCurrentTracks).Methods("GET")
	r.HandleFunc("/eastbound", handlers.GetCurrentEastboundTracks).Methods("GET")
	r.HandleFunc("/westbound", handlers.GetCurrentWestboundTracks).Methods("GET")
	r.HandleFunc("/now", handlers.GetCurrentTracksValidNow).Methods("GET")
	r.HandleFunc("/later", handlers.GetCurrentTracksValidLater).Methods("GET")
	r.HandleFunc("/earlier", handlers.GetCurrentTracksValidEarlier).Methods("GET")
	r.HandleFunc("/{track_id}", handlers.GetCurrentTrack).Methods("GET")
}

// Cached routes
func (s Server) RouteCached(r *mux.Router) {
	// GET
	r.HandleFunc("", handlers.GetAllCachedTracks).Methods("GET")
	r.HandleFunc("/", handlers.GetAllCachedTracks).Methods("GET")
	r.HandleFunc("/eastbound", handlers.GetCachedEastboundTracks).Methods("GET")
	r.HandleFunc("/westbound", handlers.GetCachedWestboundTracks).Methods("GET")
	r.HandleFunc("/{track_id}", handlers.GetCachedTrack).Methods("GET")
	r.HandleFunc("/days/{days_old}", handlers.GetCachedTracksByDaysOld).Methods("GET")
	r.HandleFunc("/check/{track_id}", handlers.CheckIsTrackCached).Methods("GET")
}

// Event routes
func (s Server) RouteEvent(r *mux.Router) {
	// GET
	r.HandleFunc("", handlers.GetAllEventTracks).Methods("GET")
	r.HandleFunc("/", handlers.GetAllEventTracks).Methods("GET")
	r.HandleFunc("/{track_id}", handlers.GetEventTrack).Methods("GET")
	// POST
	r.HandleFunc("", handlers.PostEventTrack).Methods("POST")
	r.HandleFunc("/", handlers.PostEventTrack).Methods("POST")
	// PUT
	r.HandleFunc("", handlers.PutEventTrack).Methods("PUT")
	r.HandleFunc("/", handlers.PutEventTrack).Methods("PUT")
	// DELETE
	r.HandleFunc("/{track_id}", handlers.DeleteEventTrack).Methods("DELETE")
}

// Concorde routes
func (s Server) RouteConcorde(r *mux.Router) {
	// GET
	r.HandleFunc("", handlers.GetAllConcordeTracks).Methods("GET")
	r.HandleFunc("/", handlers.GetAllConcordeTracks).Methods("GET")
	r.HandleFunc("/{track_id}", handlers.GetConcordeTrack).Methods("GET")
}

// Fix routes
func (s Server) RouteFixes(r *mux.Router) {
	// GET
	r.HandleFunc("", handlers.GetAllFixes).Methods("GET")
	r.HandleFunc("/", handlers.GetAllFixes).Methods("GET")
	r.HandleFunc("/concorde", handlers.GetConcordeFixes).Methods("GET")
	r.HandleFunc("/{fix_name}", handlers.GetFix).Methods("GET")
	// POST
	r.HandleFunc("", handlers.PostFix).Methods("POST")
	r.HandleFunc("/", handlers.PostFix).Methods("POST")
	// PUT
	r.HandleFunc("", handlers.PutFix).Methods("PUT")
	r.HandleFunc("/", handlers.PutFix).Methods("PUT")
	// DELETE
	r.HandleFunc("/{fix_name}", handlers.DeleteFix).Methods("DELETE")
}

// Run the server
func (s Server) Run() error {
	// Initialise router
	s.Router = mux.NewRouter()

	// Get the first message
	s.StoredMessage = ""
	CompareMessage(&s)

	// Start routing
	s.RouteIndex(s.Router.PathPrefix("").Subrouter())
	s.RouteCurrent(s.Router.PathPrefix("/current").Subrouter())
	s.RouteCached(s.Router.PathPrefix("/cached").Subrouter())
	s.RouteEvent(s.Router.PathPrefix("/event").Subrouter())
	s.RouteConcorde(s.Router.PathPrefix("/concorde").Subrouter())
	s.RouteFixes(s.Router.PathPrefix("/fixes").Subrouter())

	// Every 10 minutes compare the message with what we have in memory
	cron := cron.New()
	cron.AddFunc("@every 10s", func() { CompareMessage(&s) })
	cron.Start()

	// Now serve
	err := http.ListenAndServe(":5000", s.Router)

	// Error
	return err
}
