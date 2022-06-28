package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aogden41/tracks/internal/db"
	"github.com/aogden41/tracks/internal/db/models"
	"github.com/gorilla/mux"
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

	// Encode
	json.NewEncoder(w).Encode(fixes)
}

// GET /fixes/{fix_name}
func GetFix(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")

	// Get parameters, request data and encode
	params := mux.Vars(r)
	fix, err := db.SelectFix(params["fix_name"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		json.NewEncoder(w).Encode(fix)
	}
}

// GET /fixes
func GetConcordeFixes(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")

	// Fetch fixes and check error
	fixes, err := db.SelectConcordeFixes()
	if err != nil {
		panic(err)
	}

	// Encode
	json.NewEncoder(w).Encode(fixes)
}

// POST /fixes
func PostFix(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")

	// Read body and error check
	var fix models.Fix
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Read form
	name := r.Form.Get("Name")
	lat, _ := strconv.ParseFloat(r.Form.Get("Latitude"), 64)
	lon, _ := strconv.ParseFloat(r.Form.Get("Longitude"), 64)
	if name != "" && lat != 0 && lon != 0 {
		fix = models.CreateValidFix(r.Form.Get("Name"), lat, lon)
	} else {
		http.Error(w, "Invalid or empty query parameters.", http.StatusBadRequest)
		return
	}

	// Decoded fine, so insert and return
	err = db.InsertFix(&fix)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(""))
	}
}

// PUT /fixes
func PutFix(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")

	// Read body and error check
	var fix models.Fix
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Read form
	name := r.Form.Get("Name")
	lat, _ := strconv.ParseFloat(r.Form.Get("Latitude"), 64)
	lon, _ := strconv.ParseFloat(r.Form.Get("Longitude"), 64)
	if name != "" && lat != 0 && lon != 0 {
		fix = models.CreateValidFix(r.Form.Get("Name"), lat, lon)
	} else {
		http.Error(w, "Invalid or empty query parameters.", http.StatusBadRequest)
		return
	}

	// Update and return
	err = db.UpdateFix(&fix)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(""))
	}
}

// DELETE /fixes/{fix_name}
func DeleteFix(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")

	// Get parameters, delete and return
	params := mux.Vars(r)
	err := db.DeleteFix(params["fix_name"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(""))
	}
}
