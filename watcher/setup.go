package watcher

import (
	"github.com/gin-gonic/gin"
	"path/filepath"
)

func Setup(engine *gin.Engine) {
	// Load HTML templates
	absPath, _ := filepath.Abs("./static/templates/*")
	engine.LoadHTMLGlob(absPath)
	// Use the logging middleware
	engine.Use(RequestLogger)
	engine.Use(ResponseLogger)
	// Re route to static folders
	engine.Static("/js", "../static/js/")
	engine.Static("/css", "../static/css")
	// Setup necessary routes
	watcherGroup := engine.Group("/watcher")
	watcherGroup.GET("/", Dashboard)
	watcherGroup.GET("/requests", GetRequests)
	watcherGroup.GET("/requests/:id", ShowRequest)
	watcherGroup.GET("/responses", GetResponses)
	watcherGroup.GET("/responses/:id", ShowResponse)
}
