package goscope

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ShowDashboard(c *gin.Context) {
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
