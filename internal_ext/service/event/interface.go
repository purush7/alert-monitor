package event

import (
	"alert_monitor/internal_ext/models"
	"time"
)

type UseCase interface {
	Event(models.Event) (models.Event, error)
}

type repo interface {
	LogEvent(models.Event) (models.Event, error)
	CountEvents(client string, eventType string) (int, error)
	CountEventsInRange(client string, eventType string, startTime time.Time, endTime time.Time) (int, error)
}

type dispatcher interface {
	DispatchAlerts(alertId int, dispatchModes []models.DispatchStrategy, client, eventType string)
}
