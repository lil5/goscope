package goscope

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ShowDashboard(c *gin.Context) {
	/*// Markup
	baseTemplate, _ := Asset("../static/html/dashboard.html")

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
	//utilScripts, _ := Asset("../static/js/utils.js")
	//abstractDashboard, _ := Asset("../static/js/abstractDashboard.js")
	dashboard, _ := Asset("../static/js/dashboard.js")

	logsDashboard, _ := Asset("../static/js/logsDashboard.js")
	requestDashboard, _ := Asset("../static/js/requestDashboard.js")

	variables := map[string]string{
		"APPLICATION_NAME":  os.Getenv("APPLICATION_NAME"),
		"COMMON_HEADER":     MinifyHTML(header),
		"HIGHLIGHT_STYLES":  MinifyCSS(string(highlightStyles)),
		"GOSCOPE_STYLES":    MinifyCSS(string(goscopeStyles)),
		"ENTRIES_PER_PAGE":  os.Getenv("GOSCOPE_ENTRIES_PER_PAGE"),
		"COMMON_NAVBAR":     MinifyHTML(navbar),
		"COMMON_FOOTER":     MinifyHTML(string(footer)),
		"DASHBOARD_SCRIPTS": MinifyJs(string(dashboard)),
	}
	variables["LOG_DASHBOARD"] = string(logsDashboard)
	variables["REQUEST_DASHBOARD"] = string(requestDashboard)

	ShowGoScopePage(c, MinifyHTML(string(baseTemplate)), variables)*/
	file, _ := Asset("../static/goscope/dist/goscope/index.html")
	reader := strings.NewReader(string(file))
	c.DataFromReader(http.StatusOK, reader.Size(), "text/html", reader, nil)
}

func GetStaticFile(c *gin.Context) {
	pathSplit := strings.Split(c.Request.RequestURI, "/")
	requestedFile := pathSplit[len(pathSplit)-1]
	file, _ := Asset(fmt.Sprintf("../static/goscope/dist/goscope/%s", requestedFile))
	reader := strings.NewReader(string(file))
	c.DataFromReader(http.StatusOK, reader.Size(), "text/html", reader, nil)
}
