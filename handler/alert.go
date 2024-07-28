package handler

import (
	"alert_monitor/internal_ext/models"
	"encoding/json"
	"net/http"
	"strconv"
)

// AlertHandler processes different HTTP methods for the /alerts endpoint
func AlertHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		CreateAlertHandler(w, r)
	case http.MethodGet:
		GetAlertHandler(w, r)
	case http.MethodPut:
		UpdateAlertHandler(w, r)
	case http.MethodDelete:
		DeleteAlertHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// CreateAlertHandler handles POST requests to create an alert
func CreateAlertHandler(w http.ResponseWriter, r *http.Request) {
	var alert models.Alert
	if err := json.NewDecoder(r.Body).Decode(&alert); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var err error
	if alert, err = alertService.CreateAlert(alert); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(alert)

}

// GetAlertHandler handles GET requests to retrieve an alert
func GetAlertHandler(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id := -1
	var err error
	if idString != "" {
		id, err = strconv.Atoi(idString)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
	}

	alerts, err := alertService.GetAlert(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(alerts)
}

// UpdateAlertHandler handles PUT requests to update an alert
func UpdateAlertHandler(w http.ResponseWriter, r *http.Request) {
	var alert models.Alert
	if err := json.NewDecoder(r.Body).Decode(&alert); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := alertService.UpdateAlert(alert); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(alert)
}

// DeleteAlertHandler handles DELETE requests to delete an alert
func DeleteAlertHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := alertService.DeleteAlert(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
