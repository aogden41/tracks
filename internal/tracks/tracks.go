package tracks

import (
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/aogden41/tracks/internal/db"
)

// Constants
const trackUrl = "https://www.notams.faa.gov/common/nat.html"

// Months of the year
var months = [12]string{"JAN", "FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEP", "OCT", "NOV", "DEC"}

// Parse the track data
func ParseTracks(isMetres bool) {
	// Download the tracks
	tracksRes, err := http.Get(trackUrl)

	// Returned string data from files
	var natData string

	// Handle any error from web requests
	if err != nil {
		// do something...
	}

	// Fetching is done, defer closure until end of function
	defer tracksRes.Body.Close()

	// Get bytes and assign string values
	bytes, _ := io.ReadAll(tracksRes.Body)
	natData = string(bytes) // Track message

	// Split the data
	natDataLines := strings.Split(string(natData), "\n")

	// Keep track of line numbers to get correct track validity
	var lineNumbers []int

	// Store extracted track slices
	var trackSlices [][]string

	// Store data currently being processed
	var track []string
	var validities []string
	var tmi string
	var validFrom time.Time
	var validTo time.Time

	// Flag to get the next 2 lines for a track
	getNextTwo := 0

	// Iterate through
	for i := 0; i < len(natDataLines); i++ {
		// Rune slice of current line
		lineRunes := []rune(natDataLines[i])

		// If the row is a track (we assume this if the structure is a letter then a space then another 2 letters, or if a track is already being processed)
		if (unicode.IsLetter(lineRunes[0]) && unicode.IsSpace(lineRunes[1]) && unicode.IsLetter(lineRunes[2]) && unicode.IsLetter(lineRunes[3])) || getNextTwo > 0 {
			// If greater than 2 then we've collected a track, so process and move on
			if getNextTwo > 2 {
				getNextTwo = 0
				trackSlices = append(trackSlices, track)
				track = []string{}
				lineNumbers = append(lineNumbers, i-2)
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
				tmi = tmi + string(natDataLines[i][tmiStart+4])
			}
		} else {
			// Convert validities TODO
			for j := 0; j < len(months); j++ {
				reached := false

				// Check if the line contains a month
				if strings.Contains(natDataLines[i], months[j]) {
					// Split the to and from times
					splitString := strings.Split(natDataLines[i], " TO ")
					validFromSplit := strings.Split(splitString[0][4:], "/")
					validToSplit := strings.Split(splitString[1][4:], "/")

					// Valid from
					validFromDay, _ := strconv.Atoi(validFromSplit[0])
					validFromHour, _ := strconv.Atoi(validFromSplit[1][0:2])
					validFromMinute, _ := strconv.Atoi(validFromSplit[1][2:])

					// Valid to
					validToDay, _ := strconv.Atoi(validToSplit[0])
					validToHour, _ := strconv.Atoi(validToSplit[1][0:2])
					validToMinute, _ := strconv.Atoi(validToSplit[1][2:])

					// Set validity
					validFrom = time.Date(time.Now().UTC().Year(), time.Month(j+1), validFromDay, validFromHour, validFromMinute, 0, 0, time.UTC)
					validTo = time.Date(time.Now().UTC().Year(), time.Month(j+1), validToDay, validToHour, validToMinute, 0, 0, time.UTC)

					// Append
					validities = append(validities, strconv.Itoa(int(validFrom.Unix()))+"?"+strconv.Itoa(int(validTo.Unix()))+"?"+strconv.Itoa(i))

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

	// TODO: half waypoints
	// Final return list
	var finalTracks []db.Track

	// Build track objects
	for _, track := range trackSlices {
		// Initialise
		dir := db.UNKNOWN
		var flightLevels []int

		// Check direction and set flight levels
		var rawFlightLevels []string
		if strings.Contains(strings.ToUpper(track[1]), "EAST LVLS NIL") {
			// Set direction
			dir = db.WEST

			// Westbound flight levels
			rawFlightLevels = strings.Split(track[2], " ")[2:]

		} else {
			// Set direction
			dir = db.EAST

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

	}
}
