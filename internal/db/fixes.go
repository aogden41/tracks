package db

import (
	"database/sql"
	"github.com/aogden41/tracks/internal/db/models"
	"strings"

	_ "github.com/lib/pq"
)

// Return a map of fixes
func SelectFixes() (map[string]models.Fix, error) {
	// Connect and defer
	db := Connect()
	defer db.Close()

	// Statement
	query := `SELECT name, latitude, longitude FROM tracks.fixes;`

	// Perform query
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Map to return
	fixes := make(map[string]models.Fix)

	// Iterate through rows
	for rows.Next() {
		// Create fix and error check
		var fix models.Fix
		if err := rows.Scan(&fix.Name, &fix.Latitude, &fix.Longitude); err != nil {
			return nil, err
		}
		// Add to the map
		fix.IsValid = true
		fixes[fix.Name] = fix
	}

	// Catch any other error
	if err = rows.Err(); err != nil {
		return nil, err
	}

	// Success, return everything
	return fixes, nil
}

// Return a single fix
func SelectFix(f string) (models.Fix, error) {
	// Connect and defer
	db := Connect()
	defer db.Close()

	// Statement
	query := `SELECT name, latitude, longitude FROM tracks.fixes WHERE name LIKE $1;`

	// Get the row and error check
	var fix models.Fix
	row := db.QueryRow(query, strings.ToUpper(f))
	if err := row.Scan(&fix.Name, &fix.Latitude, &fix.Longitude); err != nil {
		return models.CreateInvalidFix(), err
	}

	// Return selection
	fix.IsValid = true
	return fix, nil
}

func SelectConcordeFixes() (map[string]models.Fix, error) {
	// Connect and defer
	db := Connect()
	defer db.Close()

	// Statement
	query := `SELECT name, latitude, longitude FROM tracks.fixes WHERE CAST(type AS TEXT) LIKE '1';`

	// Perform query
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Map to return
	fixes := make(map[string]models.Fix)

	// Iterate through rows
	for rows.Next() {
		// Create fix and error check
		var fix models.Fix
		if err := rows.Scan(&fix.Name, &fix.Latitude, &fix.Longitude); err != nil {
			return nil, err
		}
		if len(fix.Name) > 5 {
			fix.Name = fix.Name[:5]
		}
		// Add to the map
		fix.IsValid = true
		fixes[fix.Name] = fix
	}

	// Catch any other error
	if err = rows.Err(); err != nil {
		return nil, err
	}

	// Success, return everything
	return fixes, nil
}

// Insert a fix
func InsertFix(f *models.Fix) error {
	// Connect and defer
	db := Connect()
	defer db.Close()

	// Statement
	query := `INSERT INTO tracks.fixes (name, latitude, longitude) VALUES ($1, $2, $3) 
				ON CONFLICT ON CONSTRAINT name_idx DO NOTHING RETURNING id;`

	// Complete query, check error and return
	rowId := 0 // Need this for auto increment but we discard it straight after
	err := db.QueryRow(query, f.Name, f.Latitude, f.Longitude).Scan(&rowId)
	switch err {
	case sql.ErrNoRows: // Duplicate entry, send status 200
		return nil
	case nil: // No error
		return nil
	default:
		return err
	}
}

// Update a fix
func UpdateFix(f *models.Fix) error {
	// Connect and defer
	db := Connect()
	defer db.Close()

	// Statement
	query := `UPDATE tracks.fixes SET latitude = $2, longitude = $3 WHERE name LIKE $1;`

	// Complete query, check error and return
	_, err := db.Exec(query, f.Name, f.Latitude, f.Longitude)
	if err != nil {
		return err
	}
	return nil
}

// Delete a fix
func DeleteFix(f string) error {
	// Connect and defer
	db := Connect()
	defer db.Close()

	// Statement
	query := `DELETE FROM tracks.fixes WHERE name LIKE $1;`

	// Complete query, check error and return
	_, err := db.Exec(query, strings.ToUpper(f))
	if err != nil {
		return err
	}
	return nil
}
