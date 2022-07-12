package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aogden41/tracks/internal/db"
	"github.com/aogden41/tracks/internal/db/models"
	"github.com/gorilla/mux"
)

// GetAllFixes godoc
// @Summary Get all stored fixes
// @Description JSON output of all fixes stored in the API database
// @Tags fixes
// @Produce json
// @Success 200 {array} models.Fix
// @Failure 500 {object} InternalServerError
// @Router /fixes [get]
func GetAllFixes(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Fetch fixes and check error
	fixes, err := db.SelectFixes()
	if err != nil {
		json.NewEncoder(w).Encode(Error500(&w, err.Error()))
		return
	}

	// Encode
	json.NewEncoder(w).Encode(fixes)
}

// GetFix godoc
// @Summary Get one stored fixe
// @Description JSON output of one specific fixe stored in the API database
// @Tags fixes
// @Produce json
// @Param fix_name path string true "The name of the fix"
// @Success 200 {array} models.Fix
// @Failure 404 {object} NotFound
// @Failure 500 {object} InternalServerError
// @Router /fixes/{fix_name} [get]
func GetFix(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Get parameters, request data and encode
	params := mux.Vars(r)
	fix, err := db.SelectFix(params["fix_name"])
	if err != nil {
		if err == sql.ErrNoRows {
			json.NewEncoder(w).Encode(Error404(&w, "The requested fix was not found."))
		} else {
			json.NewEncoder(w).Encode(Error500(&w, err.Error()))
		}
		return
	}

	// Encode
	json.NewEncoder(w).Encode(fix)
}

// GetConcordeFixes godoc
// @Summary Get all stored Concorde fixes
// @Description JSON output of all Concorde fixes stored in the API database
// @Tags fixes
// @Produce json
// @Success 200 {array} models.Fix
// @Failure 500 {object} InternalServerError
// @Router /fixes/concorde [get]
func GetConcordeFixes(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Fetch fixes and check error
	fixes, err := db.SelectFixes()
	if err != nil {
		json.NewEncoder(w).Encode(Error500(&w, err.Error()))
		return
	}

	// Encode
	json.NewEncoder(w).Encode(fixes)
}

// TODO Errors, docs and auth for POST/PUT/DELETE
// POST /fixes
func PostFix(w http.ResponseWriter, r *http.Request) {
	// Set json header
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

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
	w.Header().Set("Access-Control-Allow-Origin", "*")

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
	w.Header().Set("Access-Control-Allow-Origin", "*")

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
