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
		"ELSE error END AS error, time FROM logs WHERE application = '%[1]v' AND " +
		"(uid LIKE '%%%[3]v%%' OR application LIKE '%%%[3]v%%' " +
		"OR error LIKE '%%%[3]v%%' OR time LIKE '%%%[3]v%%') " +
		"ORDER BY time DESC LIMIT %[2]v OFFSET %[4]v;"
	preparedQuery := fmt.Sprintf(
		query,
		os.Getenv("APPLICATION_ID"),
		os.Getenv("GOSCOPE_ENTRIES_PER_PAGE"),
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
