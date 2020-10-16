// License: MIT
// Authors:
// 		- Josep Jesus Bigorra Algaba (@averageflow)
package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/doug-martin/goqu/v9"

	"github.com/gin-gonic/gin"
	uuid "github.com/nu7hatch/gouuid"

	"github.com/averageflow/goscope/utils"
)

func GetDetailedRequest(db *sql.DB, connection, requestUID string) *sql.Row {
	qu := GetGoquDialect(connection)

	ds := qu.Select("uid", "client_ip", "method", "path", "url",
		"host", "time", "headers", "body", "referrer", "user_agent").From("requests").Where(goqu.C("uid").Eq(requestUID)).Limit(1)

	qsql, args, _ := ds.Prepared(true).ToSQL()

	row := db.QueryRow(qsql, args...)

	return row
}

func GetRequests(db *sql.DB, connection string, offset int) *sql.Rows {
	var query string
	if connection == MySQL || connection == SQLite {
		query = "SELECT `requests`.`uid`, `requests`.`method`, `requests`.`path`, `requests`.`time`, " +
			"`responses`.`status` FROM `requests` " +
			"INNER JOIN `responses` ON `requests`.`uid` = `responses`.`request_uid` " +
			"WHERE `requests`.`application` = ? ORDER BY `requests`.`time` DESC LIMIT ? OFFSET ?;"
	} else if connection == PostgreSQL {
		query = `SELECT "requests"."uid", "requests"."method", "requests"."path", "requests"."time",
			"responses"."status" FROM "requests"
			INNER JOIN "responses" ON "requests"."uid" = "responses"."request_uid"
			WHERE "requests"."application" = ? ORDER BY "requests"."time" DESC LIMIT ? OFFSET ?;`
	}

	rows, err := db.Query(
		query,
		utils.Config.ApplicationID,
		utils.Config.GoScopeEntriesPerPage,
		offset,
	)

	if err != nil {
		log.Println(err.Error())

		return nil
	}

	return rows
}

func SearchRequests(db *sql.DB, connection, search string, filter *RequestFilter, offset int) (*sql.Rows, error) {
	hasMethodFilter := false
	if filter != nil {
		hasMethodFilter = len(filter.Method) != 0
	}

	hasSearch := search != ""

	qu := GetGoquDialect(connection)

	ds := qu.Select(
		"requests.uid",
		"requests.method",
		"requests.path",
		"requests.time",
		"responses.status",
	).
		From("requests").
		InnerJoin(
			goqu.T("responses"),
			goqu.On(goqu.I("requests.uid").Eq(
				goqu.I("responses.request_uid"),
			)),
		).
		Where(
			goqu.I("requests.application").Eq(utils.Config.ApplicationID),
		)

	if hasMethodFilter && filter != nil {
		mLength := len(filter.Method)

		if mLength != 0 {
			mList := make([]interface{}, mLength)

			for i, v := range filter.Method {
				mList[i] = v
			}

			ds = ds.Where(
				goqu.I("requests.method").In(mList...),
			)
		}
	}

	if hasSearch {
		searchWildcard := fmt.Sprintf("%%%s%%", search)

		ds = ds.Where(goqu.Or(
			goqu.I("requests.uid").Like(searchWildcard),
			goqu.I("requests.application").Like(searchWildcard),
			goqu.I("requests.client_ip").Like(searchWildcard),
			goqu.I("requests.method").Like(searchWildcard),
			goqu.I("requests.path").Like(searchWildcard),
			goqu.I("requests.url").Like(searchWildcard),
			goqu.I("requests.host").Like(searchWildcard),
			goqu.I("requests.body").Like(searchWildcard),
			goqu.I("requests.referrer").Like(searchWildcard),
			goqu.I("requests.user_agent").Like(searchWildcard),
			goqu.I("requests.time").Like(searchWildcard),
			goqu.I("responses.uid").Like(searchWildcard),
			goqu.I("responses.request_uid").Like(searchWildcard),
			goqu.I("responses.application").Like(searchWildcard),
			goqu.I("responses.client_ip").Like(searchWildcard),
			goqu.I("responses.status").Like(searchWildcard),
			goqu.I("responses.body").Like(searchWildcard),
			goqu.I("responses.path").Like(searchWildcard),
			goqu.I("responses.headers").Like(searchWildcard),
			goqu.I("responses.size").Like(searchWildcard),
			goqu.I("responses.time").Like(searchWildcard),
		))
	}

	ds = ds.Order(
		goqu.I("requests.time").Desc(),
	).Limit(uint(offset))

	qsql, args, _ := ds.Prepared(true).ToSQL()

	rows, err := db.Query(qsql, args...)

	if err != nil {
		return nil, err
	}

	return rows, nil
}

func GetDetailedResponse(db *sql.DB, connection, requestUID string) *sql.Row {
	var query string

	if connection == MySQL || connection == SQLite {
		query = "SELECT `uid`, `client_ip`, `status`, `time`, " +
			"`body`, `path`, `headers`, `size` FROM `responses` " +
			"WHERE `request_uid` = ? LIMIT 1;"
	} else if connection == PostgreSQL {
		query = `SELECT "uid", "client_ip", "status", "time",
			"body", "path", "headers", "size" FROM "responses"
			WHERE "request_uid" = ? LIMIT 1;`
	}

	row := db.QueryRow(query, requestUID)

	return row
}

func DumpResponse(c *gin.Context, responsePayload utils.DumpResponsePayload, body string) {
	now := time.Now().Unix()
	requestUID, _ := uuid.NewV4()
	headers, _ := json.Marshal(c.Request.Header)
	query := "INSERT INTO requests (uid, application, client_ip, method, path, host, time, " +
		"headers, body, referrer, url, user_agent) VALUES " +
		"(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"

	requestPath := c.FullPath()
	if requestPath == "" {
		requestPath = c.Request.URL.String()
	}

	_, err := utils.DB.Exec(
		query,
		requestUID.String(),
		utils.Config.ApplicationID,
		c.ClientIP(),
		c.Request.Method,
		requestPath,
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
	headers, _ = json.Marshal(responsePayload.Headers)
	query = "INSERT INTO responses (uid, request_uid, application, client_ip, status, time, " +
		"body, path, headers, size) VALUES " +
		"(?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
	_, err = utils.DB.Exec(
		query,
		responseUID.String(),
		requestUID.String(),
		utils.Config.ApplicationID,
		c.ClientIP(),
		responsePayload.Status,
		now,
		responsePayload.Body.String(),
		c.FullPath(),
		string(headers),
		responsePayload.Body.Len(),
	)

	if err != nil {
		log.Println(err.Error())
	}
}
