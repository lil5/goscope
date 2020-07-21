// License: MIT
// Authors:
// 		- Josep Bigorra (averageflow)
package goscope

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
	"log"
	"os"
	"regexp"
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
func MinifyCss(uncompressed string) string {
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

func MinifyHtml(uncompressed string) string {
	m := minify.New()
	m.AddFunc("text/html", html.Minify)
	minified, err := m.String("text/html", uncompressed)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	return minified
}

func NavbarComponent(selected string) string {
	navlinkKeys := []string {
		"REQUESTS", "LOGS",
	}
	navLinks := []string{
		"<a class=\"%s\" href=\"/goscope/\"><h3 style=\"margin: 1.2em;\" class=\"font-l\">🌐&nbsp;&nbsp;Requests</h3></a>",
		"<a class=\"%s\" href=\"/goscope/logs\"><h3 style=\"margin: 1.2em;\" class=\"font-l\">📃&nbsp;&nbsp;Logs</h3></a>",
	}
	navbarCode := `
	<div class="flex m-2 p-2">
		<h3 class="font-l" style="margin: 1.2em;">⚙️&nbsp;%s</h3>
		%s
	</div>
	`
	var generatedLinks string
	for i, s := range navlinkKeys {
		if s == selected {
			generatedLinks += fmt.Sprintf(navLinks[i], "active-navbar-link")
		} else {
			generatedLinks += fmt.Sprintf(navLinks[i], "navbar-link")
		}
	}
	return MinifyHtml(fmt.Sprintf(navbarCode, os.Getenv("APPLICATION_NAME"), generatedLinks))
}