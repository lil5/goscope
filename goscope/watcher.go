// Copyright © 2020 Pro Warehouse B.V.
// All Rights Reserved
package goscope

import (
	"bitbucket.org/prowarehouse-nl/goscope/goscope_templates"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func ShowGoScopePage(c *gin.Context, rawTemplate string, variables map[string]string) {
	cleanTemplate := ReplaceVariablesInTemplate(rawTemplate, variables)
	reader := strings.NewReader(cleanTemplate)
	c.DataFromReader(http.StatusOK, reader.Size(), "text/html", reader, nil)
}

func Dashboard(c *gin.Context) {
	variables := map[string]string{
		"APPLICATION_NAME": os.Getenv("APPLICATION_NAME"),
	}
	ShowGoScopePage(c, goscope_templates.DashboardView(), variables)
}

func LogDashboard(c *gin.Context) {
	variables := map[string]string{
		"APPLICATION_NAME": os.Getenv("APPLICATION_NAME"),
	}
	ShowGoScopePage(c, goscope_templates.LogsView(), variables)
}

func ShowRequest(c *gin.Context) {
	var request RecordByUri
	err := c.ShouldBindUri(&request)
	if err != nil {
		log.Println(err.Error())
	}
	requestDetails := GetDetailedRequest(request.Uid)
	responseDetails := GetDetailedResponse(request.Uid)
	variables := map[string]string{
		"APPLICATION_NAME":   os.Getenv("APPLICATION_NAME"),
		"REQUEST_BODY":       prettifyJson(requestDetails.Body),
		"REQUEST_CLIENT_IP":  requestDetails.ClientIp,
		"REQUEST_HEADERS":    prettifyJson(requestDetails.Headers),
		"REQUEST_HOST":       requestDetails.Host,
		"REQUEST_METHOD":     requestDetails.Method,
		"REQUEST_PATH":       requestDetails.Path,
		"REQUEST_REFERRER":   requestDetails.Referrer,
		"REQUEST_TIME":       UnixTimeToAmsterdam(requestDetails.Time),
		"REQUEST_UID":        requestDetails.Uid,
		"REQUEST_URL":        requestDetails.Url,
		"REQUEST_USER_AGENT": requestDetails.UserAgent,
		"RESPONSE_BODY":      prettifyJson(responseDetails.Body),
		"RESPONSE_CLIENT_IP": responseDetails.ClientIp,
		"RESPONSE_HEADERS":   prettifyJson(responseDetails.Headers),
		"RESPONSE_PATH":      responseDetails.Path,
		"RESPONSE_SIZE":      strconv.Itoa(responseDetails.Size),
		"RESPONSE_STATUS":    responseDetails.Status,
		"RESPONSE_TIME":      UnixTimeToAmsterdam(responseDetails.Time),
		"RESPONSE_UID":       responseDetails.Uid,
	}
	ShowGoScopePage(c, goscope_templates.RequestTemplate, variables)
}
