package api

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