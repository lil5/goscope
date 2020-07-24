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

	"github.com/gin-gonic/gin"
	uuid "github.com/nu7hatch/gouuid"
)

func GetDetailedRequest(requestUID string) DetailedRequest {
	db := GetDB()
	defer db.Close()

	resultingQuery := "SELECT uid, client_ip, method, path, " +
		"url, host, time, headers, body, referrer, user_agent " +
		"FROM requests WHERE uid = ? LIMIT 1;"
	row := db.QueryRow(resultingQuery, requestUID)

	var body string

	var headers string

	var result DetailedRequest

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
	db := GetDB()
	defer db.Close()

	resultingQuery := "SELECT uid, client_ip, status, time, " +
		"body, path, headers, size FROM responses " +
		"WHERE request_uid = ? LIMIT 1;"
	row := db.QueryRow(resultingQuery, requestUID)

	var body string

	var headers string

	var result DetailedResponse

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
	db := GetDB()

	defer db.Close()

	query := "SELECT requests.uid, requests.method, requests.path, requests.time, responses.status FROM requests " +
		"INNER JOIN responses ON requests.uid = responses.request_uid " +
		"WHERE requests.application = ? ORDER BY time DESC LIMIT %v OFFSET ?;"
	rows, err := db.Query(
		fmt.Sprintf(query, os.Getenv("GOSCOPE_ENTRIES_PER_PAGE")),
		os.Getenv("APPLICATION_ID"),
		offset,
	)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())

		return
	}

	defer rows.Close()

	var result []SummarizedRequest

	for rows.Next() {
		var request SummarizedRequest

		_ = rows.Scan(&request.UID, &request.Method, &request.Path, &request.Time, &request.ResponseStatus)
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

	db := GetDB()

	defer db.Close()

	query := "SELECT requests.uid, requests.method, requests.path, requests.time, responses.status FROM requests " +
		"INNER JOIN responses ON requests.uid = responses.request_uid WHERE requests.application = '%[2]v' AND " +
		"(requests.uid LIKE '%%%[3]v%%' OR requests.application LIKE '%%%[3]v%%' " +
		"OR requests.client_ip LIKE '%%%[3]v%%' OR requests.method LIKE '%%%[3]v%%' " +
		"OR requests.path LIKE '%%%[3]v%%' " +
		"OR requests.url LIKE '%%%[3]v%%' OR requests.host LIKE '%%%[3]v%%' " +
		"OR requests.body LIKE '%%%[3]v%%' OR requests.referrer LIKE '%%%[3]v%%' " +
		"OR requests.user_agent LIKE '%%%[3]v%%' OR requests.time LIKE '%%%[3]v%%'" +
		"OR responses.uid LIKE '%%%[3]v%%' OR responses.request_uid LIKE '%%%[3]v%%' " +
		"OR responses.application LIKE '%%%[3]v%%' OR responses.client_ip LIKE '%%%[3]v%%' " +
		"OR responses.status LIKE '%%%[3]v%%' " +
		"OR responses.body LIKE '%%%[3]v%%' OR responses.path LIKE '%%%[3]v%%' " +
		"OR responses.headers LIKE '%%%[3]v%%' OR responses.size LIKE '%%%[3]v%%' " +
		"OR responses.time LIKE '%%%[3]v%%') ORDER BY time DESC LIMIT %[1]v OFFSET %[4]v;"
	preparedQuery := fmt.Sprintf(
		query,
		os.Getenv("GOSCOPE_ENTRIES_PER_PAGE"),
		os.Getenv("APPLICATION_ID"),
		searchString,
		offset,
	)
	rows, err := db.Query(preparedQuery)

	if err != nil {
		log.Println(err.Error())
		return result
	}
	defer rows.Close()

	for rows.Next() {
		var request SummarizedRequest

		_ = rows.Scan(&request.UID, &request.Method, &request.Path, &request.Time, &request.ResponseStatus)

		result = append(result, request)
	}

	return result
}
