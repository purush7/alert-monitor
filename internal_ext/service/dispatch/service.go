package dispatch

import (
	"alert_monitor/internal_ext/models"
	"fmt"
	"sync"
)

var dispatcherService *Service = nil
var serviceOnce sync.Once

func GetService() UseCase {
	serviceOnce.Do(func() {
		dispatcherService = &Service{}
	})

	return dispatcherService
}

type Service struct {
}

func (s *Service) DispatchAlerts(alertId int, dispatchModes []models.DispatchStrategy, client, eventType string) {
	for _, strategy := range dispatchModes {
		switch strategy.Type {
		case models.CONSOLE:
			fmt.Printf("[INFO] AlertingService: Dispatching to Console\n")
			fmt.Printf("[WARN] Alert[%d]: `%s` for client: %s, event-type: %s\n", alertId, strategy.Message, client, eventType)
		case models.EMAIL:
			go s.dispatchEmail(alertId, strategy, client, eventType)
		}
	}
}

func (s *Service) dispatchEmail(alertId int, strategy models.DispatchStrategy, client, eventType string) {
	fmt.Printf("[INFO] AlertingService: Dispatching an Email\n")
	fmt.Printf("[INFO] Subject: %s\n", strategy.Subject)
	fmt.Printf("[WARN] Alert[%d]: for client: %s, event-type: %s\n", alertId, client, eventType)
}
