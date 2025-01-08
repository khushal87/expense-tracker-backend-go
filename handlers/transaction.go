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

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction models.Transaction

	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		http.Error(w, "Invalid request payload:"+err.Error(), http.StatusInternalServerError)
		return
	}

	_, err := db.DB.NewInsert().Model(&transaction).Exec(r.Context())

	if err != nil {
		http.Error(w, "Failed to create transaction:"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(transaction)
}

func GetTransaction(w http.ResponseWriter, r *http.Request) {
	cacheKey := "transaction_cache"

	if cachedTransactions, err := db.RedisClient.Get(db.Ctx, cacheKey).Result(); err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(cachedTransactions))
		return
	}

	var transactions []models.Transaction

	if err := db.DB.NewSelect().Model(&transactions).Scan(r.Context()); err != nil {
		http.Error(w, "Error fetching transactions:", http.StatusInternalServerError)
		return
	}

	transactionsJSON, _ := json.Marshal(transactions)
	if err := db.RedisClient.Set(db.Ctx, cacheKey, transactionsJSON, 5*time.Minute); err != nil {
		log.Printf("Failed to cache transactions: %v\n", err)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transactions)
}

func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, exists := vars["id"]

	if !exists {
		http.Error(w, "ID not provided", http.StatusBadRequest)
		return
	}

	result, err := db.DB.NewDelete().Model((*models.Transaction)(nil)).Where("id = ?", id).Exec(r.Context())

	if err != nil {
		http.Error(w, "Failed to delete transaction.", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		http.Error(w, "Transaction ID not found/incorrect", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Transaction Deleted Successfully!")
}
