package dispatch

import "alert_monitor/internal_ext/models"

type UseCase interface {
	DispatchAlerts(alertId int, dispatchModes []models.DispatchStrategy, client, eventType string)
}
