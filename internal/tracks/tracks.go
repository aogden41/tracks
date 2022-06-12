package tracks

// Constants
const trackUrl = "https://www.notams.faa.gov/common/nat.html"
const fixesJson = "https://ganderoceanicoca.ams3.digitaloceanspaces.com/resources/data/fixes.json"

// Months of the year
var months = [12]string{"JAN", "FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEP", "OCT", "NOV", "DEC"}

// Direction enum
type Direction int32

const (
	UNKNOWN Direction = iota
	WEST
	EAST
)

// Fix model
type Fix struct {
	Name      string  `json:name`
	Latitude  float64 `json:latitude`
	Longitude float64 `json:longitude`
}

// Track model
type Track struct {
	ID           string    `json:id`
	TMI          string    `json:tmi`
	Route        []Fix     `json:route`
	FlightLevels []int32   `json:flightLevels`
	Direction    Direction `json:direction`
	ValidFrom    string    `json:validFrom`
	ValidTo      string    `json:validTo`
}

// Parse the track data
func ParseTracks(isMetres bool) []Track {
	// Download the tracks
}
