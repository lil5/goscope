// License: MIT
// Authors:
// 		- Josep Jesus Bigorra Algaba (@averageflow)

package goscope

import (
	"github.com/averageflow/goscope/database"
	"github.com/gin-gonic/gin"
)

type LoggerGoScope struct {
	RoutingEngine *gin.Engine
}

func (logger LoggerGoScope) Write(p []byte) (n int, err error) {
	go database.WriteLogs(string(p))
	return len(p), nil
}
