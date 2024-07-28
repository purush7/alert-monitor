package main

import (
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

// Generates a file in migrations/scripts/ directory in required migration format for a given tableName
func GenerateSQLFile(tableName string) {
	var sb strings.Builder
	timeString := time.Now().Format("20060102150405.003059_")
	regexString := regexp.MustCompile("^(.*?)\\.(.*)$")
	replaceString := "${1}$2"
	sb.WriteString("internal_ext/repo/migrations/")
	sb.WriteString(regexString.ReplaceAllString(timeString, replaceString))
	sb.WriteString(tableName)
	sb.WriteString(".sql")
	fileName := sb.String()
	emptyFile, err := os.Create(fileName)
	if err != nil {
		log.Println(err)
	}
	log.Println("Created SQL File:", fileName)
	emptyFile.Close()
}

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("Please specify a table name as first argument ")
	}
	tableName := os.Args[1]
	GenerateSQLFile(tableName)
}
