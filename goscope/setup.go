// License: MIT
// Authors:
// 		- Josep Bigorra (averageflow)
package goscope

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	// Import MYSQL Driver
	_ "github.com/go-sql-driver/mysql"
	// Import PostgreSQL Driver
	_ "github.com/lib/pq"
	// Import SQLite driver
	_ "github.com/mattn/go-sqlite3"
)

// Ensure necessary application variables are set.
func CheckVariablesAreSet() {
	variables := []string{
		"APPLICATION_ID",
		"APPLICATION_NAME",
		"APPLICATION_TIMEZONE",
		"GOSCOPE_DATABASE_CONNECTION",
		"GOSCOPE_DATABASE_TYPE",
		"GOSCOPE_ENTRIES_PER_PAGE",
	}
	for _, s := range variables {
		if os.Getenv(s) == "" {
			panic(fmt.Sprintf("%s variable is not set", s))
		}
	}
}

func Setup(engine *gin.Engine, goscopeGroup *gin.RouterGroup) {
	CheckVariablesAreSet()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	logger := &LoggerGoScope{}
	gin.DefaultErrorWriter = logger

	log.SetFlags(log.Lshortfile)
	log.SetOutput(logger)
	// Use the logging middleware
	engine.Use(ResponseLogger)

	// Setup necessary routes
	// Static content from SPA
	goscopeGroup.GET("/", ShowDashboard)
	goscopeGroup.GET("/img/logo.svg", GetStaticFile)
	goscopeGroup.GET("/js/app.js", GetStaticFile)
	goscopeGroup.GET("/js/app.js.map", GetStaticFile)
	goscopeGroup.GET("/css/app.css", GetStaticFile)
	goscopeGroup.GET("/css/dark.css", GetStaticFile)
	goscopeGroup.GET("/css/styles.css", GetStaticFile)
	goscopeGroup.GET("/favicon.ico", GetStaticFile)
	// SPA routes
	goscopeGroup.GET("/logs", ShowDashboard)
	goscopeGroup.GET("/logs/:uuid", ShowDashboard)
	goscopeGroup.GET("/requests", ShowDashboard)
	goscopeGroup.GET("/requests/:uuid", ShowDashboard)
	goscopeGroup.GET("/info", ShowDashboard)
	// GoScope API
	apiGroup := goscopeGroup.Group("/api")
	apiGroup.GET("/application-name", GetAppName)
	apiGroup.GET("/logs", LogList)
	apiGroup.GET("/requests/:id", ShowRequest)
	apiGroup.GET("/logs/:id", ShowLog)
	apiGroup.GET("/requests", RequestList)
	apiGroup.POST("/search/requests", SearchRequest)
	apiGroup.OPTIONS("/search/requests", SearchRequestOptions)
	apiGroup.POST("/search/logs", SearchLog)
	apiGroup.OPTIONS("/search/logs", SearchLogOptions)
	apiGroup.GET("/info", ShowSystemInfo)
}
