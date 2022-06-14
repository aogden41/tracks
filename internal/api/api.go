package api

import (
	"encoding/json"
	"github.com/aogden41/tracks/internal/db"
	"github.com/aogden41/tracks/internal/tracks"
	"net/http"
	"strconv"
)

// Route  "/"
func Index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://ganderoceanic.com/nat-track-api-usage", http.StatusPermanentRedirect)
}

///
/// NORMAL TRACKS
///

// Route "/data"
func GetAllTracks(w http.ResponseWriter, r *http.Request) {
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

// Route "/data/{track_id}"
func GetOneTrack(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")
}

func UpdateOneTrack(w http.ResponseWriter, r *http.Request) {

}

func DeleteOneTrack(w http.ResponseWriter, r *http.Request) {

}

///
/// EVENT TRACKS
///

// Route "/event"
func GetAllEventTracks(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")

}

// Route "/event/{track_id}"
func GetOneEventTrack(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")
}

func PostOneEventTrack(w http.ResponseWriter, r *http.Request) {

}

func PostManyEventTracks(w http.ResponseWriter, r *http.Request) {

}

func DeleteOneEventTrack(w http.ResponseWriter, r *http.Request) {

}

///
/// CONCORDE TRACKS
///

// Route "/concorde"
func GetAllConcordeTracks(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")

}

// Route "/concorde/{track_id}"
func GetOneConcordeTrack(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")
}

///
/// FIXES
///

// Route "/fixes"
func GetAllFixes(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")
	// Fetch fixes and check error
	fixes, err := db.SelectFixes()
	if err != nil {
		panic(err)
	}

	// Encode and return
	json.NewEncoder(w).Encode(fixes)
}

// Route "/fixes/{fix_name}"
func GetOneFix(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")
}

func PostOneFix(w http.ResponseWriter, r *http.Request) {

}

func UpdateOneFix(w http.ResponseWriter, r *http.Request) {

}
