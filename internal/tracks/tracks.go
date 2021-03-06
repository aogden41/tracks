package tracks

import (
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/aogden41/tracks/internal/db"
	"github.com/aogden41/tracks/internal/db/models"
	"github.com/aogden41/tracks/internal/tracks/track_utils"
)

// Constants
const trackUrl = "https://www.notams.faa.gov/common/nat.html"

// Months of the year
var months = [12]string{"JAN", "FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEP", "OCT", "NOV", "DEC"}

func DownloadTracks() string {
	// Download the tracks
	tracksRes, err := http.Get(trackUrl)

	// Handle any error from web requests
	if err != nil {
		panic(err)
	}

	// Fetching is done, defer closure until end of function
	defer tracksRes.Body.Close()

	// Get bytes and assign data
	bytes, _ := io.ReadAll(tracksRes.Body)
	return string(bytes) // Track message
}

// Parse the track data
func ParseTracks(isMetres bool, direction models.Direction, validity models.Validity) (map[string]models.Track, error) {
	// First get all fixes from the database and error check
	fixes, err := db.SelectFixes()
	if err != nil {
		return nil, err
	}

	// Get the tracks
	natData := DownloadTracks()

	// Split the data
	natDataLines := strings.Split(string(natData), "\n")

	// Store extracted track slices
	var trackSlices [][]string

	// Store data currently being processed
	var track []string
	var trackValidities map[string][2]int64 = make(map[string][2]int64)
	var tmi string
	var validFrom time.Time = time.Now() // Initial value
	var validTo time.Time = time.Now()   // Initial value

	// Flag to get the next 2 lines for a track
	getNextTwo := 0

	// Iterate through
	for i := 0; i < len(natDataLines); i++ {
		// Rune slice of current line
		lineRunes := []rune(natDataLines[i])

		// Continue if the line is just empty
		if len(natDataLines[i]) == 0 {
			continue
		}

		// If the row is a track (we assume this if the structure is a letter then a space then another 2 letters, or if a track is already being processed)
		if (unicode.IsLetter(lineRunes[0]) && unicode.IsSpace(lineRunes[1]) && unicode.IsLetter(lineRunes[2]) && unicode.IsLetter(lineRunes[3])) || getNextTwo > 0 {
			// If greater than 2 then we've collected a track, so process and move on
			if getNextTwo > 2 {
				getNextTwo = 0
				trackSlices = append(trackSlices, track)
				trackValidities[string(track[0][0])] = [2]int64{validFrom.Unix(), validTo.Unix()}
				track = []string{}
				continue
			}

			// Otherwise we add to the track being collected
			track = append(track, natDataLines[i])

			// Increment
			getNextTwo++

		} else if strings.Contains(natDataLines[i], "TMI IS") {
			// Extract the TMI
			tmiStart := strings.Index(natDataLines[i], "TMI IS") + 7
			tmi = string(natDataLines[i][tmiStart : tmiStart+3])

			// Add amendment character if it exists
			if unicode.IsLetter(lineRunes[tmiStart+3]) {
				tmi = tmi + string(natDataLines[i][tmiStart+3])
			}
		} else {
			// Convert validities
			for j := 0; j < len(months); j++ {
				reached := false

				// Check if the line contains a month
				if strings.HasPrefix(natDataLines[i], months[j]+" ") {
					// Split the to and from times
					splitString := strings.Split(natDataLines[i], " TO ")
					validFromSplit := strings.Split(splitString[0][4:], "/")
					validToSplit := strings.Split(splitString[1][4:], "/")

					// Get rid of any crap on the end
					if len(validToSplit[1]) > 5 {
						validToSplit = []string{validToSplit[0], validToSplit[1][:5]}
					}

					// Valid from
					validFromDay, _ := strconv.Atoi(validFromSplit[0])
					validFromHour, _ := strconv.Atoi(validFromSplit[1][:2])
					validFromMinute, _ := strconv.Atoi(validFromSplit[1][2:])

					// Valid to
					validToDay, _ := strconv.Atoi(validToSplit[0])
					validToHour, _ := strconv.Atoi(validToSplit[1][:2])
					validToMinute, _ := strconv.Atoi(validToSplit[1][2:])

					// Set validity
					validFrom = time.Date(time.Now().UTC().Year(), time.Month(j+1), validFromDay, validFromHour, validFromMinute, 0, 0, time.UTC)
					validTo = time.Date(time.Now().UTC().Year(), time.Month(j+1), validToDay, validToHour, validToMinute, 0, 0, time.UTC)

					// Terminate loop
					reached = true
				}

				// Terminate if operation complete for performance
				if reached {
					break
				}
			}
		}
	}

	// Final return list
	var finalTracks map[string]models.Track = make(map[string]models.Track)

	// Build track objects
	for _, track := range trackSlices {
		// Let's check validities now so we can save on performance
		timeNow := time.Now().UTC()
		tValidity := models.NA

		// Convert unix stamp to time
		validFrom = time.Unix(trackValidities[string(track[0][0])][0], 0)
		validTo = time.Unix(trackValidities[string(track[0][0])][1], 0)

		// validFrom is in the past but validTo is in the future, so track is valid now
		if timeNow.After(validFrom) && timeNow.Before(validTo) {
			tValidity = models.NOW
		}
		// validTo is in the past, therefore track was valid earlier
		if timeNow.After(validTo) {
			tValidity = models.EARLIER
		}
		// validFrom is in the future, therefore track is valid later
		if timeNow.Before(validFrom) {
			tValidity = models.LATER
		}

		// Check against the supplied validity
		if validity != models.NA {
			if tValidity != validity {
				continue
			}
		}

		// Initialise
		dir := models.UNKNOWN
		var flightLevels []int

		// Check direction and set flight levels
		var rawFlightLevels []string
		if strings.Contains(strings.ToUpper(track[1]), "EAST LVLS NIL") {
			// Set direction
			dir = models.WEST

			// Westbound flight levels
			rawFlightLevels = strings.Split(track[2], " ")[2:]

		} else {
			// Set direction
			dir = models.EAST

			// Westbound flight levels
			rawFlightLevels = strings.Split(track[1], " ")[2:]
		}

		// Convert to full altitudes
		for _, level := range rawFlightLevels {
			if isMetres { // Metres
				flInt, _ := strconv.Atoi(level)
				flightLevels = append(flightLevels, int(float64((flInt*100))*0.3048))
			} else { // Feet
				flInt, _ := strconv.Atoi(level)
				flightLevels = append(flightLevels, flInt*100)
			}
		}

		// Translate route strings into decimal coordinates
		route := strings.Split(track[0][2:], " ")
		var finalRoute []models.Fix
		for _, point := range route {
			// Check if the point is a coordinate
			if strings.Contains(point, "/") {
				// Lat/lon to use
				latlon, err := track_utils.ParseSlashedCoordinate(point)
				if err != nil {
					return nil, err
				}
				// Append fix
				finalRoute = append(finalRoute, models.CreateValidFix(point, latlon[0], latlon[1]))
			} else { // The point is a waypoint
				// Append waypoint
				finalRoute = append(finalRoute, fixes[point])
			}
		}

		// Finally, build the track object
		trackObj := models.Track{
			ID:           string(track[0][0]),
			TMI:          tmi,
			Route:        finalRoute,
			Direction:    dir,
			FlightLevels: flightLevels,
			ValidFrom:    trackValidities[string(track[0][0])][0],
			ValidTo:      trackValidities[string(track[0][0])][1],
		}

		// Finally, append the track but filter direction if argument supplied
		if (direction == models.EAST && dir == models.WEST) || (direction == models.WEST && dir == models.EAST) {
			continue
		} else {
			finalTracks[trackObj.ID] = trackObj
		}
	}

	// Return
	return finalTracks, nil
}
