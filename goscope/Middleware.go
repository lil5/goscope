// License: MIT
// Authors:
// 		- Josep Jesus Bigorra Algaba (@averageflow)

package goscope

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/averageflow/goscope/database"

	"github.com/averageflow/goscope/utils"

	"github.com/gin-gonic/gin"
)

// Log an HTTP response to the DB and print to Stdout.
func ResponseLogger(c *gin.Context) {
	details := ObtainBodyLogWriter(c)

	c.Next()

	dumpPayload := utils.DumpResponsePayload{
		Headers: details.Blw.Header(),
		Body:    details.Blw.body,
		Status:  c.Writer.Status(),
	}

	if CheckExcludedPaths(c.FullPath()) {
		go database.DumpResponse(c, dumpPayload, readBody(details.Rdr))
	}
}

func NoRouteResponseLogger(c *gin.Context) {
	details := ObtainBodyLogWriter(c)

	dumpPayload := utils.DumpResponsePayload{
		Headers: details.Blw.Header(),
		Body:    details.Blw.body,
		Status:  http.StatusNotFound,
	}

	go database.DumpResponse(c, dumpPayload, readBody(details.Rdr))

	c.JSON(http.StatusNotFound, gin.H{
		"code":    http.StatusNotFound,
		"message": "The requested resource could not be found!",
	})
}

func ObtainBodyLogWriter(c *gin.Context) BodyLogWriterResponse {
	blw := &BodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}

	c.Writer = blw

	buf, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err.Error())
	}

	rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
	// We have to create a new Buffer, because rdr1 will be read and consumed.
	rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))
	c.Request.Body = rdr2

	return BodyLogWriterResponse{
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
