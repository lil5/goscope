package watcher

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/nu7hatch/gouuid"
	"net/http"
	"os"
	"strconv"
	"time"
)

func GetDetailedRequest(requestUid string) DetailedRequest {
	db, err := sql.Open("mysql", os.Getenv("WATCHER_DATABASE_CONNECTION"))
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	fmt.Println(fmt.Sprintf("SELECT * FROM `requests` WHERE `uid` = '%s' LIMIT 1;", requestUid))
	resultingQuery := fmt.Sprintf("SELECT * FROM `requests` WHERE `uid` = '%s' LIMIT 1;", requestUid)
	fmt.Println(resultingQuery)
	row := db.QueryRow(resultingQuery)

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

	err = row.Scan(&uid, &application, &clientIp, &method, &path, &url, &host, &t, &headers, &body, &referrer, &userAgent)
	if err != nil {
		panic(err.Error())
	}

	return DetailedRequest{
		Body:      body,
		ClientIp:  clientIp,
		Headers:   headers,
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

func GetDetailedResponse(responseUid string) DetailedResponse {
	db, err := sql.Open("mysql", os.Getenv("WATCHER_DATABASE_CONNECTION"))
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	resultingQuery := fmt.Sprintf("SELECT * FROM `responses` WHERE `uid` = '%s' LIMIT 1;", responseUid)
	fmt.Println(resultingQuery)
	row := db.QueryRow(resultingQuery)

	var application string
	var body string
	var clientIp string
	var headers string
	var path string
	var size int
	var status string
	var t int
	var uid string

	err = row.Scan(&uid, &application, &clientIp, &status, &t, &body, &path, &headers, &size)
	if err != nil {
		panic(err.Error())
	}
	return DetailedResponse{
		Body:     body,
		ClientIp: clientIp,
		Headers:  headers,
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
	db, err := sql.Open("mysql", os.Getenv("WATCHER_DATABASE_CONNECTION"))
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	query := "SELECT `uid`, `method`, `path`, `time` FROM `requests` WHERE `application` = '%s' ORDER BY `time` DESC LIMIT 15 OFFSET %d;"
	resultingQuery := fmt.Sprintf(query, os.Getenv("APPLICATION_ID"), offset)
	fmt.Println(resultingQuery)
	rows, _ := db.Query(resultingQuery)
	var result []SummarizedRequest
	for rows.Next() {
		var uid string
		var method string
		var path string
		var t int

		_ = rows.Scan(&uid, &method, &path, &t)
		request := SummarizedRequest{
			Method: method,
			Path:   path,
			Time:   t,
			Uid:    uid,
		}
		result = append(result, request)
	}
	c.JSON(http.StatusOK, result)
}


func DumpResponse(c *gin.Context,  blw *BodyLogWriter, body string) {
	db, err := sql.Open("mysql", os.Getenv("WATCHER_DATABASE_CONNECTION"))
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	now := time.Now().Unix()

	requestUid, _ := uuid.NewV4()
	headers, _ := json.Marshal(c.Request.Header)
	query := "INSERT INTO `requests` (`uid`, `application`, `client_ip`, `method`, `path`, `host`, `time`, `headers`, `body`, `referrer`, `url`, `user_agent`) VALUES " +
		"('%s', '%s', '%s', '%s', '%s', '%s', %v, '%s', '%s', '%s', '%s', '%s');"
	resultingQuery := fmt.Sprintf(query, requestUid, os.Getenv("APPLICATION_ID"), c.ClientIP(), c.Request.Method, c.FullPath(), c.Request.Host, now, headers, body,
		c.Request.Referer(), c.Request.RequestURI, c.Request.UserAgent())
	_, err = db.Exec(resultingQuery)
	if err != nil {
		panic(err.Error())
	}

	responseUid, _ := uuid.NewV4()
	headers, _ = json.Marshal(blw.Header())
	query = "INSERT INTO `responses` (`uid`, `request_uid`, `application`, `client_ip`, `status`, `time`, `body`, `path`, `headers`, `size`) VALUES " +
		"('%s', '%s', '%s', %v, %v, '%s', '%s', '%s', %v);"
	resultingQuery = fmt.Sprintf(query, responseUid, requestUid, os.Getenv("APPLICATION_ID"), c.ClientIP(), blw.Status(), now, blw.body.String(), c.FullPath(), headers, blw.body.Len())
	_, err = db.Exec(resultingQuery)
	if err != nil {
		panic(err.Error())
	}
}