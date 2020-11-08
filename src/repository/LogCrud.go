package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/averageflow/goscope/src/types"

	uuid "github.com/nu7hatch/gouuid"

	"github.com/averageflow/goscope/src/utils"
)

func FetchDetailedLog(requestUID string) types.ExceptionRecord {
	row := QueryDetailedLog(
		utils.DB,
		utils.Config.GoScopeDatabaseType,
		requestUID,
	)

	var request types.ExceptionRecord

	err := row.Scan(&request.UID, &request.Error, &request.Time)
	if err != nil {
		log.Println(err.Error())
		return request
	}

	return request
}

func FetchSearchLogs(searchString string, offset int) []types.ExceptionRecord {
	var result []types.ExceptionRecord

	searchWildcard := fmt.Sprintf("%%%s%%", searchString)
	rows := QuerySearchLogs(utils.DB, utils.Config.GoScopeDatabaseType, searchWildcard, offset)

	defer rows.Close()

	for rows.Next() {
		var request types.ExceptionRecord

		err := rows.Scan(&request.UID, &request.Error, &request.Time)
		if err != nil {
			log.Println(err.Error())
			return result
		}

		result = append(result, request)
	}

	return result
}

// Get a summarized list of application logs from the DB.
func FetchLogs(offset int) []types.ExceptionRecord {
	var result []types.ExceptionRecord

	rows := QueryGetLogs(utils.DB, utils.Config.GoScopeDatabaseType, offset)
	if rows == nil {
		return result
	}

	defer rows.Close()

	for rows.Next() {
		var request types.ExceptionRecord

		err := rows.Scan(&request.UID, &request.Error, &request.Time)
		if err != nil {
			log.Println(err.Error())

			return result
		}

		result = append(result, request)
	}

	return result
}

func QueryDetailedLog(db *sql.DB, connection, requestUID string) *sql.Row {
	var query string
	if connection == MySQL || connection == SQLite {
		query = "SELECT `uid`, `error`, `time` FROM `logs` WHERE `uid` = ?;"
	} else if connection == PostgreSQL {
		query = `SELECT "uid", "error", "time" FROM "logs" WHERE "uid" = ?;`
	}

	row := db.QueryRow(query, requestUID)

	return row
}

func QuerySearchLogs(db *sql.DB, connection, searchWildcard string, offset int) *sql.Rows {
	var query string
	if connection == MySQL {
		query = "SELECT `uid`, CASE WHEN LENGTH(`error`) > 80 THEN CONCAT(SUBSTRING(`error`, 1, 80), '...') " +
			"ELSE `error` END AS `error`, `time` FROM `logs` WHERE `application` = ? AND " +
			"(`uid` LIKE ? OR `application` LIKE ? " +
			"OR `error` LIKE ? OR `time` LIKE ?) " +
			"ORDER BY `time` DESC LIMIT ? OFFSET ?;"
	} else if connection == PostgreSQL {
		query = `SELECT "uid", CASE WHEN LENGTH("error") > 80 THEN CONCAT(SUBSTRING("error", 1, 80), '...')
			ELSE "error" END AS "error", "time" FROM "logs" WHERE "application" = ? AND
			("uid" LIKE ? OR "application" LIKE ?
			OR "error" LIKE ? OR "time" LIKE ?)
			ORDER BY "time" DESC LIMIT ? OFFSET ?;`
	} else if connection == SQLite {
		query = "SELECT `uid`, CASE WHEN LENGTH(`error`) > 80 THEN SUBSTR(`error`, 1, 80) || '...' " +
			"ELSE `error` END AS `error`, `time` FROM `logs` WHERE `application` = ? AND " +
			"(`uid` LIKE ? OR `application` LIKE ? " +
			"OR `error` LIKE ? OR `time` LIKE ?) " +
			"ORDER BY `time` DESC LIMIT ? OFFSET ?;"
	}

	rows, err := db.Query(
		query,
		utils.Config.ApplicationID,
		searchWildcard, searchWildcard, searchWildcard, searchWildcard,
		utils.Config.GoScopeEntriesPerPage,
		offset,
	)

	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return rows
}

func QueryGetLogs(db *sql.DB, connection string, offset int) *sql.Rows {
	var query string

	if connection == MySQL {
		query = "SELECT `uid`, CASE WHEN LENGTH(`error`) > 80 THEN CONCAT(SUBSTRING(`error`, 1, 80), '...') " +
			"ELSE `error` END AS `error`, `time` FROM `logs` WHERE `application` = ? " +
			"ORDER BY `time` DESC LIMIT ? OFFSET ?;"
	} else if connection == PostgreSQL {
		query = `SELECT "uid", CASE WHEN LENGTH("error") > 80 THEN CONCAT(SUBSTRING("error", 1, 80), '...') 
		ELSE "error" END AS "error", "time" FROM "logs" WHERE "application" = ?
		ORDER BY "time" DESC LIMIT ? OFFSET ?;`
	} else if connection == SQLite {
		query = "SELECT `uid`, CASE WHEN LENGTH(`error`) > 80 THEN SUBSTR(`error`, 1, 80) || '...' " +
			"ELSE `error` END AS `error`, `time` FROM `logs` WHERE `application` = ? " +
			"ORDER BY `time` DESC LIMIT ? OFFSET ?;"
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

func WriteLogs(message string) {
	fmt.Printf("%v", message)

	uid, _ := uuid.NewV4()
	query := "INSERT INTO logs (uid, application, error, time) VALUES " +
		"(?, ?, ?, ?)"

	_, err := utils.DB.Exec(
		query,
		uid.String(),
		utils.Config.ApplicationID,
		message,
		time.Now().Unix(),
	)
	if err != nil {
		log.Println(err.Error())
	}
}
