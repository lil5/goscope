// Copyright © 2020 Pro Warehouse B.V.
// All Rights Reserved
package goscope

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

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

func Setup(engine *gin.Engine) {
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
	goscopeGroup := engine.Group("/goscope")
	goscopeGroup.GET("/", Dashboard)
	goscopeGroup.GET("/logs", LogDashboard)
	goscopeGroup.GET("/log-records", GetLogs)
	goscopeGroup.GET("/log-records/:id", ShowLog)
	goscopeGroup.GET("/requests", GetRequests)
	goscopeGroup.GET("/requests/:id", ShowRequest)
}
