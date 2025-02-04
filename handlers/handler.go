package handlers

import (
	"encoding/json"
	"fetch-assessment/calculator"
	"fetch-assessment/model"
	"fetch-assessment/store"
	"fetch-assessment/validation"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type totalResponse struct {
	Points int `json:"points"`
}

type uuidResponse struct {
	Id string `json:"id"`
}

func StoreReceiptHandler(w http.ResponseWriter, r *http.Request, receiptStore *store.ReceiptStore) {

	decoder := json.NewDecoder(r.Body)
	var rc model.Receipt
	err := decoder.Decode(&rc)
	if err != nil {
		writeErrorResponse(w, err)
		return
	}

	valid, err := validation.ValidateReceipt(rc)
	if err != nil || !valid {
		writeErrorResponse(w, err)
		return
	}

	item, err := receiptStore.Store(rc)
	if err != nil {
		writeErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	id := item.String()
	response := uuidResponse{Id: id}
	json.NewEncoder(w).Encode(response)
}

func writeErrorResponse(w http.ResponseWriter, err error) {
	fmt.Println("Invalid request format", err)
	http.Error(w, "Invalid request format", http.StatusBadRequest)
}

func GetPointsHandler(w http.ResponseWriter, r *http.Request, receiptStore *store.ReceiptStore) {

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		fmt.Println("id is missing in parameters")
		http.Error(w, "Receipt ID missing", http.StatusNotFound)
		return
	}
	receipt := receiptStore.GetReceipt(id)
	if receipt == nil {
		fmt.Printf("receipt for id %s not present\n", id)
		http.Error(w, "Receipt ID not found", http.StatusNotFound)
		return
	}
	total := calculator.CalculateTotals(*receipt)
	w.Header().Set("Content-Type", "application/json")
	var totalRsp = totalResponse{
		Points: total,
	}
	json.NewEncoder(w).Encode(totalRsp)

}
