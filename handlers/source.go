package handlers

import (
	"encoding/json"
	"expense-tracker/db"
	"expense-tracker/models"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func CreateSource(w http.ResponseWriter, r *http.Request) {
	var source models.Source

	if err := json.NewDecoder(r.Body).Decode(&source); err != nil {
		http.Error(w, "Invalid request payload", http.StatusInternalServerError)
		return
	}

	_, err := db.DB.NewInsert().Model(&source).Exec(r.Context())

	if err != nil {
		http.Error(w, "Failed to create source:", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(source)
}

func GetSource(w http.ResponseWriter, r *http.Request) {
	cacheKey := "sources_cache"

	if cachedSources, err := db.RedisClient.Get(db.Ctx, cacheKey).Result(); err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(cachedSources))
		return
	}

	var sources []models.Source

	if err := db.DB.NewSelect().Model(&sources).Scan(r.Context()); err != nil {
		http.Error(w, "Error fetching sources:", http.StatusInternalServerError)
		return
	}

	sourcesJSON, _ := json.Marshal(sources)

	if err := db.RedisClient.Set(db.Ctx, cacheKey, sourcesJSON, 5*time.Minute).Err(); err != nil {
		log.Printf("Failed to cache sources: %v\n", err)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sources)
}

func DeleteSource(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, exists := vars["id"]

	if !exists {
		http.Error(w, "ID not provided", http.StatusBadRequest)
		return
	}

	result, err := db.DB.NewDelete().Model((*models.Source)(nil)).Where("id = ?", id).Exec(r.Context())

	if err != nil {
		http.Error(w, "Failed to delete source.", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		http.Error(w, "Source ID not found/incorrect", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Source Deleted Successfully!")
}
