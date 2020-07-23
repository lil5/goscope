package goscope

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
	"os"
	"strconv"
)

func LogDashboard(c *gin.Context) {
	// Markup
	logsView, _ := Asset("static/html/log_dashboard.html")
	commonHeader, _ := Asset("static/html/common_head.html")
	footer, _ := Asset("static/html/common_footer.html")
	commonNavbar, _ := Asset("static/html/common_navbar.html")
	navbar := ReplaceVariablesInTemplate(string(commonNavbar), map[string]string{"APPLICATION_NAME": os.Getenv("APPLICATION_NAME")})
	header := ReplaceVariablesInTemplate(string(commonHeader), map[string]string{"APPLICATION_NAME": os.Getenv("APPLICATION_NAME")})
	// Styles
	highlightStyles, _ := Asset("static/css/highlight.css")
	goscopeStyles, _ := Asset("static/css/goscope.css")
	// Scripts
	utilScripts, _ := Asset("static/js/utils.js")
	abstractDashboard, _ := Asset("static/js/abstractDashboard.js")
	logsDashboard, _ := Asset("static/js/logsDashboard.js")

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
		"LOG_DASHBOARD":      MinifyJs(string(logsDashboard)),
	}
	ShowGoScopePage(c, MinifyHtml(string(logsView)), variables)
}

func ShowLog(c *gin.Context) {
	var request RecordByUri
	err := c.ShouldBindUri(&request)
	if err != nil {
		log.Println(err.Error())
	}
	logDetails := GetDetailedLog(request.Uid)
	// Markup
	logView, _ := Asset("static/html/single_log.html")
	commonHeader, _ := Asset("static/html/common_head.html")
	header := ReplaceVariablesInTemplate(string(commonHeader), map[string]string{"APPLICATION_NAME": os.Getenv("APPLICATION_NAME")})
	// Styles
	highlightStyles, _ := Asset("static/css/highlight.css")
	goscopeStyles, _ := Asset("static/css/goscope.css")
	// Scripts
	singleLog, _ := Asset("static/js/singleLog.js")

	variables := map[string]string{
		"APPLICATION_NAME": os.Getenv("APPLICATION_NAME"),
		"COMMON_HEADER":    MinifyHtml(header),
		"HIGHLIGHT_STYLES": MinifyCss(string(highlightStyles)),
		"GOSCOPE_STYLES":   MinifyCss(string(goscopeStyles)),
		"SINGLE_LOG":       MinifyJs(string(singleLog)),
		"TIME":             UnixTimeToHuman(logDetails.Time),
		"MESSAGE":          logDetails.Error,
	}
	ShowGoScopePage(c, MinifyHtml(string(logView)), variables)
}

func SearchLog(c *gin.Context) {
	var request SearchRequestPayload
	err := c.ShouldBindBodyWith(&request, binding.JSON)
	if err != nil {
		log.Println(err.Error())
	}
	offsetQuery := c.DefaultQuery("offset", "0")
	offset, _ := strconv.ParseInt(offsetQuery, 10, 32)
	result := SearchLogs(request.Query, int(offset))
	c.JSON(http.StatusOK, result)
}
