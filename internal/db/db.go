package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// Connection config
type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

// Create Postgres connection
func Connect() (ctx *sql.DB) {
	// Connection config
	var dbc DbConfig = DbConfig{os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("USER"), os.Getenv("PASS"), os.Getenv("DB_NAME")}

	// Create connection string
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbc.Host, dbc.Port, dbc.User, dbc.Password, dbc.Database)

	// Open connection
	db, err := sql.Open("postgres", conn)

	// Error check
	if err != nil {
		panic(err)
	}

	// Check if DB can be reached
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// Log & return
	fmt.Println("Postgres connection succeeded.")
	return db
}

func SelectCurrentTracks() {
	// Connect and defer
	ctx := Connect()
	defer ctx.Close()

}

func SelectConcordeTracks() {
	// Connect and defer
	ctx := Connect()
	defer ctx.Close()

}

func SelectEventTracks() {
	// Connect and defer
	ctx := Connect()
	defer ctx.Close()

}

func SelectFixes() (fixesSlice *[]Fix, err error) {
	// Connect and defer
	ctx := Connect()
	defer ctx.Close()

	// Statement
	query := `SELECT * from tracks.fixes;`

	// Perform query
	rows, err := ctx.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Return slice
	var fixes []Fix

	// Iterate through rows
	for rows.Next() {
		// Create fix and error check
		var fix Fix
		if err := rows.Scan(&fix.ID, &fix.Name, &fix.Latitude, &fix.Longitude); err != nil {
			return &fixes, err
		}
		// Append
		fixes = append(fixes, fix)
	}

	// Catch any other error
	if err = rows.Err(); err != nil {
		return &fixes, err
	}

	// Success, return everything
	return &fixes, nil
}

func SelectAllCachedTracks() {
	// Connect and defer
	ctx := Connect()
	defer ctx.Close()

}

func DeleteAllCachedTracks() {
	// Connect and defer
	ctx := Connect()
	defer ctx.Close()

}

func InsertFix() {
	// Connect and defer
	ctx := Connect()
	defer ctx.Close()

}

func InsertTrack() {
	// Connect and defer
	ctx := Connect()
	defer ctx.Close()

}

func InsertMultipleTracks() {
	// Connect and defer
	ctx := Connect()
	defer ctx.Close()

}

func UpdateFix() {
	// Connect and defer
	ctx := Connect()
	defer ctx.Close()

}

func UpdateTrack() {
	// Connect and defer
	ctx := Connect()
	defer ctx.Close()

}

func DeleteFix() {
	// Connect and defer
	ctx := Connect()
	defer ctx.Close()

}

func DeleteCachedTracksOneDay() {
	// Connect and defer
	ctx := Connect()
	defer ctx.Close()

}

func DeleteTrack() {
	// Connect and defer
	ctx := Connect()
	defer ctx.Close()

}
