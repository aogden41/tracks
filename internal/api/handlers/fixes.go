package handlers

import (
	"encoding/json"
	"github.com/aogden41/tracks/internal/db"
	"net/http"
)

// GET /fixes
func GetAllFixes(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")
	// Fetch fixes and check error
	fixes, err := db.SelectFixes()
	if err != nil {
		panic(err)
	}

	// Encode and return
	json.NewEncoder(w).Encode(fixes)
}

// GET /fixes/{fix_name}
func GetFix(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "This function has not yet been implemented.", http.StatusNotImplemented)
}

// POST /fixes/{fix_obj}
func PostFix(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "This function has not yet been implemented.", http.StatusNotImplemented)
}

// UPDATE /fixes/{fix_obj}
func UpdateFix(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "This function has not yet been implemented.", http.StatusNotImplemented)
}

// DELETE /fixes/{fix_name}
func DeleteFix(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "This function has not yet been implemented.", http.StatusNotImplemented)
}
