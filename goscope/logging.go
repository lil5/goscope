// License: MIT
// Authors:
// 		- Josep Bigorra (averageflow)
package goscope

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/nu7hatch/gouuid"
)

// Log an HTTP response to the DB and print to Stdout.
func ResponseLogger(c *gin.Context) {
	if CheckExcludedPaths(c.FullPath()) {
		blw := &BodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		buf, _ := ioutil.ReadAll(c.Request.Body)
		rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
		// We have to create a new Buffer, because rdr1 will be read and consumed.
		rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))
		c.Request.Body = rdr2

		go DumpResponse(c, blw, readBody(rdr1))
	}

	c.Next()
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
	go writeLogs(string(p))
	return len(p), nil
}

func writeLogs(message string) {
	fmt.Printf("%v", message)

	db := GetDB()

	defer db.Close()

	uid, _ := uuid.NewV4()
	query := "INSERT INTO logs (uid, application, error, time) VALUES " +
		"(?, ?, ?, ?)"

	_, err := db.Exec(
		query,
		uid.String(),
		os.Getenv("APPLICATION_ID"),
		message,
		time.Now().Unix(),
	)
	if err != nil {
		log.Println(err.Error())
	}
}
