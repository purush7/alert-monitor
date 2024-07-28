package handler

import (
	"alert_monitor/internal_ext/models"
	"encoding/json"
	"net/http"
	"time"
)

// LogEventHandler handles POST requests to log an event
func EventHandler(w http.ResponseWriter, r *http.Request) {
	var event models.Event
	var err error
	if err = json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	event.CreatedAt = time.Now()

	if event, err = eventService.Event(event); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(event)
}
