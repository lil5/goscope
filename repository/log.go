package repository

import (
	"database/sql"
	"log"
	"os"
)

func GetDetailedLog(connection, requestUID string) *sql.Row {
	db := GetDB()
	defer db.Close()

	var query string
	if connection == MySQL {
		query = "SELECT `uid`, `error`, `time` FROM `logs` WHERE `uid` = ?;"
	} else if connection == PostgreSQL {
		query = `SELECT "uid", "error", "time" FROM "logs" WHERE "uid" = ?;`
	}

	row := db.QueryRow(query, requestUID)

	return row
}

func SearchLogs(connection, searchWildcard string, offset int) *sql.Rows {
	db := GetDB()
	defer db.Close()

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
	}

	rows, err := db.Query(
		query,
		os.Getenv("APPLICATION_ID"),
		searchWildcard, searchWildcard, searchWildcard, searchWildcard,
		os.Getenv("GOSCOPE_ENTRIES_PER_PAGE"),
		offset,
	)

	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return rows
}

func GetLogs(connection string, offset int) *sql.Rows {
	db := GetDB()
	defer db.Close()

	var query string
	if connection == MySQL {
		query = "SELECT `uid`, CASE WHEN LENGTH(`error`) > 80 THEN CONCAT(SUBSTRING(`error`, 1, 80), '...') " +
			"ELSE `error` END AS `error`, `time` FROM `logs` WHERE `application` = ? " +
			"ORDER BY `time` DESC LIMIT ? OFFSET ?;"
	} else if connection == PostgreSQL {
		query = `SELECT "uid", CASE WHEN LENGTH("error") > 80 THEN CONCAT(SUBSTRING("error", 1, 80), '...') 
		ELSE "error" END AS "error", "time" FROM "logs" WHERE "application" = ?
		ORDER BY "time" DESC LIMIT ? OFFSET ?;`
	}

	rows, err := db.Query(
		query,
		os.Getenv("APPLICATION_ID"),
		os.Getenv("GOSCOPE_ENTRIES_PER_PAGE"),
		offset,
	)

	if err != nil {
		log.Println(err.Error())

		return nil
	}

	return rows
}
