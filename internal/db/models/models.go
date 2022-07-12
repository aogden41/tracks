package models

import "strings"

// Direction enum
type Direction int32

const (
	UNKNOWN Direction = iota // Default direction
	WEST
	EAST
)

// Currency enum
type Validity int32

const (
	NA Validity = iota // Default direction
	NOW
	EARLIER
	LATER
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
	Name      string  `json:"Name"`
	Latitude  float64 `json:"Latitude"`
	Longitude float64 `json:"Longitude"`
	IsValid   bool    `json:"IsValid"`
}

func CreateValidFix(name string, lat float64, lon float64) Fix {
	return Fix{Name: strings.ToUpper(name), Latitude: lat, Longitude: lon, IsValid: true}
}

func CreateInvalidFix() Fix {
	return Fix{IsValid: false}
}

// Track model
type Track struct {
	ID           string    `json:"ID"`
	TMI          string    `json:"TMI"`
	Route        []Fix     `json:"Route"`
	FlightLevels []int     `json:"FlightLevels"`
	Direction    Direction `json:"Direction"`
	ValidFrom    int64     `json:"ValidFrom"`
	ValidTo      int64     `json:"ValidTo"`
	Type         TrackType `json:"Type"`
}
