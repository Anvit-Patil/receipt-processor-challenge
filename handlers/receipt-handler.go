package handlers

import (
	"encoding/json"
	"net/http"

	"receipt-processor/models"
	"receipt-processor/storage"
	"receipt-processor/utils"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Handle POST /receipts/process
func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt models.Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Generate a unique ID
	id := uuid.New().String()

	// Calculate points
	points := utils.CalculatePoints(receipt)

	// Save receipt and points in memory
	storage.SaveReceipt(id, points)

	// Respond with generated ID
	response := map[string]string{"id": id}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Handle GET /receipts/{id}/points
func GetPoints(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Retrieve points
	points, exists := storage.GetPoints(id)
	if !exists {
		http.Error(w, "No receipt found for that ID", http.StatusNotFound)
		return
	}

	// Respond with points
	response := map[string]int{"points": points}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
