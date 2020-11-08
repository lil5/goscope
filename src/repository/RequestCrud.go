package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/averageflow/goscope/src/types"

	"github.com/gin-gonic/gin"
	uuid "github.com/nu7hatch/gouuid"

	"github.com/averageflow/goscope/src/utils"
)

// Get all details from a request via its UID.
func GetDetailedRequest(requestUID string) types.DetailedRequest {
	var body string

	var headers string

	var result types.DetailedRequest

	row := QueryDetailedRequest(utils.DB, utils.Config.GoScopeDatabaseType, requestUID)

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

	result.Body = utils.PrettifyJSON(body)
	result.Headers = utils.PrettifyJSON(headers)

	return result
}

func GetDetailedResponse(requestUID string) types.DetailedResponse {
	var body string

	var headers string

	var result types.DetailedResponse

	row := QueryDetailedResponse(utils.DB, utils.Config.GoScopeDatabaseType, requestUID)

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

	result.Body = utils.PrettifyJSON(body)
	result.Headers = utils.PrettifyJSON(headers)

	return result
}

func GetRequests(offset int) []types.SummarizedRequest {
	var result []types.SummarizedRequest

	rows := QueryGetRequests(utils.DB, utils.Config.GoScopeDatabaseType, offset)

	if rows == nil {
		return result
	}
	defer rows.Close()

	for rows.Next() {
		var request types.SummarizedRequest

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

func SearchRequests(search string, filter *types.RequestFilter, offset int) []types.SummarizedRequest {
	var result []types.SummarizedRequest

	rows, err := QuerySearchRequests(utils.DB, utils.Config.GoScopeDatabaseType, search, filter, offset)

	if err != nil {
		log.Println(err.Error())
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var request types.SummarizedRequest

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

func QueryDetailedRequest(db *sql.DB, connection, requestUID string) *sql.Row {
	query := `
		SELECT uid,
		   client_ip,
		   method,
		   path,
		   url,
		   host,
		   time,
		   headers,
		   body,
		   referrer,
		   user_agent
		FROM requests
		WHERE uid = ?
		LIMIT 1;
	`

	row := db.QueryRow(query, requestUID)

	return row
}

func QueryGetRequests(db *sql.DB, connection string, offset int) *sql.Rows {
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

func QuerySearchRequests(db *sql.DB, connection, search string, //nolint:gocognit,funlen,gocyclo
	filter *types.RequestFilter, offset int) (*sql.Rows, error) { //nolint:gocognit,funlen,gocyclo
	var query string

	var methodQuery string

	var searchQuery string

	var methodSQL []string

	hasMethodFilter := false
	if filter != nil {
		hasMethodFilter = len(filter.Method) != 0
	}

	hasSearch := search != ""

	var searchQueryCols [][2]string

	var searchWildcard string

	if hasSearch {
		searchWildcard = fmt.Sprintf("%%%s%%", search)

		searchQueryCols = [][2]string{
			{"requests", "uid"},
			{"requests", "application"},
			{"requests", "client_ip"},
			{"requests", "method"},
			{"requests", "path"},
			{"requests", "url"},
			{"requests", "host"},
			{"requests", "body"},
			{"requests", "referrer"},
			{"requests", "user_agent"},
			{"requests", "time"},
			{"responses", "uid"},
			{"responses", "request_uid"},
			{"responses", "application"},
			{"responses", "client_ip"},
			{"responses", "status"},
			{"responses", "body"},
			{"responses", "path"},
			{"responses", "headers"},
			{"responses", "size"},
			{"responses", "time"},
		}
	}

	if connection == MySQL || connection == SQLite { //nolint:nestif
		if hasMethodFilter && filter != nil {
			for i, m := range filter.Method {
				if i == 0 {
					methodQuery += "AND (`requests`.`method` = ? "
				} else {
					methodQuery += "OR `requests`.`method` = ? "
				}

				methodSQL = append(methodSQL, m)
			}

			methodQuery += ") " //nolint:goconst
		}

		if hasSearch {
			searchQuery += "AND (" //nolint:goconst

			for i, col := range searchQueryCols {
				if i != 0 {
					searchQuery += "OR " //nolint:goconst
				}

				searchQuery += fmt.Sprintf("`%s`.`%s` LIKE ? ", col[0], col[1])
			}

			searchQuery += ") "
		}

		query = "SELECT `requests`.`uid`, `requests`.`method`, `requests`.`path`, `requests`.`time`, " +
			"`responses`.`status` FROM `requests` " +
			"INNER JOIN `responses` ON `requests`.`uid` = `responses`.`request_uid` " +
			"WHERE `requests`.`application` = ? " +
			methodQuery +
			searchQuery +
			"ORDER BY `requests`.`time` DESC LIMIT ? OFFSET ?;"
	} else if connection == PostgreSQL {
		if hasMethodFilter && filter != nil {
			for i, m := range filter.Method {
				if i == 0 {
					methodQuery += `AND ("requests"."method" = ? `
				} else {
					methodQuery += `OR "requests"."method" = ? `
				}
				methodSQL = append(methodSQL, m)
			}
			methodQuery += `) `
		}

		if hasSearch {
			searchQuery += "AND ("
			for i, col := range searchQueryCols {
				if i != 0 {
					searchQuery += "OR "
				}
				searchQuery += fmt.Sprintf(`"%s"."%s" LIKE ? `, col[0], col[1])
			}
			searchQuery += ") "
		}

		query = `SELECT "requests"."uid", "requests"."method", "requests"."path",
			"requests"."time", "responses"."status" FROM "requests"
			INNER JOIN "responses" ON "requests"."uid" = "responses"."request_uid"
			WHERE "requests"."application" = ?
			` + methodQuery + searchQuery + `
			ORDER BY "requests"."time" DESC LIMIT ? OFFSET ?;`
	}

	var args []interface{}
	args = append(args, utils.Config.ApplicationID)

	if hasMethodFilter && filter != nil {
		for _, ms := range methodSQL {
			args = append(args, ms)
		}
	}

	if hasSearch {
		args = append(args,
			searchWildcard,
			searchWildcard,
			searchWildcard,
			searchWildcard,
			searchWildcard,
			searchWildcard,
			searchWildcard,
			searchWildcard,
			searchWildcard,
			searchWildcard,
			searchWildcard,
			searchWildcard,
			searchWildcard,
			searchWildcard,
			searchWildcard,
			searchWildcard,
			searchWildcard,
			searchWildcard,
			searchWildcard,
			searchWildcard,
			searchWildcard,
		)
	}

	args = append(
		args,
		utils.Config.GoScopeEntriesPerPage,
		offset)

	rows, err := db.Query(query, args...)

	if err != nil {
		return nil, err
	}

	return rows, nil
}

func QueryDetailedResponse(db *sql.DB, connection, requestUID string) *sql.Row {
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

func DumpResponse(c *gin.Context, responsePayload types.DumpResponsePayload, body string) {
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
