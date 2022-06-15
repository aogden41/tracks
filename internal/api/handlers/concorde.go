package handlers

import "net/http"

// GET /concorde
func GetAllConcordeTracks(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "This function has not yet been implemented.", http.StatusNotImplemented)

}

// GET /concorde/{track_id}
func GetConcordeTrack(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "This function has not yet been implemented.", http.StatusNotImplemented)
}
