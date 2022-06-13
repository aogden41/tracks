package db

// Direction enum
type Direction int32

const (
	UNKNOWN Direction = iota
	WEST
	EAST
)

// Track type enum
type TrackType int32

const (
	REGULAR TrackType = iota
	EVENT
	CONCORDE
)

// Fix model
type Fix struct {
	ID        int     `json:id`
	Name      string  `json:name`
	Latitude  float64 `json:latitude`
	Longitude float64 `json:longitude`
}

// Track model
type Track struct {
	ID           int       `json:uid`
	TrackID      string    `json:id`
	TMI          string    `json:tmi`
	Route        []Fix     `json:route`
	FlightLevels []int32   `json:flightLevels`
	Direction    Direction `json:direction`
	ValidFrom    string    `json:validFrom`
	ValidTo      string    `json:validTo`
	DaysOld      int       `json:daysOld`
	Type         TrackType `json:type`
}
