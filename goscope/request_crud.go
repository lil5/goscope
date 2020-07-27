package goscope

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/averageflow/goscope/repository"

	"github.com/gin-gonic/gin"
	uuid "github.com/nu7hatch/gouuid"
)

// Get all details from a request via its UID.
func GetDetailedRequest(requestUID string) DetailedRequest {
	var body string

	var headers string

	var result DetailedRequest

	row := repository.GetDetailedRequest(os.Getenv("GOSCOPE_DATABASE_TYPE"), requestUID)

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

	result.Body = html.UnescapeString(body)
	result.Headers = html.UnescapeString(headers)

	return result
}

func GetDetailedResponse(requestUID string) DetailedResponse {
	var body string

	var headers string

	var result DetailedResponse

	row := repository.GetDetailedResponse(os.Getenv("GOSCOPE_DATABASE_TYPE"), requestUID)

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

	result.Body = html.UnescapeString(body)
	result.Headers = html.UnescapeString(headers)

	return result
}

func GetRequests(c *gin.Context) {
	offsetQuery := c.DefaultQuery("offset", "0")
	offset, _ := strconv.ParseInt(offsetQuery, 10, 32)
	rows := repository.GetRequests(os.Getenv("GOSCOPE_DATABASE_TYPE"), int(offset))

	if rows == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error querying DB"})
		return
	}
	defer rows.Close()

	var result []SummarizedRequest

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

	c.JSON(http.StatusOK, result)
}

func DumpResponse(c *gin.Context, blw *BodyLogWriter, body string) {
	db := GetDB()
	defer db.Close()

	now := time.Now().Unix()
	requestUID, _ := uuid.NewV4()
	headers, _ := json.Marshal(c.Request.Header)
	query := "INSERT INTO requests (uid, application, client_ip, method, path, host, time, " +
		"headers, body, referrer, url, user_agent) VALUES " +
		"(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
	_, err := db.Exec(query, requestUID.String(), os.Getenv("APPLICATION_ID"), c.ClientIP(), c.Request.Method,
		c.FullPath(), c.Request.Host, now, html.EscapeString(string(headers)), html.EscapeString(body),
		c.Request.Referer(), c.Request.RequestURI, c.Request.UserAgent())

	if err != nil {
		log.Println(err.Error())
	}

	responseUID, _ := uuid.NewV4()
	headers, _ = json.Marshal(blw.Header())
	query = "INSERT INTO responses (uid, request_uid, application, client_ip, status, time, " +
		"body, path, headers, size) VALUES " +
		"(?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
	_, err = db.Exec(
		query,
		responseUID.String(),
		requestUID.String(),
		os.Getenv("APPLICATION_ID"),
		c.ClientIP(),
		blw.Status(),
		now,
		html.EscapeString(blw.body.String()),
		c.FullPath(),
		html.EscapeString(string(headers)),
		blw.body.Len(),
	)

	if err != nil {
		log.Println(err.Error())
	}
}

func SearchRequests(searchString string, offset int) []SummarizedRequest {
	var result []SummarizedRequest

	searchWildcard := fmt.Sprintf("%%%s%%", searchString)
	rows := repository.SearchRequests(os.Getenv("GOSCOPE_DATABASE_TYPE"), searchWildcard, offset)

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
