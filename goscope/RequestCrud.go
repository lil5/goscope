// License: MIT
// Authors:
// 		- Josep Jesus Bigorra Algaba (@averageflow)
package goscope

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/averageflow/goscope/utils"

	"github.com/averageflow/goscope/database"

	"github.com/gin-gonic/gin"
	uuid "github.com/nu7hatch/gouuid"
)

// Get all details from a request via its UID.
func GetDetailedRequest(requestUID string) DetailedRequest {
	var body string

	var headers string

	var result DetailedRequest

	row := database.GetDetailedRequest(utils.DB, utils.Config.GoScopeDatabaseType, requestUID)

	err := row.Scan(
		&result.UID,
		&result.ClientIP,
		&result.Method,
		&result.Path,
		&result.URL,
		&result.Host,
		&result.Time,
		&headers,
		&body,
		&result.Referrer,
		&result.UserAgent,
	)
	if err != nil {
		log.Println(err.Error())
	}

	result.Body = prettifyJSON(body)
	result.Headers = prettifyJSON(headers)

	return result
}

func GetDetailedResponse(requestUID string) DetailedResponse {
	var body string

	var headers string

	var result DetailedResponse

	row := database.GetDetailedResponse(utils.DB, utils.Config.GoScopeDatabaseType, requestUID)

	err := row.Scan(
		&result.UID,
		&result.ClientIP,
		&result.Status,
		&result.Time,
		&body,
		&result.Path,
		&headers,
		&result.Size,
	)
	if err != nil {
		log.Println(err.Error())
	}

	result.Body = prettifyJSON(body)
	result.Headers = prettifyJSON(headers)

	return result
}

func GetRequests(offset int) []SummarizedRequest {
	var result []SummarizedRequest

	rows := database.GetRequests(utils.DB, utils.Config.GoScopeDatabaseType, offset)

	if rows == nil {
		return result
	}
	defer rows.Close()

	for rows.Next() {
		var request SummarizedRequest

		err := rows.Scan(
			&request.UID,
			&request.Method,
			&request.Path,
			&request.Time,
			&request.ResponseStatus,
		)
		if err != nil {
			log.Println(err.Error())
			return result
		}

		result = append(result, request)
	}

	return result
}

func DumpResponse(c *gin.Context, blw *BodyLogWriter, body string) {
	now := time.Now().Unix()
	requestUID, _ := uuid.NewV4()
	headers, _ := json.Marshal(c.Request.Header)
	query := "INSERT INTO requests (uid, application, client_ip, method, path, host, time, " +
		"headers, body, referrer, url, user_agent) VALUES " +
		"(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
	_, err := utils.DB.Exec(
		query,
		requestUID.String(),
		utils.Config.ApplicationID,
		c.ClientIP(),
		c.Request.Method,
		c.FullPath(),
		c.Request.Host,
		now,
		string(headers),
		body,
		c.Request.Referer(),
		c.Request.RequestURI,
		c.Request.UserAgent(),
	)

	if err != nil {
		log.Println(err.Error())
	}

	responseUID, _ := uuid.NewV4()
	headers, _ = json.Marshal(blw.Header())
	query = "INSERT INTO responses (uid, request_uid, application, client_ip, status, time, " +
		"body, path, headers, size) VALUES " +
		"(?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
	_, err = utils.DB.Exec(
		query,
		responseUID.String(),
		requestUID.String(),
		utils.Config.ApplicationID,
		c.ClientIP(),
		blw.Status(),
		now,
		blw.body.String(),
		c.FullPath(),
		string(headers),
		blw.body.Len(),
	)

	if err != nil {
		log.Println(err.Error())
	}
}

func SearchRequests(searchString string, offset int) []SummarizedRequest {
	var result []SummarizedRequest

	searchWildcard := fmt.Sprintf("%%%s%%", searchString)
	rows := database.SearchRequests(utils.DB, utils.Config.GoScopeDatabaseType, searchWildcard, offset)

	defer rows.Close()

	for rows.Next() {
		var request SummarizedRequest

		err := rows.Scan(
			&request.UID,
			&request.Method,
			&request.Path,
			&request.Time,
			&request.ResponseStatus,
		)

		if err != nil {
			log.Println(err.Error())
		}

		result = append(result, request)
	}

	return result
}
