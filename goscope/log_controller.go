package goscope

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func ShowDashboard(c *gin.Context, mode int) {
	// Markup
	var baseTemplate string

	if mode == LogDashboardMode {
		logsView, _ := Asset("../static/html/log_dashboard.html")
		baseTemplate = string(logsView)
	} else {
		requestView, _ := Asset("../static/html/request_dashboard.html")
		baseTemplate = string(requestView)
	}

	commonHeader, _ := Asset("../static/html/common_head.html")
	footer, _ := Asset("../static/html/common_footer.html")
	commonNavbar, _ := Asset("../static/html/common_navbar.html")
	navbarVariables := map[string]string{"APPLICATION_NAME": os.Getenv("APPLICATION_NAME")}
	headerVariables := map[string]string{"APPLICATION_NAME": os.Getenv("APPLICATION_NAME")}
	navbar := ReplaceVariablesInTemplate(string(commonNavbar), navbarVariables)
	header := ReplaceVariablesInTemplate(string(commonHeader), headerVariables)
	// Styles
	highlightStyles, _ := Asset("../static/css/highlight.css")
	goscopeStyles, _ := Asset("../static/css/goscope.css")
	// Scripts
	utilScripts, _ := Asset("../static/js/utils.js")
	abstractDashboard, _ := Asset("../static/js/abstractDashboard.js")

	var dashboardScript string

	if mode == LogDashboardMode {
		logsDashboard, _ := Asset("../static/js/logsDashboard.js")
		dashboardScript = string(logsDashboard)
	} else {
		requestDashboard, _ := Asset("../static/js/requestDashboard.js")
		dashboardScript = string(requestDashboard)
	}

	variables := map[string]string{
		"APPLICATION_NAME":   os.Getenv("APPLICATION_NAME"),
		"COMMON_HEADER":      MinifyHTML(header),
		"HIGHLIGHT_STYLES":   MinifyCSS(string(highlightStyles)),
		"GOSCOPE_STYLES":     MinifyCSS(string(goscopeStyles)),
		"ENTRIES_PER_PAGE":   os.Getenv("GOSCOPE_ENTRIES_PER_PAGE"),
		"COMMON_NAVBAR":      MinifyHTML(navbar),
		"COMMON_FOOTER":      MinifyHTML(string(footer)),
		"UTIL_SCRIPTS":       MinifyJs(string(utilScripts)),
		"ABSTRACT_DASHBOARD": MinifyJs(string(abstractDashboard)),
	}
	if mode == LogDashboardMode {
		variables["LOG_DASHBOARD"] = dashboardScript
	} else {
		variables["REQUEST_DASHBOARD"] = dashboardScript
	}

	ShowGoScopePage(c, MinifyHTML(baseTemplate), variables)
}

func LogDashboard(c *gin.Context) {
	ShowDashboard(c, LogDashboardMode)
}

func ShowLog(c *gin.Context) {
	var request RecordByURI

	err := c.ShouldBindUri(&request)
	if err != nil {
		log.Println(err.Error())
	}

	logDetails := GetDetailedLog(request.UID)
	// Markup
	logView, _ := Asset("../static/html/single_log.html")
	commonHeader, _ := Asset("../static/html/common_head.html")
	headerVariables := map[string]string{"APPLICATION_NAME": os.Getenv("APPLICATION_NAME")}
	header := ReplaceVariablesInTemplate(string(commonHeader), headerVariables)
	// Styles
	highlightStyles, _ := Asset("../static/css/highlight.css")
	goscopeStyles, _ := Asset("../static/css/goscope.css")
	// Scripts
	singleLog, _ := Asset("../static/js/singleLog.js")

	variables := map[string]string{
		"APPLICATION_NAME": os.Getenv("APPLICATION_NAME"),
		"COMMON_HEADER":    MinifyHTML(header),
		"HIGHLIGHT_STYLES": MinifyCSS(string(highlightStyles)),
		"GOSCOPE_STYLES":   MinifyCSS(string(goscopeStyles)),
		"SINGLE_LOG":       MinifyJs(string(singleLog)),
		"TIME":             UnixTimeToHuman(logDetails.Time),
		"MESSAGE":          logDetails.Error,
	}

	ShowGoScopePage(c, MinifyHTML(string(logView)), variables)
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
