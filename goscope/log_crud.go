package goscope

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/averageflow/goscope/repository"

	"github.com/gin-gonic/gin"
)

func GetDetailedLog(requestUID string) ExceptionRecord {
	row := repository.GetDetailedLog(
		os.Getenv("GOSCOPE_DATABASE_TYPE"),
		requestUID,
	)

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

	searchWildcard := fmt.Sprintf("%%%s%%", searchString)
	rows := repository.SearchLogs(os.Getenv("GOSCOPE_DATABASE_TYPE"), searchWildcard, offset)

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

	rows := repository.GetLogs(os.Getenv("GOSCOPE_DATABASE_TYPE"), int(offset))
	if rows == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error querying DB"})
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
