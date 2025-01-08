package main

import (
	"expense-tracker/db"
	"expense-tracker/handlers"
	"expense-tracker/migrations"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db.InitDB()
	defer db.DB.Close()

	db.InitRedis()

	migrations.RunMigrations(db.DB)

	r := mux.NewRouter()

	r.HandleFunc("/sources", handlers.GetSource).Methods("GET")
	r.HandleFunc("/sources/create", handlers.CreateSource).Methods("POST")
	r.HandleFunc("/sources/delete/{id}", handlers.DeleteSource).Methods("DELETE")

	r.HandleFunc("/transactions", handlers.GetTransaction).Methods("GET")
	r.HandleFunc("/transactions/create", handlers.CreateTransaction).Methods("POST")
	r.HandleFunc("/transactions/delete/{id}", handlers.DeleteTransaction).Methods("DELETE")

	port := ":8080"

	fmt.Printf("Server running on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
