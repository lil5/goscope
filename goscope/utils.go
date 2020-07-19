// Copyright © 2020 Pro Warehouse B.V.
// All Rights Reserved
package goscope

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func CheckExcludedPaths(path string) bool {
	result := true
	items := []string{"", "/goscope/", "/goscope/log-records", "/goscope/log-records/:id", "/goscope/logs", "css", "/goscope", "/js/*", "/css/*", "/css/*filepath", "/js/*filepath", "/goscope/requests", "/js", "/goscope/responses", "/goscope/responses/:id", "/goscope/requests/:id"}
	for _, s := range items {
		if path == s {
			result = false
		}
	}
	return result
}

func UnixTimeToHuman(rawTime int) string {
	loc, err := time.LoadLocation(os.Getenv("APPLICATION_TIMEZONE"))
	if err != nil {
		log.Println(err.Error())
		loc, _ = time.LoadLocation("Europe/Amsterdam")
	}
	timeInstance := time.Unix(int64(rawTime), 0)
	return timeInstance.In(loc).Format("15:04:05 Mon, 2 Jan 2006 ")
}

func prettifyJson(rawString string) string {
	if rawString == "" {
		return ""
	}
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, []byte(rawString), "", "    ")
	if err != nil {
		log.Println(err.Error())
		return rawString
	}
	return prettyJSON.String()
}

func ReplaceVariablesInTemplate(rawTemplate string, variables map[string]string) string {
	for i, s := range variables {
		rawTemplate = strings.ReplaceAll(rawTemplate, fmt.Sprintf("{{.%s}}", i), s)
	}
	return rawTemplate
}

func GetDB() *sql.DB {
	db, err := sql.Open(os.Getenv("GOSCOPE_DATABASE_TYPE"), os.Getenv("GOSCOPE_DATABASE_CONNECTION"))
	if err != nil {
		log.Println(err.Error())
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Println(err.Error())
		panic(err.Error())
	}
	return db
}
