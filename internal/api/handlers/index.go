package handlers

import "net/http"

// Index godoc
// @Summary API index page
// @Description Redirect to the Gander Oceanic page with API instructions and link to this documentation
// @Tags index
// @Success 308
// @Router / [get]
func Index(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Redirect to info page
	http.Redirect(w, r, "https://ganderoceanic.ca/nat-track-api-usage", http.StatusPermanentRedirect)
}
