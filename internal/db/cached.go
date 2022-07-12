package db

import (
	"database/sql"
	"github.com/aogden41/tracks/internal/db/models"
	"github.com/aogden41/tracks/internal/tracks/track_utils"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

func SelectCachedTMIs() ([]string, error) {
	// Connect and defer
	db := Connect()
	defer db.Close()

	// Statement
	query := `SELECT tmi
				FROM tracks.cache 
				WHERE type <> 1 AND type <> 2;`

	// Perform query
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Array of return values
	tmis := []string{}

	// Iterate and assign
	lastTMI := ""
	for rows.Next() {
		rowTMI := ""
		if err := rows.Scan(&rowTMI); err != nil {
			return nil, err
		}
		// Compare row TMI to the last tmi
		if rowTMI != lastTMI {
			lastTMI = rowTMI
			tmis = append(tmis, lastTMI)
		}
	}

	// Catch any other error
	if err = rows.Err(); err != nil {
		return nil, err
	}

	// Success, return everything
	return tmis, nil
}

func SelectCachedTracksByTMI(tmi string, dir models.Direction) (map[string]models.Track, error) {
	// Connect and defer
	db := Connect()
	defer db.Close()

	// Statement
	query := `SELECT track_id, tmi, route, flight_levels, direction, valid_from, valid_to 
				FROM tracks.cache 
				WHERE type <> 1 AND type <> 2 AND tmi LIKE $1;`

	// Perform query
	rows, err := db.Query(query, tmi)
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

		// Check direction, throw away if parameter provided and not a match
		if dir != models.UNKNOWN && direction != int(dir) {
			continue
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
		track.Type = models.REGULAR

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

func SelectCachedTrack(tmi string, tk string) (models.Track, error) {
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
	row := db.QueryRow(query, tmi, tk)
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
	track.Type = models.REGULAR

	// Return selection
	return track, nil
}

func InsertCachedTrack(track *models.Track) error {
	// Connect and defer
	db := Connect()
	defer db.Close()

	// Statement (we don't insert if the cached track already exists in the DB)
	query := `INSERT INTO tracks.cache (track_id, tmi, route, flight_levels, direction, valid_from, valid_to, type) 
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
				ON CONFLICT ON CONSTRAINT tk_tmi DO NOTHING RETURNING id;`

	// Stringify route
	var route string
	for _, fix := range track.Route {
		route += fix.Name + " "
	}
	route = route[:len(route)-1] // Get rid of the last whitespace

	// Stringify flight levels
	var levels string
	for _, level := range track.FlightLevels {
		levels += strconv.Itoa(level) + " "
	}
	levels = levels[:len(levels)-1] // Get rid of the last whitespace

	// Complete query, check error and return
	rowId := 0 // Need this for auto increment but we discard it straight after
	err := db.QueryRow(query, track.ID, track.TMI, route, levels, track.Direction, track.ValidFrom, track.ValidTo, track.Type).Scan(&rowId)
	switch err {
	case sql.ErrNoRows: // Added successfully entry, send status 200
		return nil
	case nil: // No error
		return nil
	default:
		return err
	}
}

func DeleteCachedTrack(trackID string) error {
	// Connect and defer
	db := Connect()
	defer db.Close()

	// Statement
	query := `DELETE FROM tracks.cache WHERE track_id LIKE $1;`

	// Complete query, check error and return
	_, err := db.Exec(query, strings.ToUpper(trackID))
	if err != nil {
		return err
	}
	return nil
}
