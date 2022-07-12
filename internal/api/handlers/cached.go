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

// GetCachedTracks godoc
// @Summary Get all cached tracks with a given TMI
// @Description JSON output of all cached tracks with the requested TMI
// @Tags cached
// @Produce json
// @Param tmi path string true "Track Message Identifier (Julian calendar day numbered 1 to 365, incl. any amendment characters)"
// @Success 200 {array} models.Track
// @Failure 404 {object} NotFound
// @Failure 500 {object} InternalServerError
// @Router /cached/{tmi} [get]
func GetCachedTracks(w http.ResponseWriter, r *http.Request) {
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
	params := mux.Vars(r)
	tracks, err := db.SelectCachedTracksByTMI(params["tmi"], models.UNKNOWN)
	if err != nil {
		if err == sql.ErrNoRows {
			json.NewEncoder(w).Encode(Error404(&w, "No tracks matching the query criteria were found."))
		} else {
			json.NewEncoder(w).Encode(Error500(&w, err.Error()))
		}
		return
	}

	// No error but check not empty
	if len(tracks) == 0 {
		json.NewEncoder(w).Encode(Error404(&w, "No tracks matching the query criteria were found."))
		return
	}

	// Encode
	json.NewEncoder(w).Encode(tracks)
}

// GetCachedEastboundTracks godoc
// @Summary Get all cached eastbound tracks with a given TMI
// @Description JSON output of all cached eastbound tracks with the requested TMI
// @Tags cached
// @Produce json
// @Param tmi path string true "Track Message Identifier (Julian calendar day numbered 1 to 365, incl. any amendment characters)"
// @Success 200 {array} models.Track
// @Failure 404 {object} NotFound
// @Failure 500 {object} InternalServerError
// @Router /cached/eastbound/{tmi} [get]
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
	params := mux.Vars(r)
	tracks, err := db.SelectCachedTracksByTMI(params["tmi"], models.EAST)
	if err != nil {
		if err == sql.ErrNoRows {
			json.NewEncoder(w).Encode(Error404(&w, "No tracks matching the query criteria were found."))
		} else {
			json.NewEncoder(w).Encode(Error500(&w, err.Error()))
		}
		return
	}

	// No error but check not empty
	if len(tracks) == 0 {
		json.NewEncoder(w).Encode(Error404(&w, "No tracks matching the query criteria were found."))
		return
	}

	// Encode
	json.NewEncoder(w).Encode(tracks)
}

// GetCachedWestboundTracks godoc
// @Summary Get all cached westbound tracks with a given TMI
// @Description JSON output of all cached westbound tracks with the requested TMI
// @Tags cached
// @Produce json
// @Param tmi path string true "Track Message Identifier (Julian calendar day numbered 1 to 365, incl. any amendment characters)"
// @Success 200 {array} models.Track
// @Failure 404 {object} NotFound
// @Failure 500 {object} InternalServerError
// @Router /cached/westbound/{tmi} [get]
func GetCachedWestboundTracks(w http.ResponseWriter, r *http.Request) {
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
	params := mux.Vars(r)
	tracks, err := db.SelectCachedTracksByTMI(params["tmi"], models.WEST)
	if err != nil {
		if err == sql.ErrNoRows {
			json.NewEncoder(w).Encode(Error404(&w, "No tracks matching the query criteria were found."))
		} else {
			json.NewEncoder(w).Encode(Error500(&w, err.Error()))
		}
		return
	}

	// No error but check not empty
	if len(tracks) == 0 {
		json.NewEncoder(w).Encode(Error404(&w, "No tracks matching the query criteria were found."))
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
