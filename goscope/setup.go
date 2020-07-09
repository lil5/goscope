// Copyright © 2020 Pro Warehouse B.V.
// All Rights Reserved
package goscope

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Setup(engine *gin.Engine) {
	engine.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%d - %s - %s - %s - %d - %s\n%s",
			time.Now().Unix(),
			param.ClientIP,
			param.Method,
			param.Path,
			param.StatusCode,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	// Use the logging middleware
	engine.Use(ResponseLogger)
	// Setup necessary routes
	watcherGroup := engine.Group("/goscope")
	watcherGroup.GET("/", Dashboard)
	watcherGroup.GET("/logs", GetLogs)
	watcherGroup.GET("/requests", GetRequests)
	watcherGroup.GET("/requests/:id", ShowRequest)
}
