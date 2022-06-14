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
	CONCORDE
	EVENT
)

// Fix model
type Fix struct {
	Name      string
	Latitude  float64
	Longitude float64
}

// Track model
type Track struct {
	ID           string
	TMI          string
	Route        []Fix
	FlightLevels []int
	Direction    Direction
	ValidFrom    int64
	ValidTo      int64
	DaysOld      int
	Type         TrackType
}
