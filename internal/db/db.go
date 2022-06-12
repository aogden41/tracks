package db

import (
	"database/sql"
	"fmt"
	"os"
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
	cxt := Connect()
	defer cxt.Close()

}

func SelectConcordeTracks() {
	// Connect and defer
	cxt := Connect()
	defer cxt.Close()

}

func SelectEventTracks() {
	// Connect and defer
	cxt := Connect()
	defer cxt.Close()

}

func SelectAllFixes() {
	// Connect and defer
	cxt := Connect()
	defer cxt.Close()

}

func SelectAllCachedTracks() {
	// Connect and defer
	cxt := Connect()
	defer cxt.Close()

}

func DeleteAllCachedTracks() {
	// Connect and defer
	cxt := Connect()
	defer cxt.Close()

}

func InsertFix() {
	// Connect and defer
	cxt := Connect()
	defer cxt.Close()

}

func InsertTrack() {
	// Connect and defer
	cxt := Connect()
	defer cxt.Close()

}

func InsertMultipleTracks() {
	// Connect and defer
	cxt := Connect()
	defer cxt.Close()

}

func UpdateFix() {
	// Connect and defer
	cxt := Connect()
	defer cxt.Close()

}

func UpdateTrack() {
	// Connect and defer
	cxt := Connect()
	defer cxt.Close()

}

func DeleteFix() {
	// Connect and defer
	cxt := Connect()
	defer cxt.Close()

}

func DeleteCachedTracksOneDay() {
	// Connect and defer
	cxt := Connect()
	defer cxt.Close()

}

func DeleteTrack() {
	// Connect and defer
	cxt := Connect()
	defer cxt.Close()

}
