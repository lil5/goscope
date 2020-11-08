package controllers

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/averageflow/goscope/src/types"

	"github.com/averageflow/goscope/src/repository"

	"github.com/averageflow/goscope/src/utils"

	"github.com/gin-gonic/gin"
)

type LoggerGoScope struct {
	RoutingEngine *gin.Engine
}

func (logger LoggerGoScope) Write(p []byte) (n int, err error) {
	go repository.WriteLogs(string(p))
	return len(p), nil
}

// Log an HTTP response to the DB and print to Stdout.
func ResponseLogger(c *gin.Context) {
	details := ObtainBodyLogWriter(c)

	c.Next()

	dumpPayload := types.DumpResponsePayload{
		Headers: details.Blw.Header(),
		Body:    details.Blw.Body,
		Status:  c.Writer.Status(),
	}

	if utils.CheckExcludedPaths(c.FullPath()) {
		go repository.DumpResponse(c, dumpPayload, readBody(details.Rdr))
	}
}

func NoRouteResponseLogger(c *gin.Context) {
	details := ObtainBodyLogWriter(c)

	dumpPayload := types.DumpResponsePayload{
		Headers: details.Blw.Header(),
		Body:    details.Blw.Body,
		Status:  http.StatusNotFound,
	}

	go repository.DumpResponse(c, dumpPayload, readBody(details.Rdr))

	c.JSON(http.StatusNotFound, gin.H{
		"code":    http.StatusNotFound,
		"message": "The requested resource could not be found!",
	})
}

func ObtainBodyLogWriter(c *gin.Context) types.BodyLogWriterResponse {
	blw := &types.BodyLogWriter{Body: bytes.NewBufferString(""), ResponseWriter: c.Writer}

	c.Writer = blw

	buf, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err.Error())
	}

	rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
	// We have to create a new Buffer, because rdr1 will be read and consumed.
	rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))
	c.Request.Body = rdr2

	return types.BodyLogWriterResponse{
		Blw: blw,
		Rdr: rdr1,
	}
}

func readBody(reader io.Reader) string {
	buf := new(bytes.Buffer)

	_, err := buf.ReadFrom(reader)
	if err != nil {
		log.Println(err.Error())
	}

	s := buf.String()

	return s
}
