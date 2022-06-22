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
func Connect() *sql.DB {
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
