package alert

import "alert_monitor/internal_ext/models"

type UseCase interface {
	CreateAlert(alert models.Alert) (models.Alert, error)
	GetAlert(id int) ([]models.Alert, error)
	UpdateAlert(alert models.Alert) error
	DeleteAlert(id int) error
	GetAllAlerts(client, eventType string) ([]models.Alert, error)
}

type repo interface {
	CreateAlert(alert models.Alert) (models.Alert, error)
	GetAlert(id int) ([]models.Alert, error)
	UpdateAlert(alert models.Alert) error
	DeleteAlert(id int) error
	GetAllAlerts(client, eventType string) ([]models.Alert, error)
}
