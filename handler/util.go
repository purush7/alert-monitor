package handler

import (
	"alert_monitor/internal_ext/service/alert"
	"alert_monitor/internal_ext/service/event"
)

var alertService = alert.GetService()
var eventService = event.GetService()
