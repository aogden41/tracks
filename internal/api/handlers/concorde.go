package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/aogden41/tracks/internal/db"
	"github.com/gorilla/mux"
)

// GetAllConcordeTracks godoc
// @Summary Get all concorde tracks
// @Description JSON output of all concorde tracks available for use
// @Tags concorde
// @Produce json
// @Success 200 {array} models.Track
// @Failure 500 {object} InternalServerError
// @Router /concorde [get]
func GetAllConcordeTracks(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Fetch tracks and check error
	tracks, err := db.SelectConcordeTracks()
	if err != nil {
		json.NewEncoder(w).Encode(Error500(&w, err.Error()))
		return
	}

	// Encode
	json.NewEncoder(w).Encode(tracks)
}

// GetConcordeTrack godoc
// @Summary Get one concorde track
// @Description JSON output of one specific concorde track
// @Tags concorde
// @Produce json
// @Param track_id path string true "SM, SN, SO or SP"
// @Success 200 {object} models.Track
// @Failure 404 {object} NotFound
// @Failure 500 {object} InternalServerError
// @Router /concorde/{track_id} [get]
func GetConcordeTrack(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Fetch tracks and check error
	params := mux.Vars(r)
	track, err := db.SelectConcordeTrack(params["track_id"])
	if err != nil {
		if err == sql.ErrNoRows {
			json.NewEncoder(w).Encode(Error404(&w, "Requested track was not found."))
		} else {
			json.NewEncoder(w).Encode(Error500(&w, err.Error()))
		}
		return
	}

	// Encode
	json.NewEncoder(w).Encode(track)
}
