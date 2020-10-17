// License: MIT
// Authors:
// 		- Josep Jesus Bigorra Algaba (@averageflow)
package goscope

import (
	"log"
	"net/http"

	"github.com/averageflow/goscope/utils"
	"github.com/gin-gonic/gin"

	// Import MYSQL Driver.
	_ "github.com/go-sql-driver/mysql"
	// Import PostgreSQL Driver.
	_ "github.com/lib/pq"
	// Import SQLite driver.
	_ "github.com/mattn/go-sqlite3"
)

func Setup(router *gin.Engine, goScopeGroup *gin.RouterGroup) {
	utils.ConfigSetup()
	utils.DatabaseSetup(utils.DatabaseInformation{
		Type:                  utils.Config.GoScopeDatabaseType,
		Connection:            utils.Config.GoScopeDatabaseConnection,
		MaxOpenConnections:    utils.Config.GoScopeDatabaseMaxOpenConnections,
		MaxIdleConnections:    utils.Config.GoScopeDatabaseMaxIdleConnections,
		MaxConnectionLifetime: utils.Config.GoScopeDatabaseMaxConnLifetime,
	})

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	logger := &LoggerGoScope{}
	gin.DefaultErrorWriter = logger

	log.SetFlags(log.Lshortfile)
	log.SetOutput(logger)
	// Use the logging middleware
	router.Use(ResponseLogger)

	// Catch 404s
	router.NoRoute(func(c *gin.Context) {
		err := LogWantedResponse(c)
		if err != nil {
			log.Printf(err.Error()) //nolint:staticcheck
			c.Next()
		}

		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "The requested resource could not be found!",
		})
	})

	// SPA routes
	if !utils.Config.HasFrontendDisabled {
		goScopeGroup.GET("/", ShowDashboard)
		goScopeGroup.GET("/logo.svg", GetStaticFile)
		goScopeGroup.GET("/js/app.js", GetStaticFile)
		goScopeGroup.GET("/js/app.js.map", GetStaticFile)
		goScopeGroup.GET("/css/app.css", GetStaticFile)
		goScopeGroup.GET("/css/dark.css", GetStaticFile)
		goScopeGroup.GET("/css/light.css", GetStaticFile)
		goScopeGroup.GET("/css/code-blocks.css", GetStaticFile)
		goScopeGroup.GET("/css/styles.css", GetStaticFile)
		goScopeGroup.GET("/favicon.ico", GetStaticFile)
		goScopeGroup.GET("/favicon-32x32.png", GetStaticFile)
		goScopeGroup.GET("/apple-touch-icon-precomposed.png", GetStaticFile)
		goScopeGroup.GET("/apple-touch-icon.png", GetStaticFile)
		goScopeGroup.GET("/favicon-16x16.png", GetStaticFile)
		goScopeGroup.GET("/logs", ShowDashboard)
		goScopeGroup.GET("/logs/:uuid", ShowDashboard)
		goScopeGroup.GET("/requests", ShowDashboard)
		goScopeGroup.GET("/requests/:uuid", ShowDashboard)
		goScopeGroup.GET("/info", ShowDashboard)
	}

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
