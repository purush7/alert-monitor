package models

import (
	"time"
)

// AlertConfigType represents different types of alert configurations
type AlertConfigType string
type AlertType string

const (
	SIMPLE_COUNT    AlertConfigType = "SIMPLE_COUNT"
	TUMBLING_WINDOW AlertConfigType = "TUMBLING_WINDOW"
	SLIDING_WINDOW  AlertConfigType = "SLIDING_WINDOW"
	CONSOLE         AlertType       = "CONSOLE"
	EMAIL           AlertType       = "EMAIL"
)

// AlertConfig holds the configuration for an alert
type AlertConfig struct {
	Type             AlertConfigType `json:"type"`
	Count            int             `json:"count"`
	WindowSizeInSecs int             `json:"windowSizeInSecs"`
}

// DispatchStrategy holds the dispatch configuration for alerts
type DispatchStrategy struct {
	Type    AlertType `json:"type"`
	Message string    `json:"message,omitempty"`
	Subject string    `json:"subject,omitempty"`
}

// Alert represents the alert configuration in the database
type Alert struct {
	ID               int                `json:"id"`
	Client           string             `json:"client"`
	EventType        string             `json:"eventType"`
	AlertConfig      AlertConfig        `json:"alertConfig"`
	DispatcherConfig []DispatchStrategy `json:"dispatcherConfig"` // JSON config as a string
}

// Event represents an event logged in the database
type Event struct {
	ID        int       `json:"id"`
	Client    string    `json:"client"`
	EventType string    `json:"eventType"`
	CreatedAt time.Time `json:"createdAt"`
}
