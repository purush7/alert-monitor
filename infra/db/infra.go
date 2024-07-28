package db

import (
	"alert_monitor/internal_ext/conf"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var database *sqlx.DB = nil
var dbOnce sync.Once

func GetDB() *sqlx.DB {
	var err error
	dbOnce.Do(func() {
		database, err = sqlx.Open("postgres", conf.GetDBConn())
		if err != nil {
			panic(err)
		}
		maxOpenConn := 50
		database.SetMaxOpenConns(maxOpenConn)
		database.SetMaxIdleConns(50)
		database.SetConnMaxIdleTime(2 * time.Minute)
		database.SetConnMaxLifetime(5 * time.Minute)
	})
	return database
}
