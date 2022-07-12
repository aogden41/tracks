package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/aogden41/tracks/internal/db"
	"github.com/aogden41/tracks/internal/db/models"
	"github.com/gorilla/mux"
)

// GetAllCachedTracks godoc
// @Summary Get all cached tracks
// @Description JSON output of all tracks cached in the API database (except Concorde and event)
// @Tags cached
// @Produce json
// @Success 200 {array} models.Track
// @Failure 404 {object} NotFound
// @Failure 500 {object} InternalServerError
// @Router /cached [get]
func GetAllCachedTracks(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Fetch fixes and check error
	tracks, err := db.SelectCachedTracks(models.UNKNOWN)
	if err != nil {
		if err == sql.ErrNoRows {
			json.NewEncoder(w).Encode(Error404(&w, "No tracks were found in the cache"))
		} else {
			json.NewEncoder(w).Encode(Error500(&w, err.Error()))
		}
		return
	}

	// Encode
	json.NewEncoder(w).Encode(tracks)
}

// GetCachedTrack godoc
// @Summary Get one cached track by ID and TMI
// @Description JSON output of a specific track cached in the API database
// @Tags cached
// @Produce json
// @Param tmi path string true "Track Message Identifier (Julian calendar day numbered 1 to 365, incl. any amendment characters)"
// @Param track_id path string true "Track ID (letter from A-Z)"
// @Success 200 {object} models.Track
// @Failure 404 {object} NotFound
// @Failure 500 {object} InternalServerError
// @Router /cached/{tmi}/{track_id} [get]
func GetCachedTrack(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Get parameters, request data and encode
	params := mux.Vars(r)
	track, err := db.SelectCachedTrack(strings.ToUpper(params["tmi"]), strings.ToUpper(params["track_id"]))
	if err != nil {
		if err == sql.ErrNoRows {
			json.NewEncoder(w).Encode(Error404(&w, "Requested track was not found in cache."))
		} else {
			json.NewEncoder(w).Encode(Error500(&w, err.Error()))
		}
		return
	}

	// Encode
	json.NewEncoder(w).Encode(track)
}

// GetCachedEastboundTracks godoc
// @Summary Get all cached eastbound tracks
// @Description JSON output of all eastbound tracks cached in the API database
// @Tags cached
// @Produce json
// @Success 200 {array} models.Track
// @Failure 404 {object} NotFound
// @Failure 500 {object} InternalServerError
// @Router /cached/eastbound [get]
func GetCachedEastboundTracks(w http.ResponseWriter, r *http.Request) {
	// SI units?
	/*isMetres := true // Default
	si, err := strconv.ParseBool(r.URL.Query().Get("si"))
	if err != nil || !si { // If not
		isMetres = false
	}*/

	// Set json header
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Fetch fixes and check error
	tracks, err := db.SelectCachedTracks(models.EAST)
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

// GetCachedWestboundTracks godoc
// @Summary Get all cached westbound tracks
// @Description JSON output of all westbound tracks cached in the API database
// @Tags cached
// @Produce json
// @Success 200 {array} models.Track
// @Failure 404 {object} NotFound
// @Failure 500 {object} InternalServerError
// @Router /cached/westbound [get]
func GetCachedWestboundTracks(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Fetch fixes and check error
	tracks, err := db.SelectCachedTracks(models.WEST)
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

// CheckIsTrackCached godoc
// @Summary Check cached status of a given track in the database. Warning: Intensive operation
// @Description Determine whether, given a particular track ID and a route, if the track is cached in the database
// @Tags cached
// @Produce json
// @Param track_id path string true "Track ID (letter from A-Z)"
// @Param route query string true "Route string as returned by other API requests (POINT XX/XX XX/XX ...)"
// Success 200 {object} RequestOK "If the track has been found, the TMI will be returned in the OK object message"
// Failure 404 {object} NotFound
// Failure 500 {object} InternalServerError
// @Failure 501 {object} NotImplemented
// @Router /cached/check/{track_id} [get]
func CheckIsTrackCached(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// Error
	json.NewEncoder(w).Encode(Error501(&w, "This operation has not yet been implemented."))

}
