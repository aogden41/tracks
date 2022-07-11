package api

import (
	"github.com/aogden41/tracks/internal/db"
	"github.com/aogden41/tracks/internal/db/models"
	"github.com/aogden41/tracks/internal/tracks"
	"strconv"
	"unicode"
)

func CompareMessage(server *Server) {
	msg := tracks.DownloadTracks()

	// Check if what we have matches what we've just obtained
	if server.StoredMessage != msg {

		// Update the stored message
		server.StoredMessage = msg

		// Run job
		CacheJob(msg)

		return
	}
}

func CacheJob(newMsg string) error {
	// Get all tracks parsed from this current message
	tracks, err := tracks.ParseTracks(false, models.UNKNOWN, models.NA)
	if err != nil {
		return err
	}

	// Store current TMI for later
	var currentTMINumeric int = -1

	// Let's cache them
	for _, track := range tracks {
		// Set TMI
		if currentTMINumeric == -1 {
			if !unicode.IsDigit(rune(track.TMI[len(track.TMI)-1])) {
				currentTMINumeric, _ = strconv.Atoi(track.TMI[:len(track.TMI)-1])
			} else {
				currentTMINumeric, _ = strconv.Atoi(track.TMI)
			}
		}

		// Insert the track
		db.InsertCachedTrack(&track)
	}

	// Check now for tracks in the cache out of date by more than 7 days
	tracks, err = db.SelectCachedTracks()
	if err != nil {
		return err
	}
	for _, track := range tracks {
		// Strip any alpha runes off the end in case of amendments
		trackTMI := track.TMI
		if !unicode.IsDigit(rune(trackTMI[len(trackTMI)-1])) {
			trackTMI = trackTMI[:len(trackTMI)-1]
		}

		// Convert TMI to integer
		tmiInt, _ := strconv.Atoi(trackTMI)

		// Finally compare them
		if (currentTMINumeric - tmiInt) > 7 {
			db.DeleteCachedTrack(track.ID)
		}
	}

	return nil
}
