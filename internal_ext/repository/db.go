package repository

import (
	infraDB "alert_monitor/infra/db"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var db = infraDB.GetDB()

// reads the metadata from migrations_metadata table and return last executed script name
func readMetadata() string {
	// create migrations table if not exists
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS migrations_metadata (migrated_at TIMESTAMP, script_name TEXT)")
	if err != nil {
		log.Fatal("error while creating table migrations_metadata", err)
	}
	// read last executed script name
	var lastScriptName string
	err = db.Get(&lastScriptName, "SELECT script_name from migrations_metadata ORDER BY migrated_at DESC LIMIT 1")
	if err != nil {
		return ""
	}
	return lastScriptName
}

// writeMetadata writes a new row in the migrations_metadata table to record our action
func writeMetadata(scriptName string) bool {
	sql := "INSERT INTO migrations_metadata values (NOW(), $1);"
	_, err := db.Exec(sql, scriptName)
	if err != nil {
		log.Println("error while wrting migrations metadata", err)
		return false
	}
	return true
}

// MigrateDB finds the last run migration, and run all those after it in order
func MigrateDB() {
	// Get a list of migration files
	files, err := filepath.Glob("migrations/*.sql")
	if err != nil {
		log.Printf("Error running restore %s", err)
		return
	}
	// Sort the list alphabetically
	sort.Strings(files)
	// get last run migration
	log.Println("Reading from Metadata table...")
	lastScriptName := readMetadata()
	log.Println("Last migrated script:", lastScriptName)

	var lastCompleted string
	completedCount := 0
	for _, file := range files {

		// if no migrations were made or the migration file is newer than last migrated file
		if lastScriptName == "" || strings.Compare(file, lastScriptName) > 0 {
			// reading contents of SQL file
			content, _ := os.ReadFile(file)
			// Convert []byte to string
			sqlQueries := string(content)

			// Execute queries in a transaction If at any point we fail, rollback it and break
			tx, _ := db.Begin()
			_, err = tx.Exec(sqlQueries)
			if err != nil {
				log.Println(sqlQueries)
				log.Println(err)
				tx.Rollback()
				break
			}
			tx.Commit()

			lastCompleted = file
			completedCount += 1

			log.Println("Completed migration:", file)
			writeMetadata(file)
		}
	}

	if completedCount > 0 {
		log.Println(completedCount, "Migrations completed. Last completed:", lastCompleted)
	} else {
		log.Println("No migrations performed")
	}
}
