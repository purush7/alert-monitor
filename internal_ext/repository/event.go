package repository

import (
	"alert_monitor/internal_ext/models"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
)

var eventsRepo *EventsRepo
var eventsOnce sync.Once

type EventsRepo struct {
	db *sqlx.DB
}

func GetEventsRepo() *EventsRepo {
	eventsOnce.Do(func() {
		eventsRepo = &EventsRepo{
			db: db,
		}
	})
	return eventsRepo
}

// LogEvent logs an event in the database
func (e *EventsRepo) LogEvent(event models.Event) (models.Event, error) {
	query := `INSERT INTO events (client, event_type, created_at) VALUES ($1, $2, $3) RETURNING id`
	err := db.QueryRow(query, event.Client, event.EventType, event.CreatedAt).Scan(&event.ID)
	return event, err
}

// CountEvents returns the number of events for a given client and event type.
func (e *EventsRepo) CountEvents(client string, eventType string) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM events WHERE client = $1 AND event_type = $2`
	err := db.QueryRow(query, client, eventType).Scan(&count)
	return count, err
}

// CountEventsInRange returns the number of events within a specified time range.
func (e *EventsRepo) CountEventsInRange(client string, eventType string, startTime time.Time, endTime time.Time) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM events WHERE client = $1 AND event_type = $2 AND created_at BETWEEN $3 AND $4`
	err := db.QueryRow(query, client, eventType, startTime, endTime).Scan(&count)
	return count, err
}
