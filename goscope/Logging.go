// License: MIT
// Authors:
// 		- Josep Jesus Bigorra Algaba (@averageflow)
package goscope

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"

	"github.com/averageflow/goscope/database"

	"github.com/averageflow/goscope/utils"

	"github.com/gin-gonic/gin"
)

// Log an HTTP response to the DB and print to Stdout.
func ResponseLogger(c *gin.Context) {
	if CheckExcludedPaths(c.FullPath()) {
		err := LogWantedResponse(c)
		if err != nil {
			log.Printf(err.Error()) //nolint:staticcheck
			return
		}
	}

	c.Next()
}

func LogWantedResponse(c *gin.Context) error {
	blw := &BodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw

	buf, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}

	rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
	// We have to create a new Buffer, because rdr1 will be read and consumed.
	rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))
	c.Request.Body = rdr2

	go database.DumpResponse(c, utils.DumpResponsePayload{
		Headers: blw.Header(),
		Body:    blw.body,
		Status:  blw.Status(),
	}, readBody(rdr1))

	return nil
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

type LoggerGoScope struct {
	RoutingEngine *gin.Engine
}

func (logger LoggerGoScope) Write(p []byte) (n int, err error) {
	go database.WriteLogs(string(p))
	return len(p), nil
}
