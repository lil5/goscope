package watcher

import (
	"bytes"
	"github.com/gin-gonic/gin"
)

type SummarizedRequest struct {
	Method         string `json:"method,exists"`
	Path           string `json:"path,exists"`
	Time           int    `json:"time,exists"`
	Uid            string `json:"uid,exists"`
	ResponseStatus int    `json:"response_status,exists"`
}

type RecordByUri struct {
	Uid string `uri:"id" binding:"required"`
}

type SummarizedResponse struct {
	RequestUid string `json:"request_uid,exists"`
	ClientIp   string `json:"client_ip,exists"`
	Path       string `json:"path,exists"`
	Status     string `json:"status,exists"`
	Time       int    `json:"time,exists"`
	Uid        string `json:"uid,exists"`
}

type DetailedResponse struct {
	Body       string `json:"body,exists"`
	ClientIp   string `json:"client_ip,exists"`
	Headers    string `json:"headers,exists"`
	Path       string `json:"path,exists"`
	Size       int    `json:"size,exists"`
	Status     string `json:"status,exists"`
	Time       int    `json:"time,exists"`
	RequestUid string `json:"request_uid,exists"`
	Uid        string `json:"uid,exists"`
}

type DetailedRequest struct {
	Body      string `json:"body,exists"`
	ClientIp  string `json:"client_ip,exists"`
	Headers   string `json:"headers,exists"`
	Host      string `json:"host,exists"`
	Method    string `json:"method,exists"`
	Path      string `json:"path,exists"`
	Referrer  string `json:"referrer,exists"`
	Time      int    `json:"time,exists"`
	Uid       string `json:"uid,exists"`
	Url       string `json:"url,exists"`
	UserAgent string `json:",exists"`
}

type BodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w BodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
