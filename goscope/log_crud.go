package goscope

import (
	"fmt"
	"log"
	"os"

	"github.com/averageflow/goscope/repository"
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
func GetLogs(offset int) []ExceptionRecord {
	var result []ExceptionRecord

	rows := repository.GetLogs(os.Getenv("GOSCOPE_DATABASE_TYPE"), offset)
	if rows == nil {
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
