package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/aogden41/tracks/internal/tracks"
)

// GET /current
func GetAllCurrentTracks(w http.ResponseWriter, r *http.Request) {
	// SI units?
	isMetres := true // Default
	siVal, err := strconv.ParseBool(r.URL.Query().Get("si"))
	if err != nil || !siVal { // If not
		isMetres = false
	}

	// Set json header
	w.Header().Set("Content-Type", "application/json")

	// Parse tracks and return
	tracks, err := tracks.ParseTracks(isMetres)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// Encode and return
	json.NewEncoder(w).Encode(&tracks)
}

// GET /current/{track_id}
func GetCurrentTrack(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "This function has not yet been implemented.", http.StatusNotImplemented)
}

// GET /current/eastbound
func GetCurrentEastboundTracks(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "This function has not yet been implemented.", http.StatusNotImplemented)
}

// GET /current/westbound
func GetCurrentWestboundTracks(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "This function has not yet been implemented.", http.StatusNotImplemented)
}

// GET /current/now
func GetCurrentTracksValidNow(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "This function has not yet been implemented.", http.StatusNotImplemented)
}

// GET /current/later
func GetCurrentTracksValidLater(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "This function has not yet been implemented.", http.StatusNotImplemented)
}

// GET /current/earlier
func GetCurrentTracksValidEarlier(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "This function has not yet been implemented.", http.StatusNotImplemented)
}
