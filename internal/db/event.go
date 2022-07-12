package db

import (
	"strconv"
	"strings"

	"github.com/aogden41/tracks/internal/db/models"
	"github.com/aogden41/tracks/internal/tracks/track_utils"

	_ "github.com/lib/pq"
)

func SelectEventTracks(event_tmi string) (map[string]models.Track, error) {
	// Connect and defer
	db := Connect()
	defer db.Close()

	// Statement
	query := `SELECT track_id, tmi, route, flight_levels, direction, valid_from, valid_to 
				FROM tracks.cache 
				WHERE type LIKE 2 AND tmi LIKE $1;`

	// Perform query
	rows, err := db.Query(query, event_tmi)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Get all fixes
	fixes, err := SelectFixes()
	if err != nil {
		panic(err)
	}

	// Map of return values
	returnTracks := make(map[string]models.Track)

	// Iterate and assign
	for rows.Next() {
		// Extract values and error check
		var altitudes string
		var direction int
		var route string
		var track models.Track
		if err := rows.Scan(&track.ID, &track.TMI, &route, &altitudes, &direction, &track.ValidFrom, &track.ValidTo); err != nil {
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
		for _, fix := range strings.Split(route, " ") {
			if selectedFix, ok := fixes[fix]; ok {
				// Append
				track.Route = append(track.Route, selectedFix)
			} else {
				var newCoord models.Fix
				newCoord.Name = fix
				coordinates, _ := track_utils.ParseSlashedCoordinate(fix)
				newCoord.Latitude = coordinates[0]
				newCoord.Longitude = coordinates[1]

				// Append
				track.Route = append(track.Route, newCoord)
			}
		}

		// Parse and append direction
		track.Direction = models.Direction(direction)

		// Other variables
		track.Type = models.EVENT

		// Compile track and add to map
		returnTracks[track.ID] = track
	}

	// Catch any other error
	if err = rows.Err(); err != nil {
		return nil, err
	}

	// Success, return everything
	return returnTracks, nil
}

func SelectEventTrack(event_tmi string, tk string) (models.Track, error) {
	// Connect and defer
	db := Connect()
	defer db.Close()

	// Get all concorde fixes
	fixes, err := SelectFixes()
	if err != nil {
		panic(err)
	}

	// Statement
	query := `SELECT track_id, tmi, route, flight_levels, direction, valid_from, valid_to FROM tracks.cache WHERE tmi LIKE $1 AND track_id LIKE $2;`

	// Get the row and error check
	row := db.QueryRow(query, event_tmi, tk)
	var altitudes string
	var direction int
	var route string
	var track models.Track
	if err := row.Scan(&track.ID, &track.TMI, &route, &altitudes, &direction, &track.ValidFrom, &track.ValidTo); err != nil {
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
	for _, fix := range strings.Split(route, " ") {
		if selectedFix, ok := fixes[fix]; ok {
			// Append
			track.Route = append(track.Route, selectedFix)
		} else {
			var newCoord models.Fix
			newCoord.Name = fix
			coordinates, _ := track_utils.ParseSlashedCoordinate(fix)
			newCoord.Latitude = coordinates[0]
			newCoord.Longitude = coordinates[1]
			newCoord.IsValid = true

			// Append
			track.Route = append(track.Route, newCoord)
		}
	}

	// Parse and append direction
	track.Direction = models.Direction(direction)

	// Other variables
	track.Type = models.EVENT

	// Return selection
	return track, nil
}
