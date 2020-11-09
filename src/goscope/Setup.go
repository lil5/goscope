package goscope

import (
	"log"

	"github.com/averageflow/goscope/src/controllers"

	"github.com/averageflow/goscope/src/utils"
	"github.com/gin-gonic/gin"
)

// Setup is the necessary step to enable GoScope in an application.
// It will setup the necessary routes and middlewares for GoScope to work.
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

	logger := &controllers.LoggerGoScope{}
	gin.DefaultErrorWriter = logger

	log.SetFlags(log.Lshortfile)
	log.SetOutput(logger)

	// Use the logging middleware
	router.Use(controllers.ResponseLogger)

	// Catch 404s
	router.NoRoute(controllers.NoRouteResponseLogger)

	// SPA routes
	if !utils.Config.HasFrontendDisabled {
		goScopeGroup.GET("/", controllers.ShowDashboard)
		goScopeGroup.GET("/logo.svg", controllers.GetStaticFile)
		goScopeGroup.GET("/js/app.js", controllers.GetStaticFile)
		goScopeGroup.GET("/js/app.js.map", controllers.GetStaticFile)
		goScopeGroup.GET("/css/app.css", controllers.GetStaticFile)
		goScopeGroup.GET("/css/dark.css", controllers.GetStaticFile)
		goScopeGroup.GET("/css/light.css", controllers.GetStaticFile)
		goScopeGroup.GET("/css/code-blocks.css", controllers.GetStaticFile)
		goScopeGroup.GET("/css/styles.css", controllers.GetStaticFile)
		goScopeGroup.GET("/favicon.ico", controllers.GetStaticFile)
		goScopeGroup.GET("/favicon-32x32.png", controllers.GetStaticFile)
		goScopeGroup.GET("/apple-touch-icon-precomposed.png", controllers.GetStaticFile)
		goScopeGroup.GET("/apple-touch-icon.png", controllers.GetStaticFile)
		goScopeGroup.GET("/favicon-16x16.png", controllers.GetStaticFile)
		goScopeGroup.GET("/logs", controllers.ShowDashboard)
		goScopeGroup.GET("/logs/:uuid", controllers.ShowDashboard)
		goScopeGroup.GET("/requests", controllers.ShowDashboard)
		goScopeGroup.GET("/requests/:uuid", controllers.ShowDashboard)
		goScopeGroup.GET("/info", controllers.ShowDashboard)
	}

	// GoScope API
	apiGroup := goScopeGroup.Group("/api")
	apiGroup.GET("/application-name", controllers.GetAppName)
	apiGroup.GET("/logs", controllers.LogList)
	apiGroup.GET("/requests/:id", controllers.ShowRequest)
	apiGroup.GET("/logs/:id", controllers.ShowLog)
	apiGroup.GET("/requests", controllers.RequestList)
	apiGroup.POST("/search/requests", controllers.SearchRequest)
	apiGroup.OPTIONS("/search/requests", controllers.SearchRequestOptions)
	apiGroup.POST("/search/logs", controllers.SearchLog)
	apiGroup.OPTIONS("/search/logs", controllers.SearchLogOptions)
	apiGroup.GET("/info", controllers.ShowSystemInfo)
}
