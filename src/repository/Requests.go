package repository

import (
	"log"

	"github.com/averageflow/goscope/src/types"
	"github.com/averageflow/goscope/src/utils"
)

// FetchDetailedRequest fetches all details from a request via its UUID.
func FetchDetailedRequest(requestUID string) types.DetailedRequest {
	var body string

	var headers string

	var result types.DetailedRequest

	row := QueryDetailedRequest(utils.DB, requestUID)

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

// FetchDetailedResponse fetches all details of a response via its UUID.
func FetchDetailedResponse(responseUUID string) types.DetailedResponse {
	var body string

	var headers string

	var result types.DetailedResponse

	row := QueryDetailedResponse(utils.DB, responseUUID)

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

// FetchRequestList fetches a list of summarized requests.
func FetchRequestList(offset int) []types.SummarizedRequest {
	var result []types.SummarizedRequest

	rows, err := QueryGetRequests(utils.DB, offset)
	if err != nil {
		log.Println(err.Error())

		return result
	}

	if rows.Err() != nil {
		log.Println(rows.Err().Error())

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

// FetchSearchRequests fetches a list of summarized requests that match the input parameters of search.
func FetchSearchRequests(search string, filter *types.RequestFilter, offset int) []types.SummarizedRequest {
	var result []types.SummarizedRequest

	rows, err := QuerySearchRequests(utils.DB, utils.Config.GoScopeDatabaseType, search, filter, offset)
	if err != nil {
		log.Println(err.Error())
		return result
	}

	if rows.Err() != nil {
		log.Println(rows.Err().Error())

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
		}

		result = append(result, request)
	}

	return result
}
