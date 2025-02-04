package main

import (
	"fetch-assessment/handlers"
	"fetch-assessment/store"
	mux2 "github.com/gorilla/mux"
	"log"
	"net/http"
)

const serverPort = ":8080"

func main() {

	mux := mux2.NewRouter()

	const serverURL = "localhost" + serverPort

	receiptStore := store.NewReceiptStore()
	mux.HandleFunc("/receipts/process", func(w http.ResponseWriter, r *http.Request) {
		handlers.StoreReceiptHandler(w, r, receiptStore)
	}).Methods("POST")
	mux.HandleFunc("/receipts/{id}/points", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetPointsHandler(w, r, receiptStore)
	}).Methods("GET")

	err := http.ListenAndServe(serverPort, mux)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Server started on %s", serverURL)
}
