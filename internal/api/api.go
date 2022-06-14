package api

import (
	"encoding/json"
	"github.com/aogden41/tracks/internal/db"
	"github.com/aogden41/tracks/internal/tracks"
	"net/http"
	"strconv"
)

///
/// INDEX/OTHERS
///

// GET /
func Index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://ganderoceanic.com/nat-track-api-usage", http.StatusPermanentRedirect)
}

///
/// CURRENT TRACKS
/// TODO should I cache current tracks?

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
	// Set json header
	w.Header().Set("Content-Type", "application/json")
}

// GET /current/eastbound
func GetCurrentEastboundTracks(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")
}

// GET /current/westbound
func GetCurrentWestboundTracks(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")
}

// GET /current/now
func GetCurrentTracksValidNow(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")
}

// GET /current/later
func GetCurrentTracksValidLater(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")
}

// GET /current/earlier
func GetCurrentTracksValidEarlier(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")
}

///
/// CACHED TRACKS
///

// GET /cached
func GetAllCachedTracks(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")
}

// GET /cached/{track_id}
func GetCachedTrack(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")
}

// GET /cached/days/{days_old}
func GetCachedTracksByDaysOld(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")
}

// GET /cached/eastbound
func GetCachedEastboundTracks(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")
}

// GET /cached/westbound
func GetCachedWestboundTracks(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")
}

// GET /cached/check/{track_id}
func CheckIsTrackCached(w http.ResponseWriter, r *http.Request) {

}

///
/// EVENT TRACKS
///

// GET /event
func GetAllEventTracks(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")

}

// GET /event/{track_id}
func GetEventTrack(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")
}

// POST /event/{track_obj}
func PostEventTrack(w http.ResponseWriter, r *http.Request) {

}

// DELETE /event/{track_id}
func DeleteEventTrack(w http.ResponseWriter, r *http.Request) {

}

///
/// CONCORDE TRACKS
///

// GET /concorde
func GetAllConcordeTracks(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")

}

// GET /concorde/{track_id}
func GetConcordeTrack(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")
}

///
/// FIXES
///

// GET /fixes
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

// GET /fixes/{fix_name}
func GetFix(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")
}

// POST /fixes/{fix_obj}
func PostFix(w http.ResponseWriter, r *http.Request) {

}

// UPDATE /fixes/{fix_obj}
func UpdateFix(w http.ResponseWriter, r *http.Request) {

}

// DELETE /fixes/{fix_name}
func DeleteFix(w http.ResponseWriter, r *http.Request) {

}
