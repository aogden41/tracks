package db

import (
	"strconv"
	"strings"

	"github.com/aogden41/tracks/internal/db/models"

	_ "github.com/lib/pq"
)

// TODO more efficient way of fetching fixes than iterating
func SelectConcordeTracks() (map[string]models.Track, error) {
	// Connect and defer
	db := Connect()
	defer db.Close()

	// Statement
	query := `SELECT track_id, tmi, flight_levels, direction, valid_from, valid_to FROM tracks.cache WHERE CAST(type AS TEXT) LIKE '1';`

	// Perform query
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Get all concorde fixes
	cFixes, err := SelectConcordeFixes()
	if err != nil {
		panic(err)
	}

	// Map of return values
	cTracks := make(map[string]models.Track)

	// Iterate and assign
	for rows.Next() {
		// Extract values and error check
		var altitudes string
		var direction int
		var track models.Track
		if err := rows.Scan(&track.ID, &track.TMI, &altitudes, &direction, &track.ValidFrom, &track.ValidTo); err != nil {
			return nil, err
		}

		// Parse and append the flight levels
		track.FlightLevels = []int{}
		for _, fl := range strings.Split(altitudes, " ") {
			flInt, _ := strconv.Atoi(fl)
			track.FlightLevels = append(track.FlightLevels, flInt)
		}

		// Parse and append route elements
		track.Route = []models.Fix{}
		for _, fix := range cFixes {
			if strings.Contains(fix.Name, track.ID) {
				track.Route = append(track.Route, fix)
			}
		}

		// Parse and append direction
		track.Direction = models.Direction(direction)

		// Other variables
		track.Type = models.CONCORDE
		track.TMI = "-1"

		// Compile track and add to map
		cTracks[track.ID] = track
	}

	return cTracks, nil
}

// TODO efficiency as well
func SelectConcordeTrack(tk string) (models.Track, error) {
	// Connect and defer
	db := Connect()
	defer db.Close()

	// Get all concorde fixes
	cFixes, err := SelectConcordeFixes()
	if err != nil {
		panic(err)
	}

	// Statement
	query := `SELECT track_id, tmi, flight_levels, direction, valid_from, valid_to FROM tracks.cache WHERE track_id LIKE $1;`

	// Get the row and error check
	row := db.QueryRow(query, strings.ToUpper(tk))
	var altitudes string
	var direction int
	var track models.Track
	if err := row.Scan(&track.ID, &track.TMI, &altitudes, &direction, &track.ValidFrom, &track.ValidTo); err != nil {
		return models.Track{}, err
	}

	// Parse and append the flight levels
	track.FlightLevels = []int{}
	for _, fl := range strings.Split(altitudes, " ") {
		flInt, _ := strconv.Atoi(fl)
		track.FlightLevels = append(track.FlightLevels, flInt)
	}

	// Parse and append route elements
	track.Route = []models.Fix{}
	for _, fix := range cFixes {
		if strings.Contains(fix.Name, track.ID) {
			track.Route = append(track.Route, fix)
		}
	}

	// Parse and append direction
	track.Direction = models.Direction(direction)

	// Other variables
	track.Type = models.CONCORDE
	track.TMI = "-1"

	// Return selection
	return track, nil
}
