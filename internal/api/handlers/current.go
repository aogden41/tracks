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

// GetAllCurrentTracks godoc
// @Summary Get all current tracks
// @Description JSON output of all tracks published in the current NAT message
// @Tags current
// @Produce json
// @Param si query boolean false "Parse altitudes in metres"
// @Success 200 {array} models.Track
// @Failure 500 {object} InternalServerError
// @Router /current [get]
func GetAllCurrentTracks(w http.ResponseWriter, r *http.Request) {
	// SI units?
	isMetres := true // Default
	si, err := strconv.ParseBool(r.URL.Query().Get("si"))
	if err != nil || !si { // If not
		isMetres = false
	}

	// Set json header
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Parse tracks and return
	tracks, err := tracks.ParseTracks(isMetres, models.UNKNOWN, models.NA)
	if err != nil {
		json.NewEncoder(w).Encode(Error500(&w, err.Error()))
		return
	}

	// Encode and return
	json.NewEncoder(w).Encode(tracks)
}

// GetCurrentTrack godoc
// @Summary Get one current track by ID
// @Description JSON output of a specific track published in the current NAT message
// @Tags current
// @Produce json
// @Param track_id path string true "Track ID (letter from A-Z)"
// @Param si query boolean false "Parse altitudes in metres"
// @Success 200 {object} models.Track
// @Failure 404 {object} NotFound
// @Failure 500 {object} InternalServerError
// @Router /current/{track_id} [get]
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
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Get the track ID parameter and isolate the fix we need
	params := mux.Vars(r)
	tid := strings.ToUpper(params["track_id"])

	// Deserialise and display
	tracks, err := tracks.ParseTracks(isMetres, models.UNKNOWN, models.NA)
	if err != nil {
		json.NewEncoder(w).Encode(Error500(&w, err.Error()))
		return
	}

	// Check if the track is actually in the map
	if tk, ok := tracks[tid]; ok {
		// Encode and return
		json.NewEncoder(w).Encode(tk)
	} else {
		// Not found
		json.NewEncoder(w).Encode(Error404(&w, "The requested track does not exist."))
	}
}

// GetCurrentEastboundTracks godoc
// @Summary Get all currently published eastbound tracks
// @Description JSON output of all eastbound tracks currently listed in the NAT message
// @Tags current
// @Produce json
// @Param si query boolean false "Parse altitudes in metres"
// @Success 200 {array} models.Track
// @Failure 404 {object} NotFound
// @Failure 500 {object} InternalServerError
// @Router /current/eastbound [get]
func GetCurrentEastboundTracks(w http.ResponseWriter, r *http.Request) {
	// SI units?
	isMetres := true // Default
	si, err := strconv.ParseBool(r.URL.Query().Get("si"))
	if err != nil || !si { // If not
		isMetres = false
	}

	// Set json header
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Parse tracks and return
	tracks, err := tracks.ParseTracks(isMetres, models.EAST, models.NA)
	if err != nil {
		json.NewEncoder(w).Encode(Error500(&w, err.Error()))
		return
	}

	// Check if empty
	if len(tracks) != 0 {
		// Encode and return
		json.NewEncoder(w).Encode(tracks)
	} else {
		// Not found
		json.NewEncoder(w).Encode(Error404(&w, "The requested tracks do not exist."))
	}
}

// GetCurrentWestboundTracks godoc
// @Summary Get all currently published westbound tracks
// @Description JSON output of all westbound tracks currently listed in the NAT message
// @Tags current
// @Produce json
// @Param si query boolean false "Parse altitudes in metres"
// @Success 200 {array} models.Track
// @Failure 404 {object} NotFound
// @Failure 500 {object} InternalServerError
// @Router /current/westbound [get]
func GetCurrentWestboundTracks(w http.ResponseWriter, r *http.Request) {
	// SI units?
	isMetres := true // Default
	si, err := strconv.ParseBool(r.URL.Query().Get("si"))
	if err != nil || !si { // If not
		isMetres = false
	}

	// Set json header
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Parse tracks and return
	tracks, err := tracks.ParseTracks(isMetres, models.WEST, models.NA)
	if err != nil {
		json.NewEncoder(w).Encode(Error500(&w, err.Error()))
	}

	// Check if empty
	if len(tracks) != 0 {
		// Encode and return
		json.NewEncoder(w).Encode(tracks)
	} else {
		// Not found
		json.NewEncoder(w).Encode(Error404(&w, "The requested tracks do not exist."))
	}
}

// GetCurrentTracksValidNow godoc
// @Summary Get all currently published tracks valid now
// @Description JSON output of all tracks currently listed in the NAT message which are valid now
// @Tags current
// @Produce json
// @Param si query boolean false "Parse altitudes in metres"
// @Success 200 {array} models.Track
// @Failure 404 {object} NotFound
// @Failure 500 {object} InternalServerError
// @Router /current/now [get]
func GetCurrentTracksValidNow(w http.ResponseWriter, r *http.Request) {
	// SI units?
	isMetres := true // Default
	si, err := strconv.ParseBool(r.URL.Query().Get("si"))
	if err != nil || !si { // If not
		isMetres = false
	}

	// Set json header
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Parse tracks and return
	tracks, err := tracks.ParseTracks(isMetres, models.UNKNOWN, models.NOW)
	if err != nil {
		json.NewEncoder(w).Encode(Error500(&w, err.Error()))
	}

	// Check if empty
	if len(tracks) != 0 {
		// Encode and return
		json.NewEncoder(w).Encode(tracks)
	} else {
		// Not found
		json.NewEncoder(w).Encode(Error404(&w, "The requested tracks do not exist."))
	}
}

// GetCurrentTracksValidLater godoc
// @Summary Get all currently published tracks valid later
// @Description JSON output of all tracks currently listed in the NAT message which are valid later
// @Tags current
// @Produce json
// @Param si query boolean false "Parse altitudes in metres"
// @Success 200 {array} models.Track
// @Failure 404 {object} NotFound
// @Failure 500 {object} InternalServerError
// @Router /current/later [get]
func GetCurrentTracksValidLater(w http.ResponseWriter, r *http.Request) {
	// SI units?
	isMetres := true // Default
	si, err := strconv.ParseBool(r.URL.Query().Get("si"))
	if err != nil || !si { // If not
		isMetres = false
	}

	// Set json header
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Parse tracks and return
	tracks, err := tracks.ParseTracks(isMetres, models.UNKNOWN, models.LATER)
	if err != nil {
		json.NewEncoder(w).Encode(Error500(&w, err.Error()))
	}

	// Check if empty
	if len(tracks) != 0 {
		// Encode and return
		json.NewEncoder(w).Encode(tracks)
	} else {
		// Not found
		json.NewEncoder(w).Encode(Error404(&w, "The requested tracks do not exist."))
	}
}

// GetCurrentTracksValidEarlier godoc
// @Summary Get all currently published tracks valid earlier
// @Description JSON output of all tracks currently listed in the NAT message which are valid earlier
// @Tags current
// @Produce json
// @Param si query boolean false "Parse altitudes in metres"
// @Success 200 {array} models.Track
// @Failure 404 {object} NotFound
// @Failure 500 {object} InternalServerError
// @Router /current/earlier [get]
func GetCurrentTracksValidEarlier(w http.ResponseWriter, r *http.Request) {
	// SI units?
	isMetres := true // Default
	si, err := strconv.ParseBool(r.URL.Query().Get("si"))
	if err != nil || !si { // If not
		isMetres = false
	}

	// Set json header
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Parse tracks and return
	tracks, err := tracks.ParseTracks(isMetres, models.UNKNOWN, models.EARLIER)
	if err != nil {
		json.NewEncoder(w).Encode(Error500(&w, err.Error()))
	}

	// Check if empty
	if len(tracks) != 0 {
		// Encode and return
		json.NewEncoder(w).Encode(tracks)
	} else {
		// Not found
		json.NewEncoder(w).Encode(Error404(&w, "The requested tracks do not exist."))
	}
}

// GetCurrentTMI godoc
// @Summary Get the current TMI
// @Description String output of the current TMI
// @Tags current
// @Produce json
// @Success 200 {object} string
// @Failure 500 {object} InternalServerError
// @Router /current/tmi [get]
func GetCurrentTMI(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Parse the tracks for the TMI
	tracks, _ := tracks.ParseTracks(false, models.UNKNOWN, models.NA)
	var tmi string

	// Dodgy but works
	for _, track := range tracks {
		// Set TMI
		tmi = track.TMI

		// Break after first iteration, we have what we need
		break
	}

	if len(tmi) != 0 {
		// Encode and return
		json.NewEncoder(w).Encode(tmi)
	} else {
		// Not found
		json.NewEncoder(w).Encode(Error500(&w, "Could not fetch TMI."))
	}
}
