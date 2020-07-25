package goscope

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetDetailedLog(requestUID string) ExceptionRecord {
	db := GetDB()
	defer db.Close()

	query := "SELECT uid, error, time FROM logs WHERE uid = ?;"
	row := db.QueryRow(query, requestUID)

	var request ExceptionRecord

	err := row.Scan(&request.UID, &request.Error, &request.Time)
	if err != nil {
		log.Println(err.Error())
		return request
	}

	return request
}

func SearchLogs(searchString string, offset int) []ExceptionRecord {
	var result []ExceptionRecord

	db := GetDB()

	defer db.Close()

	query := "SELECT uid, CASE WHEN LENGTH(error) > 80 THEN CONCAT(SUBSTRING(error, 1, 80), '...') " +
		"ELSE error END AS error, time FROM logs WHERE application = ? AND " +
		"(uid LIKE ? OR application LIKE ? " +
		"OR error LIKE ? OR time LIKE ?) " +
		"ORDER BY time DESC LIMIT ? OFFSET ?;"

	searchWildcard := fmt.Sprintf("%%%s%%", searchString)
	rows, err := db.Query(
		query,
		os.Getenv("APPLICATION_ID"),
		searchWildcard, searchWildcard, searchWildcard, searchWildcard,
		os.Getenv("GOSCOPE_ENTRIES_PER_PAGE"),
		offset,
	)

	if err != nil {
		log.Println(err.Error())
		return result
	}

	defer rows.Close()

	for rows.Next() {
		var request ExceptionRecord

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
func GetLogs(c *gin.Context) {
	offsetQuery := c.DefaultQuery("offset", "0")
	offset, _ := strconv.ParseInt(offsetQuery, 10, 32)
	db := GetDB()

	defer db.Close()

	query := "SELECT uid, CASE WHEN LENGTH(error) > 80 THEN CONCAT(SUBSTRING(error, 1, 80), '...') " +
		"ELSE error END AS error, time FROM logs WHERE application = ? " +
		"ORDER BY time DESC LIMIT ? OFFSET ?;"
	rows, err := db.Query(
		query,
		os.Getenv("APPLICATION_ID"),
		os.Getenv("GOSCOPE_ENTRIES_PER_PAGE"),
		offset,
	)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())

		return
	}

	defer rows.Close()

	var result []ExceptionRecord

	for rows.Next() {
		var request ExceptionRecord

		err := rows.Scan(&request.UID, &request.Error, &request.Time)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, err.Error())

			return
		}

		result = append(result, request)
	}

	c.JSON(http.StatusOK, result)
}
