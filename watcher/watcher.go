package watcher

import (
	"bitbucket.org/prowarehouse-nl/gohttpwatcher/watcher_templates"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func Dashboard(c *gin.Context) {
	/*variables := map[string]string{
		"APPLICATION_NAME":os.Getenv("APPLICATION_NAME"),
	}*/
	r := strings.NewReader(watcher_templates.IndexTemplate)
	c.DataFromReader(http.StatusOK, r.Size(), "text/html", r, nil)
}

func ShowRequest(c *gin.Context) {
	var request RecordByUri
	_ = c.ShouldBindUri(&request)
	requestDetails := GetDetailedRequest(request.Uid)
	variables := gin.H{
		"APPLICATION_NAME": os.Getenv("APPLICATION_NAME"),
		"BODY":             formatJson(requestDetails.Body),
		"CLIENT_IP":        requestDetails.ClientIp,
		"HEADERS":          formatJson(requestDetails.Headers),
		"HOST":             requestDetails.Host,
		"METHOD":           requestDetails.Method,
		"PATH":             requestDetails.Path,
		"REFERRER":         requestDetails.Referrer,
		"TIME":             UnixTimeToAmsterdam(requestDetails.Time),
		"UID":              requestDetails.Uid,
		"URL":              requestDetails.Url,
		"USER_AGENT":       requestDetails.UserAgent,
	}
	fmt.Printf("%+v", variables)
	r := strings.NewReader(watcher_templates.RequestTemplate)
	c.DataFromReader(http.StatusOK, r.Size(), "text/html", r, nil)
}

func ShowResponse(c *gin.Context) {
	var request RecordByUri
	_ = c.ShouldBindUri(&request)
	responseDetails := GetDetailedResponse(request.Uid)
	variables := gin.H{
		"APPLICATION_NAME": os.Getenv("APPLICATION_NAME"),
		"BODY":             formatJson(responseDetails.Body),
		"CLIENT_IP":        responseDetails.ClientIp,
		"HEADERS":          formatJson(responseDetails.Headers),
		"PATH":             responseDetails.Path,
		"SIZE":             strconv.Itoa(responseDetails.Size),
		"STATUS":           responseDetails.Status,
		"TIME":             UnixTimeToAmsterdam(responseDetails.Time),
		"UID":              responseDetails.Uid,
	}
	fmt.Printf("%+v", variables)
	r := strings.NewReader(watcher_templates.ResponseTemplate)
	c.DataFromReader(http.StatusOK, r.Size(), "text/html", r, nil)
}
