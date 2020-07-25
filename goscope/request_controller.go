package goscope

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// Show the dashboard of the HTTP request/responses.
func RequestDashboard(c *gin.Context) {
	ShowDashboard(c, RequestDashboardMode)
}

func ShowRequest(c *gin.Context) {
	var request RecordByURI

	err := c.ShouldBindUri(&request)
	if err != nil {
		log.Println(err.Error())
	}

	requestDetails := GetDetailedRequest(request.UID)
	responseDetails := GetDetailedResponse(request.UID)
	// Markup
	requestView, _ := Asset("../static/html/single_request.html")
	commonHeader, _ := Asset("../static/html/common_head.html")
	headerVariables := map[string]string{"APPLICATION_NAME": os.Getenv("APPLICATION_NAME")}
	header := ReplaceVariablesInTemplate(string(commonHeader), headerVariables)
	// Styles
	highlightStyles, _ := Asset("../static/css/highlight.css")
	goscopeStyles, _ := Asset("../static/css/goscope.css")
	// Scripts
	singleRequest, _ := Asset("../static/js/singleRequest.js")

	variables := map[string]string{
		"APPLICATION_NAME":   os.Getenv("APPLICATION_NAME"),
		"COMMON_HEADER":      MinifyHTML(header),
		"HIGHLIGHT_STYLES":   MinifyCSS(string(highlightStyles)),
		"GOSCOPE_STYLES":     MinifyCSS(string(goscopeStyles)),
		"SINGLE_REQUEST":     MinifyJs(string(singleRequest)),
		"REQUEST_BODY":       prettifyJSON(requestDetails.Body),
		"REQUEST_CLIENT_IP":  requestDetails.ClientIP,
		"REQUEST_HEADERS":    prettifyJSON(requestDetails.Headers),
		"REQUEST_HOST":       requestDetails.Host,
		"REQUEST_METHOD":     requestDetails.Method,
		"REQUEST_PATH":       requestDetails.Path,
		"REQUEST_REFERRER":   requestDetails.Referrer,
		"REQUEST_TIME":       UnixTimeToHuman(requestDetails.Time),
		"REQUEST_UID":        requestDetails.UID,
		"REQUEST_URL":        requestDetails.URL,
		"REQUEST_USER_AGENT": requestDetails.UserAgent,
		"RESPONSE_BODY":      prettifyJSON(responseDetails.Body),
		"RESPONSE_CLIENT_IP": responseDetails.ClientIP,
		"RESPONSE_HEADERS":   prettifyJSON(responseDetails.Headers),
		"RESPONSE_PATH":      responseDetails.Path,
		"RESPONSE_SIZE":      strconv.Itoa(responseDetails.Size),
		"RESPONSE_STATUS":    responseDetails.Status,
		"RESPONSE_TIME":      UnixTimeToHuman(responseDetails.Time),
		"RESPONSE_UID":       responseDetails.UID,
	}
	ShowGoScopePage(c, MinifyHTML(string(requestView)), variables)
}

func SearchRequest(c *gin.Context) {
	var request SearchRequestPayload
	err := c.ShouldBindBodyWith(&request, binding.JSON)

	if err != nil {
		log.Println(err.Error())
	}

	offsetQuery := c.DefaultQuery("offset", "0")
	offset, _ := strconv.ParseInt(offsetQuery, 10, 32)
	result := SearchRequests(request.Query, int(offset))
	c.JSON(http.StatusOK, result)
}
