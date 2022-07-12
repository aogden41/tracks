package api

import (
	"database/sql"
	"strconv"
	"unicode"

	"github.com/aogden41/tracks/internal/db"
	"github.com/aogden41/tracks/internal/db/models"
	"github.com/aogden41/tracks/internal/tracks"
)

func CompareMessage(server *Server) {
	tracks, _ := tracks.ParseTracks(false, models.UNKNOWN, models.NA)
	var tmi string

	// Dodgy but works
	for _, track := range tracks {
		// Set TMI
		tmi = track.TMI

		// Break after first iteration, we have what we need
		break
	}

	// Check if what we have matches what we've just obtained
	if server.CurrentTMI != tmi {

		// Update the stored TMI
		server.CurrentTMI = tmi

		// Run job
		CacheJob()

		return
	}
}

func CacheJob() error {
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
	tracks, err = db.SelectCachedTracksByTMI(strconv.Itoa(currentTMINumeric-8), models.UNKNOWN)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	// If there were tracks that returned then delete them
	if len(tracks) > 0 || err == sql.ErrNoRows {
		for _, track := range tracks {
			db.DeleteCachedTrack(track.ID)
		}
	}
	return nil
}
