package watcher

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

func CheckExcludedPaths(path string) bool {
	result := true
	items := []string{"", "/watcher/", "css", "/watcher", "/js/*", "/css/*", "/css/*filepath", "/js/*filepath", "/watcher/requests", "/js", "/watcher/responses", "/watcher/responses/:id", "/watcher/requests/:id"}
	for _, s := range items {
		if path == s {
			result = false
		}
	}
	return result
}

func UnixTimeToAmsterdam(rawTime int) string {
	loc, _ := time.LoadLocation("Europe/Amsterdam")
	timeInstance := time.Unix(int64(rawTime), 0)
	return timeInstance.In(loc).Format("15:04:05 Mon, 2 Jan 2006 ")
}

func prettifyJson(rawString string) string {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, []byte(rawString), "", "\t")
	if err != nil {
		panic(err.Error())
	}
	return string(prettyJSON.Bytes())
}

func ReplaceVariablesInTemplate(rawTemplate string, variables map[string]string) string {
	for i, s := range variables {
		rawTemplate = strings.ReplaceAll(rawTemplate, fmt.Sprintf("{{.%s}}", i), s)
	}
	return rawTemplate
}