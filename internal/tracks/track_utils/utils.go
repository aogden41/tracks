package track_utils

import (
	"math"
	"strconv"
	"strings"
)

// Parse a coordinate in the format 'XX/XX'
func ParseSlashedCoordinate(point string) ([2]float64, error) {
	// Lat/lon to return
	var latitude float64
	var longitude float64

	// Split the point
	pointSplit := strings.Split(point, "/")

	// Parse latitude
	if len(pointSplit[0]) > 2 {
		// Significant digits (first 2 digits)
		s := pointSplit[0][:2]
		// Decimal places (everything else)
		var d string

		// Since lat/lon minutes are out of 60, parse them to decimal (i.e. 30 becomes 50))
		d1, _ := strconv.ParseInt(string(pointSplit[0][2:][0]), 10, 8)
		d2, _ := strconv.ParseInt(string(pointSplit[0][2:][1]), 10, 8)
		d1 = int64(math.RoundToEven(10.0 * (float64(d1) / 6.0))) // Eg. 10 * (3/6) is 0.5, because 3/6 is 5
		d = strconv.Itoa(int(d1)) + strconv.Itoa(int(d2))

		// Put back together
		var err error
		latitude, err = strconv.ParseFloat(strings.Join([]string{s, d}, "."), 64)
		if err != nil { // Error check
			return [2]float64{0, 0}, err
		}
	} else {
		// Latitude isn't fractional, return float for consistency
		var err error
		latitude, err = strconv.ParseFloat(pointSplit[0], 64)
		if err != nil { // Error check
			return [2]float64{0, 0}, err
		}
	}

	// Parse longitude
	if len(pointSplit[1]) > 2 {
		// Significant digits (first 2 digits)
		s := pointSplit[1][:2]
		// Decimal places (everything else)
		var d string

		// Since lat/lon minutes are out of 60, parse them to decimal (i.e. 30 becomes 50))
		d1, _ := strconv.ParseInt(string(pointSplit[1][2:][0]), 10, 8)
		d2, _ := strconv.ParseInt(string(pointSplit[1][2:][1]), 10, 8)
		d1 = int64(math.RoundToEven(10.0 * (float64(d1) / 6.0))) // Eg. 10 * (3/6) is 0.5, because 3/6 is 5
		d = strconv.Itoa(int(d1)) + strconv.Itoa(int(d2))

		// Put back together
		var err error
		longitude, err = strconv.ParseFloat(strings.Join([]string{s, d}, "."), 64)
		if err != nil { // Error check
			return [2]float64{0, 0}, err
		}
	} else {
		// Longitude isn't fractional, return float for consistency
		var err error
		longitude, err = strconv.ParseFloat(pointSplit[1], 64)
		if err != nil { // Error check
			return [2]float64{0, 0}, err
		}
	}

	// Return lat/lon pair
	return [2]float64{latitude, -longitude}, nil
}
