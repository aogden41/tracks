package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/aogden41/tracks/internal/db"
)

// GET /cached
func GetAllCachedTracks(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")

	// Fetch fixes and check error
	tracks, err := db.SelectCachedTracks()
	if err != nil {
		panic(err)
	}

	// Encode
	json.NewEncoder(w).Encode(tracks)
}

// GET /cached/{track_id}
func GetCachedTrack(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "This function has not yet been implemented.", http.StatusNotImplemented)
}

// GET /cached/days/{days_old}
func GetCachedTracksByDaysOld(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "This function has not yet been implemented.", http.StatusNotImplemented)
}

// GET /cached/eastbound
func GetCachedEastboundTracks(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "This function has not yet been implemented.", http.StatusNotImplemented)
}

// GET /cached/westbound
func GetCachedWestboundTracks(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "This function has not yet been implemented.", http.StatusNotImplemented)
}

// GET /cached/check/{track_id}
func CheckIsTrackCached(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "This function has not yet been implemented.", http.StatusNotImplemented)
}
