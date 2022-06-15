package handlers

import "net/http"

// GET /
func Index(w http.ResponseWriter, r *http.Request) {
	// Redirect to info page
	http.Redirect(w, r, "https://ganderoceanic.ca/nat-track-api-usage", http.StatusPermanentRedirect)
}
