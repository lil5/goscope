// License: MIT
// Authors:
// 		- Josep Jesus Bigorra Algaba (@averageflow)
package goscope

import (
	"bytes"
	"encoding/json"
	"log"
	"strings"
)

// Check the wanted path is not in the do not log list.
func CheckExcludedPaths(path string) bool {
	exactMatches := []string{
		"",
	}

	for _, s := range exactMatches {
		if path == s {
			return false
		}
	}

	partialMatches := []string{
		"/goscope",
	}

	for _, s := range partialMatches {
		if strings.Contains(path, s) {
			return false
		}
	}

	return true
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
