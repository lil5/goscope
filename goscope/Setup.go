// License: MIT
// Authors:
// 		- Josep Jesus Bigorra Algaba (@averageflow)
package goscope

import (
	"log"

	"github.com/averageflow/goscope/utils"
	"github.com/gin-gonic/gin"

	// Import MYSQL Driver.
	_ "github.com/go-sql-driver/mysql"
	// Import PostgreSQL Driver.
	_ "github.com/lib/pq"
	// Import SQLite driver.
	_ "github.com/mattn/go-sqlite3"
)

func Setup(engine *gin.Engine, goScopeGroup *gin.RouterGroup) {
	utils.ConfigSetup()

	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	logger := &LoggerGoScope{}
	gin.DefaultErrorWriter = logger

	log.SetFlags(log.Lshortfile)
	log.SetOutput(logger)
	// Use the logging middleware
	engine.Use(ResponseLogger)

	// SPA routes
	// Static content from SPA
	goScopeGroup.GET("/", ShowDashboard)
	goScopeGroup.GET("/logo.svg", GetStaticFile)
	goScopeGroup.GET("/js/app.js", GetStaticFile)
	goScopeGroup.GET("/js/app.js.map", GetStaticFile)
	goScopeGroup.GET("/css/app.css", GetStaticFile)
	goScopeGroup.GET("/css/dark.css", GetStaticFile)
	goScopeGroup.GET("/css/styles.css", GetStaticFile)
	goScopeGroup.GET("/favicon.ico", GetStaticFile)
	goScopeGroup.GET("/logs", ShowDashboard)
	goScopeGroup.GET("/logs/:uuid", ShowDashboard)
	goScopeGroup.GET("/requests", ShowDashboard)
	goScopeGroup.GET("/requests/:uuid", ShowDashboard)
	goScopeGroup.GET("/info", ShowDashboard)
	// GoScope API
	apiGroup := goScopeGroup.Group("/api")
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
