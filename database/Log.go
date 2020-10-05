// License: MIT
// Authors:
// 		- Josep Jesus Bigorra Algaba (@averageflow)
package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	uuid "github.com/nu7hatch/gouuid"

	"github.com/averageflow/goscope/utils"
)

func GetDetailedLog(db *sql.DB, connection, requestUID string) *sql.Row {
	var query string
	if connection == MySQL || connection == SQLite {
		query = "SELECT `uid`, `error`, `time` FROM `logs` WHERE `uid` = ?;"
	} else if connection == PostgreSQL {
		query = `SELECT "uid", "error", "time" FROM "logs" WHERE "uid" = ?;`
	}

	row := db.QueryRow(query, requestUID)

	return row
}

func SearchLogs(db *sql.DB, connection, searchWildcard string, offset int) *sql.Rows {
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

func GetLogs(db *sql.DB, connection string, offset int) *sql.Rows {
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
