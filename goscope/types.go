// Copyright © 2020 Pro Warehouse B.V.
// All Rights Reserved
package goscope

import (
	"bytes"
	"github.com/gin-gonic/gin"
)

type SummarizedRequest struct {
	Method         string `json:"method"`
	Path           string `json:"path"`
	Time           int    `json:"time"`
	Uid            string `json:"uid"`
	ResponseStatus int    `json:"response_status"`
}

type ExceptionRecord struct {
	Error string `json:"error"`
	Time  int    `json:"time"`
	Uid   string `json:"uid"`
}
type RecordByUri struct {
	Uid string `uri:"id" binding:"required"`
}

type SummarizedResponse struct {
	RequestUid string `json:"request_uid"`
	ClientIp   string `json:"client_ip"`
	Path       string `json:"path"`
	Status     string `json:"status"`
	Time       int    `json:"time"`
	Uid        string `json:"uid"`
}

type DetailedResponse struct {
	Body       string `json:"body"`
	ClientIp   string `json:"client_ip"`
	Headers    string `json:"headers"`
	Path       string `json:"path"`
	Size       int    `json:"size"`
	Status     string `json:"status"`
	Time       int    `json:"time"`
	RequestUid string `json:"request_uid"`
	Uid        string `json:"uid"`
}

type DetailedRequest struct {
	Body      string `json:"body"`
	ClientIp  string `json:"client_ip"`
	Headers   string `json:"headers"`
	Host      string `json:"host"`
	Method    string `json:"method"`
	Path      string `json:"path"`
	Referrer  string `json:"referrer"`
	Time      int    `json:"time"`
	Uid       string `json:"uid"`
	Url       string `json:"url"`
	UserAgent string `json:""`
}

type BodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w BodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
