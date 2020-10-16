// License: MIT
// Authors:
// 		- Josep Jesus Bigorra Algaba (@averageflow)
package goscope

import (
	"log"

	"github.com/averageflow/goscope/utils"

	"github.com/averageflow/goscope/database"
)

// Get all details from a request via its UID.
func GetDetailedRequest(requestUID string) DetailedRequest {
	var body string

	var headers string

	var result DetailedRequest

	row := database.GetDetailedRequest(utils.DB, utils.Config.GoScopeDatabaseType, requestUID)

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

	result.Body = prettifyJSON(body)
	result.Headers = prettifyJSON(headers)

	return result
}

func GetDetailedResponse(requestUID string) DetailedResponse {
	var body string

	var headers string

	var result DetailedResponse

	row := database.GetDetailedResponse(utils.DB, utils.Config.GoScopeDatabaseType, requestUID)

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

	result.Body = prettifyJSON(body)
	result.Headers = prettifyJSON(headers)

	return result
}

func GetRequests(offset int) []SummarizedRequest {
	var result []SummarizedRequest

	rows := database.GetRequests(utils.DB, utils.Config.GoScopeDatabaseType, offset)

	if rows == nil {
		return result
	}
	defer rows.Close()

	for rows.Next() {
		var request SummarizedRequest

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

func SearchRequests(search string, filter *database.RequestFilter, offset int) []SummarizedRequest {
	var result []SummarizedRequest

	rows, err := database.SearchRequests(utils.DB, utils.Config.GoScopeDatabaseType, search, filter, offset)

	if err != nil {
		log.Println(err.Error())
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var request SummarizedRequest

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
