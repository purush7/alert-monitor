package main

import (
	"alert_monitor/handler"
	"alert_monitor/internal_ext/repository"
	"fmt"
	"net/http"
)

func startHttpServer() {

	http.HandleFunc("/alerts", handler.AlertHandler)
	http.HandleFunc("/events", handler.EventHandler)
	fmt.Println("Alert System is running on :3336")
	http.ListenAndServe(":3333", nil)
}

func main() {
	// run migrations
	repository.MigrateDB()
	//start server
	startHttpServer()
}
