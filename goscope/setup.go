// Copyright © 2020 Pro Warehouse B.V.
// All Rights Reserved
package goscope

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Setup(engine *gin.Engine) {
	logger := &LoggerGoScope{}
	gin.DefaultErrorWriter = logger
	log.SetFlags(0)
	log.SetOutput(logger)
	// Use the logging middleware
	engine.Use(ResponseLogger)
	// Setup necessary routes
	watcherGroup := engine.Group("/goscope")
	watcherGroup.GET("/", Dashboard)
	watcherGroup.GET("/logs", LogDashboard)
	watcherGroup.GET("/log-records", GetLogs)
	watcherGroup.GET("/requests", GetRequests)
	watcherGroup.GET("/requests/:id", ShowRequest)
}
