package handlers

import (
	"encoding/json"
	"strings"

	"net/http"
	"strconv"

	"github.com/aogden41/tracks/internal/db/models"
	"github.com/aogden41/tracks/internal/tracks"
	"github.com/gorilla/mux"
)

// GET /current
func GetAllCurrentTracks(w http.ResponseWriter, r *http.Request) {
	// SI units?
	isMetres := true // Default
	si, err := strconv.ParseBool(r.URL.Query().Get("si"))
	if err != nil || !si { // If not
		isMetres = false
	}

	// Set json header
	w.Header().Set("Content-Type", "application/json")

	// Parse tracks and return
	tracks, err := tracks.ParseTracks(isMetres, models.UNKNOWN)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// Encode and return
	json.NewEncoder(w).Encode(tracks)
}

// GET /current/{track_id}
func GetCurrentTrack(w http.ResponseWriter, r *http.Request) {
	// SI units?
	isMetres := true
	si, err := strconv.ParseBool(r.URL.Query().Get("si"))
	if err != nil || !si {
		// SI units not requested
		isMetres = false
	}

	// Set json header
	w.Header().Set("Content-Type", "application/json")

	// Get the track ID parameter and isolate the fix we need
	params := mux.Vars(r)
	tid := strings.ToUpper(params["track_id"])

	// Deserialise and display
	tracks, err := tracks.ParseTracks(isMetres, models.UNKNOWN)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// Encode and return
	json.NewEncoder(w).Encode(tracks[tid])
}

// GET /current/eastbound
func GetCurrentEastboundTracks(w http.ResponseWriter, r *http.Request) {
	// SI units?
	isMetres := true // Default
	si, err := strconv.ParseBool(r.URL.Query().Get("si"))
	if err != nil || !si { // If not
		isMetres = false
	}

	// Set json header
	w.Header().Set("Content-Type", "application/json")

	// Parse tracks and return
	tracks, err := tracks.ParseTracks(isMetres, models.EAST)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// Encode and return
	json.NewEncoder(w).Encode(tracks)
}

// GET /current/westbound
func GetCurrentWestboundTracks(w http.ResponseWriter, r *http.Request) {
	// SI units?
	isMetres := true // Default
	si, err := strconv.ParseBool(r.URL.Query().Get("si"))
	if err != nil || !si { // If not
		isMetres = false
	}

	// Set json header
	w.Header().Set("Content-Type", "application/json")

	// Parse tracks and return
	tracks, err := tracks.ParseTracks(isMetres, models.WEST)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// Encode and return
	json.NewEncoder(w).Encode(tracks)
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
