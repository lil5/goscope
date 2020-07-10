// Copyright © 2020 Pro Warehouse B.V.
// All Rights Reserved
package goscope

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	uuid "github.com/nu7hatch/gouuid"
	"html"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

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
	db, err := sql.Open("mysql", os.Getenv("WATCHER_DATABASE_CONNECTION"))
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer db.Close()
	uid, _ := uuid.NewV4()
	query := "INSERT INTO `logs` (`uid`, `application`, `error`, `time`) VALUES " +
		"('%s', '%s', '%s', %v)"
	resultingQuery := fmt.Sprintf(query, uid, os.Getenv("APPLICATION_ID"), html.EscapeString(message), time.Now().Unix())
	_, err = db.Exec(resultingQuery)
	if err != nil {
		log.Println(err.Error())
	}
}