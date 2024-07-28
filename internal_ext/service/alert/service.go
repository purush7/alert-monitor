package alert

import (
	"alert_monitor/internal_ext/models"
	"alert_monitor/internal_ext/repository"
	"sync"
)

var alertService *Service = nil
var serviceOnce sync.Once

func GetService() UseCase {
	serviceOnce.Do(func() {
		alertService = &Service{
			repo: repository.GetAlertsRepo(),
		}
	})

	return alertService
}

type Service struct {
	repo repo
}

func (s *Service) CreateAlert(alert models.Alert) (models.Alert, error) {
	return s.repo.CreateAlert(alert)
}
func (s *Service) GetAlert(id int) ([]models.Alert, error) {
	return s.repo.GetAlert(id)
}
func (s *Service) UpdateAlert(alert models.Alert) error {
	return s.repo.UpdateAlert(alert)
}
func (s *Service) DeleteAlert(id int) error {
	return s.repo.DeleteAlert(id)
}

func (s *Service) GetAllAlerts(client, eventType string) ([]models.Alert, error) {
	return s.repo.GetAllAlerts(client, eventType)
}
