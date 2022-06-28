package models

import "strings"

// Direction enum
type Direction int32

const (
	UNKNOWN Direction = iota // Default direction
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
	IsValid   bool
}

func CreateValidFix(name string, lat float64, lon float64) Fix {
	return Fix{Name: strings.ToUpper(name), Latitude: lat, Longitude: lon, IsValid: true}
}

func CreateInvalidFix() Fix {
	return Fix{IsValid: false}
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
