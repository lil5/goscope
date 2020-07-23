// License: MIT
// Authors:
// 		- Josep Bigorra (averageflow)
package goscope

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/nu7hatch/gouuid"
	"html"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func GetDetailedRequest(requestUid string) DetailedRequest {
	db := GetDB()
	defer db.Close()
	resultingQuery := "SELECT uid, application, client_ip, method, path, url, host, time, headers, body, referrer, user_agent FROM requests WHERE uid = ? LIMIT 1;"
	row := db.QueryRow(resultingQuery, requestUid)
	var application string
	var body string
	var clientIp string
	var headers string
	var host string
	var method string
	var path string
	var referrer string
	var t int
	var uid string
	var url string
	var userAgent string

	err := row.Scan(&uid, &application, &clientIp, &method, &path, &url, &host, &t, &headers, &body, &referrer, &userAgent)
	if err != nil {
		log.Println(err.Error())
	}

	return DetailedRequest{
		Body:      html.UnescapeString(body),
		ClientIp:  clientIp,
		Headers:   html.UnescapeString(headers),
		Host:      host,
		Method:    method,
		Path:      path,
		Referrer:  referrer,
		Time:      t,
		Uid:       uid,
		Url:       url,
		UserAgent: userAgent,
	}
}

func GetDetailedResponse(requestUid string) DetailedResponse {
	db := GetDB()
	defer db.Close()
	resultingQuery := "SELECT uid, application, client_ip, status, time, body, path, headers, size FROM responses WHERE request_uid = ? LIMIT 1;"
	row := db.QueryRow(resultingQuery, requestUid)

	var application string
	var body string
	var clientIp string
	var headers string
	var path string
	var size int
	var status string
	var t int
	var uid string

	err := row.Scan(&uid, &application, &clientIp, &status, &t, &body, &path, &headers, &size)
	if err != nil {
		log.Println(err.Error())
	}
	return DetailedResponse{
		Body:     html.UnescapeString(body),
		ClientIp: clientIp,
		Headers:  html.UnescapeString(headers),
		Path:     path,
		Size:     size,
		Status:   status,
		Time:     t,
		Uid:      uid,
	}
}

func GetRequests(c *gin.Context) {
	offsetQuery := c.DefaultQuery("offset", "0")
	offset, _ := strconv.ParseInt(offsetQuery, 10, 32)
	db := GetDB()
	defer db.Close()
	query := "SELECT requests.uid, requests.method, requests.path, requests.time, responses.status FROM requests " +
		"INNER JOIN responses ON requests.uid = responses.request_uid WHERE requests.application = ? ORDER BY time DESC LIMIT %v OFFSET ?;"
	rows, err := db.Query(fmt.Sprintf(query, os.Getenv("GOSCOPE_ENTRIES_PER_PAGE")), os.Getenv("APPLICATION_ID"), offset)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	var result []SummarizedRequest
	for rows.Next() {
		var uid string
		var method string
		var path string
		var t int
		var status int

		_ = rows.Scan(&uid, &method, &path, &t, &status)
		request := SummarizedRequest{
			Method:         method,
			Path:           path,
			Time:           t,
			Uid:            uid,
			ResponseStatus: status,
		}
		result = append(result, request)
	}
	c.JSON(http.StatusOK, result)
}

func GetDetailedLog(requestUid string) ExceptionRecord {
	db := GetDB()
	defer db.Close()
	query := "SELECT uid, error, time FROM logs WHERE uid = ?;"
	row := db.QueryRow(query, requestUid)
	var uid string
	var t int
	var errorMessage string
	err := row.Scan(&uid, &errorMessage, &t)
	if err != nil {
		log.Println(err.Error())
		return ExceptionRecord{}
	}
	return ExceptionRecord{
		Error: errorMessage,
		Time:  t,
		Uid:   uid,
	}
}

func GetLogs(c *gin.Context) {
	offsetQuery := c.DefaultQuery("offset", "0")
	offset, _ := strconv.ParseInt(offsetQuery, 10, 32)
	db := GetDB()
	defer db.Close()
	query := "SELECT uid, CASE WHEN LENGTH(error) > 80 THEN CONCAT(SUBSTRING(error, 1, 80), '...') " +
		"ELSE error END AS error, time FROM logs WHERE application = ? " +
		"ORDER BY time DESC LIMIT %v OFFSET ?;"
	rows, err := db.Query(fmt.Sprintf(query, os.Getenv("GOSCOPE_ENTRIES_PER_PAGE")), os.Getenv("APPLICATION_ID"), offset)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	var result []ExceptionRecord
	for rows.Next() {
		var uid string
		var t int
		var errorMessage string

		err := rows.Scan(&uid, &errorMessage, &t)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		request := ExceptionRecord{
			Error: errorMessage,
			Time:  t,
			Uid:   uid,
		}
		result = append(result, request)
	}
	c.JSON(http.StatusOK, result)
}

func DumpResponse(c *gin.Context, blw *BodyLogWriter, body string) {
	db := GetDB()
	defer db.Close()
	now := time.Now().Unix()
	requestUid, _ := uuid.NewV4()
	headers, _ := json.Marshal(c.Request.Header)
	query := "INSERT INTO requests (uid, application, client_ip, method, path, host, time, " +
		"headers, body, referrer, url, user_agent) VALUES " +
		"(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
	_, err := db.Exec(query, requestUid.String(), os.Getenv("APPLICATION_ID"), c.ClientIP(), c.Request.Method,
		c.FullPath(), c.Request.Host, now, html.EscapeString(string(headers)), html.EscapeString(body),
		c.Request.Referer(), c.Request.RequestURI, c.Request.UserAgent())
	if err != nil {
		log.Println(err.Error())
	}
	responseUid, _ := uuid.NewV4()
	headers, _ = json.Marshal(blw.Header())
	query = "INSERT INTO responses (uid, request_uid, application, client_ip, status, time, " +
		"body, path, headers, size) VALUES " +
		"(?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
	_, err = db.Exec(query, responseUid.String(), requestUid.String(), os.Getenv("APPLICATION_ID"), c.ClientIP(),
		blw.Status(), now, html.EscapeString(blw.body.String()), c.FullPath(), html.EscapeString(string(headers)), blw.body.Len())
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
		"(requests.uid LIKE '%%%[3]v%%' OR requests.application LIKE '%%%[3]v%%' OR requests.client_ip LIKE '%%%[3]v%%' OR requests.method LIKE '%%%[3]v%%' OR requests.path LIKE '%%%[3]v%%' " +
		"OR requests.url LIKE '%%%[3]v%%' OR requests.host LIKE '%%%[3]v%%' OR requests.body LIKE '%%%[3]v%%' OR requests.referrer LIKE '%%%[3]v%%' OR requests.user_agent LIKE '%%%[3]v%%' OR requests.time LIKE '%%%[3]v%%'" +
		"OR responses.uid LIKE '%%%[3]v%%' OR responses.request_uid LIKE '%%%[3]v%%' OR responses.application LIKE '%%%[3]v%%' OR responses.client_ip LIKE '%%%[3]v%%' OR responses.status LIKE '%%%[3]v%%' " +
		"OR responses.body LIKE '%%%[3]v%%' OR responses.path LIKE '%%%[3]v%%' OR responses.headers LIKE '%%%[3]v%%' OR responses.size LIKE '%%%[3]v%%' OR responses.time LIKE '%%%[3]v%%') " +
		"ORDER BY time DESC LIMIT %[1]v OFFSET %[4]v;"
	preparedQuery := fmt.Sprintf(query, os.Getenv("GOSCOPE_ENTRIES_PER_PAGE"), os.Getenv("APPLICATION_ID"), searchString, offset)
	rows, err := db.Query(preparedQuery)
	if err != nil {
		log.Println(err.Error())
		return result
	}
	for rows.Next() {
		var uid string
		var method string
		var path string
		var t int
		var status int

		_ = rows.Scan(&uid, &method, &path, &t, &status)
		request := SummarizedRequest{
			Method:         method,
			Path:           path,
			Time:           t,
			Uid:            uid,
			ResponseStatus: status,
		}
		result = append(result, request)
	}
	return result
}

func SearchLogs(searchString string, offset int) []ExceptionRecord {
	var result []ExceptionRecord
	db := GetDB()
	defer db.Close()
	query := "SELECT uid, CASE WHEN LENGTH(error) > 80 THEN CONCAT(SUBSTRING(error, 1, 80), '...') " +
		"ELSE error END AS error, time FROM logs WHERE application = '%[1]v' AND (uid LIKE '%%%[3]v%%' OR application LIKE '%%%[3]v%%' OR error LIKE '%%%[3]v%%' OR time LIKE '%%%[3]v%%') " +
		"ORDER BY time DESC LIMIT %[2]v OFFSET %[4]v;"
	preparedQuery := fmt.Sprintf(query, os.Getenv("APPLICATION_ID"), os.Getenv("GOSCOPE_ENTRIES_PER_PAGE"), searchString, offset)
	rows, err := db.Query(preparedQuery)
	if err != nil {
		log.Println(err.Error())
		return result
	}
	for rows.Next() {
		var uid string
		var t int
		var errorMessage string

		err := rows.Scan(&uid, &errorMessage, &t)
		if err != nil {
			log.Println(err.Error())
			return result
		}
		request := ExceptionRecord{
			Error: errorMessage,
			Time:  t,
			Uid:   uid,
		}
		result = append(result, request)
	}
	return result
}
