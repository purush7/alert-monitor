package event

import (
	"alert_monitor/internal_ext/models"
	"alert_monitor/internal_ext/repository"
	"alert_monitor/internal_ext/service/alert"
	"alert_monitor/internal_ext/service/dispatch"
	"fmt"
	"sync"
	"time"
)

var eventService *Service = nil
var serviceOnce sync.Once

func GetService() UseCase {
	serviceOnce.Do(func() {
		eventService = &Service{
			repo:          repository.GetEventsRepo(),
			dispatcherSrv: dispatch.GetService(),
			alertSrv:      alert.GetService(),
		}
	})

	return eventService
}

type Service struct {
	repo          repo
	dispatcherSrv dispatcher
	alertSrv      alert.UseCase
}

func (s *Service) Event(event models.Event) (models.Event, error) {

	alerts, err := s.alertSrv.GetAllAlerts(event.Client, event.EventType)
	if err != nil {
		return event, fmt.Errorf("Error retrieving alerts: %s", err)
	}

	for _, alert := range alerts {
		// Step 3: Check the count based on alert configuration
		var thresholdBreached bool

		switch alert.AlertConfig.Type {
		case models.SIMPLE_COUNT:
			thresholdBreached = s.checkSimpleCount(alert)
		case models.TUMBLING_WINDOW:
			thresholdBreached = s.checkTumblingWindow(alert)
		case models.SLIDING_WINDOW:
			thresholdBreached = s.checkSlidingWindow(alert)
		default:
			fmt.Printf("Unknown alert type: %s\n", alert.AlertConfig.Type)
			continue
		}

		// If threshold is breached, dispatch alerts
		if thresholdBreached {
			go func(alert models.Alert) {
				s.dispatcherSrv.DispatchAlerts(alert.ID, alert.DispatcherConfig, alert.Client, alert.EventType)
			}(alert)
		}
	}
	return s.repo.LogEvent(event)
}

// checkSimpleCount checks if the number of events exceeds the count threshold.
func (s *Service) checkSimpleCount(alert models.Alert) bool {
	count, err := s.repo.CountEvents(alert.Client, alert.EventType)
	if err != nil {
		return false
	}
	return count+1 >= alert.AlertConfig.Count
}

func getTumblingWindowStart(now time.Time, windowSizeInSecs int) time.Time {

	startOfDay := now.Truncate(24 * time.Hour)
	elapsedSeconds := int(now.Sub(startOfDay).Seconds())
	windowIndex := elapsedSeconds / windowSizeInSecs
	windowStartSeconds := windowIndex * windowSizeInSecs
	return startOfDay.Add(time.Duration(windowStartSeconds) * time.Second)

}

// checkTumblingWindow checks if the number of events in the current window exceeds the count threshold.
func (s *Service) checkTumblingWindow(alert models.Alert) bool {
	now := time.Now()
	windowStart := getTumblingWindowStart(now, alert.AlertConfig.WindowSizeInSecs)
	count, err := s.repo.CountEventsInRange(alert.Client, alert.EventType, windowStart, now)
	if err != nil {
		return false
	}
	return count+1 >= alert.AlertConfig.Count
}

// checkSlidingWindow checks if the number of events in the sliding window exceeds the count threshold.
func (s *Service) checkSlidingWindow(alert models.Alert) bool {
	now := time.Now()
	startTime := now.Add(-time.Duration(alert.AlertConfig.WindowSizeInSecs) * time.Second)
	count, err := s.repo.CountEventsInRange(alert.Client, alert.EventType, startTime, now)
	if err != nil {
		return false
	}
	return count+1 >= alert.AlertConfig.Count
}
