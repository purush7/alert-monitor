package repository

import (
	"alert_monitor/internal_ext/models"
	"encoding/json"
	"sync"

	"github.com/jmoiron/sqlx"
)

var alertsRepo *AlertsRepo
var alertsOnce sync.Once

type AlertsRepo struct {
	db *sqlx.DB
}

func GetAlertsRepo() *AlertsRepo {
	alertsOnce.Do(func() {
		alertsRepo = &AlertsRepo{
			db: db,
		}
	})
	return alertsRepo
}

// CreateAlert inserts a new alert into the database
func (a *AlertsRepo) CreateAlert(alert models.Alert) (models.Alert, error) {
	query := `INSERT INTO alert (client, event_type, alert_config_type, count, window_size_in_secs, dispatcher_config) 
              VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	dispatcherConfig, err := json.Marshal(alert.DispatcherConfig)
	if err != nil {
		return alert, err
	}

	err = db.QueryRow(query, alert.Client, alert.EventType, alert.AlertConfig.Type, alert.AlertConfig.Count, alert.AlertConfig.WindowSizeInSecs, string(dispatcherConfig)).Scan(&alert.ID)
	return alert, err
}

// GetAlert retrieves an alert by ID
func (a *AlertsRepo) GetAlert(id int) ([]models.Alert, error) {
	query := `SELECT id, client, event_type, alert_config_type, count, window_size_in_secs, dispatcher_config FROM alert `
	args := []any{}
	if id != -1 {
		query += ` WHERE id = $1`
		args = append(args, id)
	}
	alerts := make([]models.Alert, 0)
	rows, err := db.Query(query, args...)
	if err != nil {
		return alerts, err
	}
	for rows.Next() {
		var alert models.Alert
		var dispactherConfig json.RawMessage
		err := rows.Scan(&alert.ID, &alert.Client, &alert.EventType, &alert.AlertConfig.Type, &alert.AlertConfig.Count, &alert.AlertConfig.WindowSizeInSecs, &dispactherConfig)
		if err != nil {
			return alerts, err
		}
		err = json.Unmarshal(dispactherConfig, &alert.DispatcherConfig)
		if err != nil {
			return alerts, err
		}
		alerts = append(alerts, alert)
	}

	return alerts, nil
}

// UpdateAlert updates an existing alert
func (a *AlertsRepo) UpdateAlert(alert models.Alert) error {
	query := `UPDATE alert SET client = $1, event_type = $2, alert_config_type = $3, count = $4, window_size_in_secs = $5, dispatcher_config = $6 WHERE id = $7`
	dispatcherConfig, err := json.Marshal(alert.DispatcherConfig)
	if err != nil {
		return err
	}
	_, err = db.Exec(query, alert.Client, alert.EventType, alert.AlertConfig.Type, alert.AlertConfig.Count, alert.AlertConfig.WindowSizeInSecs, string(dispatcherConfig), alert.ID)
	return err
}

// DeleteAlert removes an alert by ID
func (a *AlertsRepo) DeleteAlert(id int) error {
	query := `DELETE FROM alert WHERE id = $1`
	_, err := db.Exec(query, id)
	return err
}

func (a *AlertsRepo) GetAllAlerts(client, eventType string) ([]models.Alert, error) {
	query := `SELECT id, client, event_type, alert_config_type, count, window_size_in_secs, dispatcher_config FROM alert where client = $1 and event_type = $2`
	alerts := make([]models.Alert, 0)
	rows, err := db.Query(query, client, eventType)
	if err != nil {
		return alerts, err
	}
	for rows.Next() {
		var alert models.Alert
		var dispactherConfig json.RawMessage
		err := rows.Scan(&alert.ID, &alert.Client, &alert.EventType, &alert.AlertConfig.Type, &alert.AlertConfig.Count, &alert.AlertConfig.WindowSizeInSecs, &dispactherConfig)
		if err != nil {
			return alerts, err
		}
		err = json.Unmarshal(dispactherConfig, &alert.DispatcherConfig)
		if err != nil {
			return alerts, err
		}
		alerts = append(alerts, alert)
	}

	return alerts, nil
}
