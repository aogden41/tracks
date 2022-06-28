package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/aogden41/tracks/internal/db"
	"github.com/gorilla/mux"
)

// GET /concorde
func GetAllConcordeTracks(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")

	// Fetch tracks and check error
	tracks, err := db.SelectConcordeTracks()
	if err != nil {
		panic(err)
	}

	// Encode
	json.NewEncoder(w).Encode(tracks)
}

// GET /concorde/{track_id}
func GetConcordeTrack(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")

	// Fetch tracks and check error
	params := mux.Vars(r)
	track, err := db.SelectConcordeTrack(params["track_id"])
	if err != nil {
		panic(err)
	}

	// Encode
	json.NewEncoder(w).Encode(track)
}
