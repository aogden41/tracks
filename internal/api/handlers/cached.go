package handlers

import "net/http"

// GET /cached
func GetAllCachedTracks(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "This function has not yet been implemented.", http.StatusNotImplemented)
}

// GET /cached/{track_id}
func GetCachedTrack(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "This function has not yet been implemented.", http.StatusNotImplemented)
}

// GET /cached/days/{days_old}
func GetCachedTracksByDaysOld(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "This function has not yet been implemented.", http.StatusNotImplemented)
}

// GET /cached/eastbound
func GetCachedEastboundTracks(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "This function has not yet been implemented.", http.StatusNotImplemented)
}

// GET /cached/westbound
func GetCachedWestboundTracks(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "This function has not yet been implemented.", http.StatusNotImplemented)
}

// GET /cached/check/{track_id}
func CheckIsTrackCached(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "This function has not yet been implemented.", http.StatusNotImplemented)
}
