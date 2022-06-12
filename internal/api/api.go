package api

import (
	"fmt"
	"io"
	"net/http"

	"github.com/aogden41/tracks/internal/tracks"
)

// Route  "/"
func Index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://ganderoceanic.com/nat-track-api-usage", http.StatusPermanentRedirect)
}

// Route "/data/"
func Get(w http.ResponseWriter, r *http.Request) {
	tracks.ParseTracks(false)
}

// Route "/event"
func GetEvent(w http.ResponseWriter, r *http.Request) {
	// Fetch data
	path := "https://ganderoceanicoca.ams3.cdn.digitaloceanspaces.com/resources/data/eventTracks.json"
	res, err := http.Get(path)

	// Check for errors then defer
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer res.Body.Close()

	// Read bytes and output
	bytes, _ := io.ReadAll(res.Body)
	fmt.Fprintf(w, string(bytes))
}

// Route "/concorde"
func GetConcorde(w http.ResponseWriter, r *http.Request) {
	// Fetch data
	path := "https://ganderoceanicoca.ams3.cdn.digitaloceanspaces.com/resources/data/concordeTracks.json"
	res, err := http.Get(path)

	// Check for errors then defer
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer res.Body.Close()

	// Read bytes and output
	bytes, _ := io.ReadAll(res.Body)
	fmt.Fprintf(w, string(bytes))
}
