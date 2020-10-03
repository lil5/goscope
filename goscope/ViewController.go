// License: MIT
// Authors:
// 		- Josep Jesus Bigorra Algaba (@averageflow)
package goscope

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ShowDashboard(c *gin.Context) {
	file, _ := Asset("../frontend/dist/index.html")
	reader := strings.NewReader(string(file))
	c.DataFromReader(http.StatusOK, reader.Size(), "text/html", reader, nil)
}

func GetStaticFile(c *gin.Context) {
	requestedFile := strings.ReplaceAll(c.Request.RequestURI, "/goscope/", "")
	file, _ := Asset(fmt.Sprintf("../frontend/dist/%s", requestedFile))
	reader := strings.NewReader(string(file))
	c.DataFromReader(http.StatusOK, reader.Size(), GetMimeType(requestedFile), reader, nil)
}

// Get the mime type of a file by its name.
func GetMimeType(filename string) string {
	var mimeTypes = map[string]string{
		"3gp":     "video/3gpp",
		"avi":     "video/x-msvideo",
		"css":     "text/css",
		"doc":     "application/msword",
		"docx":    "application/msword",
		"exe":     "application/octet-stream",
		"gif":     "image/gif",
		"go":      "text/x-go",
		"groovy":  "text/x-groovy",
		"htm":     "text/html",
		"html":    "text/html",
		"icns":    "image/icns",
		"jpeg":    "image/jpg",
		"jpg":     "image/jpg",
		"js":      "application/javascript",
		"jsc":     "application/javascript",
		"key":     "application/vnd.apple.keynote",
		"mov":     "video/quicktime",
		"mp3":     "audio/mpeg",
		"mp4":     "video/mp4",
		"mpe":     "video/mpeg",
		"mpeg":    "video/mpeg",
		"mpg":     "video/mpeg",
		"numbers": "application/vnd.apple.numbers",
		"pages":   "application/vnd.apple.pages",
		"pdf":     "application/pdf",
		"php":     "text/html",
		"png":     "image/png",
		"ppt":     "application/vnd.ms-powerpoint",
		"py":      "text/x-python",
		"rtf":     "application/rtf",
		"swift":   "text/x-swift",
		"wav":     "audio/x-wav",
		"xls":     "application/vnd.ms-excel",
		"xlsx":    "application/vnd.ms-excel",
		"zip":     "application/zip",
	}

	split := strings.Split(filename, ".")
	extension := strings.ToLower(split[len(split)-1])

	if val, ok := mimeTypes[extension]; ok {
		return val
	}

	return extension
}
