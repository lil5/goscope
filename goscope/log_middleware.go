// Copyright © 2020 Pro Warehouse B.V.
// All Rights Reserved
package goscope

import (
	"bytes"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"io/ioutil"
	"log"
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
