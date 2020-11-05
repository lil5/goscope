// License: MIT
// Authors:
// 		- Josep Jesus Bigorra Algaba (@averageflow)

package goscope

import (
	"fmt"
	"log"

	"github.com/averageflow/goscope/utils"

	"github.com/averageflow/goscope/database"
)

func GetDetailedLog(requestUID string) ExceptionRecord {
	row := database.GetDetailedLog(
		utils.DB,
		utils.Config.GoScopeDatabaseType,
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

	searchWildcard := fmt.Sprintf("%%%s%%", searchString)
	rows := database.SearchLogs(utils.DB, utils.Config.GoScopeDatabaseType, searchWildcard, offset)

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

	rows := database.GetLogs(utils.DB, utils.Config.GoScopeDatabaseType, offset)
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
