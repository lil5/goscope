package watcher

import (
	"bitbucket.org/prowarehouse-nl/gohttpwatcher/watcher_templates"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func Dashboard(c *gin.Context) {
	variables := map[string]string{
		"APPLICATION_NAME": os.Getenv("APPLICATION_NAME"),
	}
	rawTemplate := watcher_templates.IndexTemplate
	cleanTemplate := ReplaceVariablesInTemplate(rawTemplate, variables)
	reader := strings.NewReader(cleanTemplate)
	c.DataFromReader(http.StatusOK, reader.Size(), "text/html", reader, nil)
}

func ShowRequest(c *gin.Context) {
	var request RecordByUri
	_ = c.ShouldBindUri(&request)
	requestDetails := GetDetailedRequest(request.Uid)
	responseDetails := GetDetailedResponse(request.Uid)
	variables := map[string]string{
		"APPLICATION_NAME":  os.Getenv("APPLICATION_NAME"),
		"REQUEST_BODY":              formatJson(requestDetails.Body),
		"REQUEST_CLIENT_IP":         requestDetails.ClientIp,
		"REQUEST_HEADERS":           formatJson(requestDetails.Headers),
		"REQUEST_HOST":              requestDetails.Host,
		"REQUEST_METHOD":            requestDetails.Method,
		"REQUEST_PATH":              requestDetails.Path,
		"REQUEST_REFERRER":          requestDetails.Referrer,
		"REQUEST_TIME":              UnixTimeToAmsterdam(requestDetails.Time),
		"REQUEST_UID":               requestDetails.Uid,
		"REQUEST_URL":               requestDetails.Url,
		"REQUEST_USER_AGENT":        requestDetails.UserAgent,
		"RESPONSE_BODY":             formatJson(responseDetails.Body),
		"RESPONSE_CLIENT_IP":        responseDetails.ClientIp,
		"RESPONSE_HEADERS":          formatJson(responseDetails.Headers),
		"RESPONSE_PATH":             responseDetails.Path,
		"RESPONSE_SIZE":             strconv.Itoa(responseDetails.Size),
		"RESPONSE_STATUS":           responseDetails.Status,
		"RESPONSE_TIME":             UnixTimeToAmsterdam(responseDetails.Time),
		"RESPONSE_UID":              responseDetails.Uid,
	}
	rawTemplate := watcher_templates.RequestTemplate
	cleanTemplate := ReplaceVariablesInTemplate(rawTemplate, variables)
	reader := strings.NewReader(cleanTemplate)
	c.DataFromReader(http.StatusOK, reader.Size(), "text/html", reader, nil)
}
