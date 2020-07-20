// Copyright © 2020 Pro Warehouse B.V.
// All Rights Reserved
package goscope

import (
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
	dashboardView, _ := Asset("static/html/dashboard.html")
	commonHeader, _ := Asset("static/html/head.html")
	highlightStyles, _ := Asset("static/css/highlight.css")
	watcherStyles, _ := Asset("static/css/watcher.css")
	footer, _ := Asset("static/html/footer.html")
	utilScripts, _ := Asset("static/js/utils.js")
	dashboardScripts, _ := Asset("static/js/dashboard.js")
	header := ReplaceVariablesInTemplate(string(commonHeader), map[string]string{"APPLICATION_NAME": os.Getenv("APPLICATION_NAME")})
	variables := map[string]string{
		"APPLICATION_NAME":  os.Getenv("APPLICATION_NAME"),
		"COMMON_HEADER":     MinifyHtml(header),
		"HIGHLIGHT_STYLES":  MinifyCss(string(highlightStyles)),
		"WATCHER_STYLES":    MinifyCss(string(watcherStyles)),
		"ENTRIES_PER_PAGE":  os.Getenv("GOSCOPE_ENTRIES_PER_PAGE"),
		"NAVBAR":            NavbarComponent("REQUESTS"),
		"FOOTER":            MinifyHtml(string(footer)),
		"UTIL_SCRIPTS":      MinifyJs(string(utilScripts)),
		"DASHBOARD_SCRIPTS": MinifyJs(string(dashboardScripts)),
	}
	ShowGoScopePage(c, MinifyHtml(string(dashboardView)), variables)
}

func LogDashboard(c *gin.Context) {
	logsView, _ := Asset("static/html/logs.html")
	commonHeader, _ := Asset("static/html/head.html")
	highlightStyles, _ := Asset("static/css/highlight.css")
	watcherStyles, _ := Asset("static/css/watcher.css")
	footer, _ := Asset("static/html/footer.html")
	utilScripts, _ := Asset("static/js/utils.js")
	logScripts, _ := Asset("static/js/logs.js")
	header := ReplaceVariablesInTemplate(string(commonHeader), map[string]string{"APPLICATION_NAME": os.Getenv("APPLICATION_NAME")})
	variables := map[string]string{
		"APPLICATION_NAME": os.Getenv("APPLICATION_NAME"),
		"COMMON_HEADER":    MinifyHtml(header),
		"HIGHLIGHT_STYLES": MinifyCss(string(highlightStyles)),
		"WATCHER_STYLES":   MinifyCss(string(watcherStyles)),
		"ENTRIES_PER_PAGE": os.Getenv("GOSCOPE_ENTRIES_PER_PAGE"),
		"NAVBAR":           NavbarComponent("LOGS"),
		"FOOTER":           MinifyHtml(string(footer)),
		"UTIL_SCRIPTS":     MinifyJs(string(utilScripts)),
		"LOG_SCRIPTS":      MinifyJs(string(logScripts)),
	}
	ShowGoScopePage(c, MinifyHtml(string(logsView)), variables)
}

func ShowRequest(c *gin.Context) {
	var request RecordByUri
	err := c.ShouldBindUri(&request)
	if err != nil {
		log.Println(err.Error())
	}
	requestDetails := GetDetailedRequest(request.Uid)
	responseDetails := GetDetailedResponse(request.Uid)
	requestView, _ := Asset("static/html/request.html")
	commonHeader, _ := Asset("static/html/head.html")
	highlightStyles, _ := Asset("static/css/highlight.css")
	watcherStyles, _ := Asset("static/css/watcher.css")
	requestScripts, _ := Asset("static/js/request.js")
	header := ReplaceVariablesInTemplate(string(commonHeader), map[string]string{"APPLICATION_NAME": os.Getenv("APPLICATION_NAME")})
	variables := map[string]string{
		"APPLICATION_NAME":   os.Getenv("APPLICATION_NAME"),
		"COMMON_HEADER":      MinifyHtml(header),
		"HIGHLIGHT_STYLES":   MinifyCss(string(highlightStyles)),
		"WATCHER_STYLES":     MinifyCss(string(watcherStyles)),
		"REQUEST_SCRIPTS":    MinifyJs(string(requestScripts)),
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

func ShowLog(c *gin.Context) {
	var request RecordByUri
	err := c.ShouldBindUri(&request)
	if err != nil {
		log.Println(err.Error())
	}
	logDetails := GetDetailedLog(request.Uid)
	logView, _ := Asset("static/html/log.html")
	commonHeader, _ := Asset("static/html/head.html")
	highlightStyles, _ := Asset("static/css/highlight.css")
	watcherStyles, _ := Asset("static/css/watcher.css")
	requestScripts, _ := Asset("static/js/log.js")
	header := ReplaceVariablesInTemplate(string(commonHeader), map[string]string{"APPLICATION_NAME": os.Getenv("APPLICATION_NAME")})
	variables := map[string]string{
		"APPLICATION_NAME": os.Getenv("APPLICATION_NAME"),
		"COMMON_HEADER":    MinifyHtml(header),
		"HIGHLIGHT_STYLES": MinifyCss(string(highlightStyles)),
		"WATCHER_STYLES":   MinifyCss(string(watcherStyles)),
		"LOG_SCRIPTS":      MinifyJs(string(requestScripts)),
		"TIME":             UnixTimeToHuman(logDetails.Time),
		"MESSAGE":          logDetails.Error,
	}
	ShowGoScopePage(c, MinifyHtml(string(logView)), variables)
}
