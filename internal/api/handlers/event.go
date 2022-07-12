package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/aogden41/tracks/internal/db"
	"github.com/gorilla/mux"
)

// GetAllEventTracks godoc
// @Summary Get all event tracks with a given TMI
// @Description JSON output of all event tracks per an assigned TMI
// @Tags event
// @Produce json
// @Param event_tmi path string true "The TMI assigned to the event"
// @Success 200 {array} models.Track
// @Failure 404 {object} NotFound
// @Failure 500 {object} InternalServerError
// @Router /event/{event_tmi} [get]
func GetAllEventTracks(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Fetch fixes and check error
	params := mux.Vars(r)
	tracks, err := db.SelectEventTracks(strings.ToUpper(params["event_tmi"]))
	if err != nil {
		if err == sql.ErrNoRows {
			json.NewEncoder(w).Encode(Error404(&w, "No tracks matching the query criteria were found."))
		} else {
			json.NewEncoder(w).Encode(Error500(&w, err.Error()))
		}
		return
	}

	// Encode
	json.NewEncoder(w).Encode(tracks)

}

// GetEventTrack godoc
// @Summary Get one event track
// @Description JSON output of a specific event track per an assigned TMI and its ID
// @Tags event
// @Produce json
// @Param event_tmi path string true "The TMI assigned to the event"
// @Param track_id path string true "The requested track ID"
// @Success 200 {object} models.Track
// @Failure 404 {object} NotFound
// @Failure 500 {object} InternalServerError
// @Router /event/{event_tmi}/{track_id} [get]
func GetEventTrack(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Get parameters, request data and encode
	params := mux.Vars(r)
	track, err := db.SelectEventTrack(strings.ToUpper(params["event_tmi"]), strings.ToUpper(params["track_id"]))
	if err != nil {
		if err == sql.ErrNoRows {
			json.NewEncoder(w).Encode(Error404(&w, "Requested track does not exist"))
		} else {
			json.NewEncoder(w).Encode(Error500(&w, err.Error()))
		}
		return
	}

	// Encode
	json.NewEncoder(w).Encode(track)
}

// POST /event/{track_obj}
func PostEventTrack(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Error
	json.NewEncoder(w).Encode(Error501(&w, "This operation has not yet been implemented."))
}

// PUT /event/{track_obj}
func PutEventTrack(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Error
	json.NewEncoder(w).Encode(Error501(&w, "This operation has not yet been implemented."))
}

// DELETE /event/{track_id}
func DeleteEventTrack(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Error
	json.NewEncoder(w).Encode(Error501(&w, "This operation has not yet been implemented."))
}
