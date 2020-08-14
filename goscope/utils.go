// License: MIT
// Authors:
// 		- Josep Bigorra (averageflow)
package goscope

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"os"
)

// Check the wanted path is not in the do not log list.
func CheckExcludedPaths(path string) bool {
	result := true
	items := []string{
		"",
		"/goscope/",
		"/goscope/log-records",
		"/goscope/log-records/:id",
		"/goscope/logs",
		"css",
		"/goscope",
		"/js/*",
		"/css/*",
		"/css/*filepath",
		"/js/*filepath",
		"/goscope/requests",
		"/js",
		"/goscope/responses",
		"/goscope/responses/:id",
		"/goscope/requests/:id",
		"/goscope/search/requests",
		"/goscope/search/logs",
		"/goscope/api/requests",
		"/goscope/api/requests/:id",
		"/goscope/api/logs",
		"/goscope/api/logs/:id",
		"/goscope/api/info",
		"/goscope/api/search/logs",
		"/goscope/api/search/requests",
		"/goscope/runtime.js",
		"/goscope/styles.css",
		"/goscope/polyfills.js",
		"/goscope/main.js",
		"/goscope/api/application-name",
		"/goscope/favicon.ico",
	}

	for _, s := range items {
		if path == s {
			result = false
		}
	}

	return result
}

func prettifyJSON(rawString string) string {
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
