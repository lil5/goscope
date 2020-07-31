// License: MIT
// Authors:
// 		- Josep Bigorra (averageflow)
package goscope

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
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

func UnixTimeToHuman(rawTime int) string {
	loc, err := time.LoadLocation(os.Getenv("APPLICATION_TIMEZONE"))
	if err != nil {
		log.Println(err.Error())

		loc, _ = time.LoadLocation("Europe/Amsterdam")
	}

	timeInstance := time.Unix(int64(rawTime), 0)

	return timeInstance.In(loc).Format("15:04:05 Mon, 2 Jan 2006 ")
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

func MinifyCSS(uncompressed string) string {
	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	minified, err := m.String("text/css", uncompressed)

	if err != nil {
		log.Println(err.Error())
		return ""
	}

	return minified
}

func MinifyJs(uncompressed string) string {
	m := minify.New()
	m.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)
	minified, err := m.String("application/javascript", uncompressed)

	if err != nil {
		log.Println(err.Error())
		return ""
	}

	return minified
}

func MinifyHTML(uncompressed string) string {
	m := minify.New()
	m.AddFunc("text/html", html.Minify)
	minified, err := m.String("text/html", uncompressed)

	if err != nil {
		log.Println(err.Error())
		return ""
	}

	return minified
}

func ShowGoScopePage(c *gin.Context, rawTemplate string, variables map[string]string) {
	cleanTemplate := ReplaceVariablesInTemplate(rawTemplate, variables)
	reader := strings.NewReader(cleanTemplate)
	c.DataFromReader(http.StatusOK, reader.Size(), "text/html", reader, nil)
}
