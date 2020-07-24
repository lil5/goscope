package goscope

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
	"os"
	"strconv"
)

func Dashboard(c *gin.Context) {
	// Markup
	dashboardView, _ := Asset("../static/html/request_dashboard.html")
	commonHeader, _ := Asset("../static/html/common_head.html")
	footer, _ := Asset("../static/html/common_footer.html")
	commonNavbar, _ := Asset("../static/html/common_navbar.html")
	navbar := ReplaceVariablesInTemplate(string(commonNavbar), map[string]string{"APPLICATION_NAME": os.Getenv("APPLICATION_NAME")})
	header := ReplaceVariablesInTemplate(string(commonHeader), map[string]string{"APPLICATION_NAME": os.Getenv("APPLICATION_NAME")})
	// Styles
	highlightStyles, _ := Asset("../static/css/highlight.css")
	goscopeStyles, _ := Asset("../static/css/goscope.css")
	// Scripts
	utilScripts, _ := Asset("../static/js/utils.js")
	abstractDashboard, _ := Asset("../static/js/abstractDashboard.js")
	requestDashboard, _ := Asset("../static/js/requestDashboard.js")
	variables := map[string]string{
		"APPLICATION_NAME":   os.Getenv("APPLICATION_NAME"),
		"COMMON_HEADER":      MinifyHtml(header),
		"HIGHLIGHT_STYLES":   MinifyCss(string(highlightStyles)),
		"GOSCOPE_STYLES":     MinifyCss(string(goscopeStyles)),
		"ENTRIES_PER_PAGE":   os.Getenv("GOSCOPE_ENTRIES_PER_PAGE"),
		"COMMON_NAVBAR":      MinifyHtml(navbar),
		"COMMON_FOOTER":      MinifyHtml(string(footer)),
		"UTIL_SCRIPTS":       MinifyJs(string(utilScripts)),
		"ABSTRACT_DASHBOARD": MinifyJs(string(abstractDashboard)),
		"REQUEST_DASHBOARD":  MinifyJs(string(requestDashboard)),
	}
	ShowGoScopePage(c, MinifyHtml(string(dashboardView)), variables)
}

func ShowRequest(c *gin.Context) {
	var request RecordByUri
	err := c.ShouldBindUri(&request)

	if err != nil {
		log.Println(err.Error())
	}

	requestDetails := GetDetailedRequest(request.Uid)
	responseDetails := GetDetailedResponse(request.Uid)
	// Markup
	requestView, _ := Asset("../static/html/single_request.html")
	commonHeader, _ := Asset("../static/html/common_head.html")
	header := ReplaceVariablesInTemplate(string(commonHeader), map[string]string{"APPLICATION_NAME": os.Getenv("APPLICATION_NAME")})
	// Styles
	highlightStyles, _ := Asset("../static/css/highlight.css")
	goscopeStyles, _ := Asset("../static/css/goscope.css")
	// Scripts
	singleRequest, _ := Asset("../static/js/singleRequest.js")

	variables := map[string]string{
		"APPLICATION_NAME":   os.Getenv("APPLICATION_NAME"),
		"COMMON_HEADER":      MinifyHtml(header),
		"HIGHLIGHT_STYLES":   MinifyCss(string(highlightStyles)),
		"GOSCOPE_STYLES":     MinifyCss(string(goscopeStyles)),
		"SINGLE_REQUEST":     MinifyJs(string(singleRequest)),
		"REQUEST_BODY":       prettifyJson(requestDetails.Body),
		"REQUEST_CLIENT_IP":  requestDetails.ClientIp,
		"REQUEST_HEADERS":    prettifyJson(requestDetails.Headers),
		"REQUEST_HOST":       requestDetails.Host,
		"REQUEST_METHOD":     requestDetails.Method,
		"REQUEST_PATH":       requestDetails.Path,
		"REQUEST_REFERRER":   requestDetails.Referrer,
		"REQUEST_TIME":       UnixTimeToHuman(requestDetails.Time),
		"REQUEST_UID":        requestDetails.Uid,
		"REQUEST_URL":        requestDetails.Url,
		"REQUEST_USER_AGENT": requestDetails.UserAgent,
		"RESPONSE_BODY":      prettifyJson(responseDetails.Body),
		"RESPONSE_CLIENT_IP": responseDetails.ClientIp,
		"RESPONSE_HEADERS":   prettifyJson(responseDetails.Headers),
		"RESPONSE_PATH":      responseDetails.Path,
		"RESPONSE_SIZE":      strconv.Itoa(responseDetails.Size),
		"RESPONSE_STATUS":    responseDetails.Status,
		"RESPONSE_TIME":      UnixTimeToHuman(responseDetails.Time),
		"RESPONSE_UID":       responseDetails.Uid,
	}
	ShowGoScopePage(c, MinifyHtml(string(requestView)), variables)
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
