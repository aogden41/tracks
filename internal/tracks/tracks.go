package tracks

import (
	"io"
	"net/http"
	"strings"
	"unicode"
)

// Constants
const trackUrl = "https://www.notams.faa.gov/common/nat.html"
const fixesJson = "https://ganderoceanicoca.ams3.digitaloceanspaces.com/resources/data/fixes.json"

// Months of the year
var months = [12]string{"JAN", "FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEP", "OCT", "NOV", "DEC"}

// TODO: half waypoints
// Parse the track data
func ParseTracks(isMetres bool) {
	// Download the tracks and the fixes
	tracksRes, err := http.Get(trackUrl)
	fixesData, err := http.Get(fixesJson)

	// Returned string data from files
	var natData string
	var fixData string

	// Handle any error from web requests
	if err != nil {
		// do something...
	}

	// Fetching is done, defer closure until end of function
	defer tracksRes.Body.Close()
	defer fixesData.Body.Close()

	// Get bytes and assign string values
	bytes, _ := io.ReadAll(tracksRes.Body)
	natData = string(bytes) // Track message
	bytes, _ = io.ReadAll(fixesData.Body)
	fixData = string(bytes) // Downloaded fixes

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
	var validTo string
	var validFrom string

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

				}

				// Terminate if operation complete for performance
				if reached {
					break
				}
			}
		}
	}
}
