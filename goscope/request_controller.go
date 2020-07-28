package goscope

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func RequestList(c *gin.Context) {
	offsetQuery := c.DefaultQuery("offset", "0")
	offset, _ := strconv.ParseInt(offsetQuery, 10, 32)
	entriesPerPage, _ := strconv.ParseInt(os.Getenv("GOSCOPE_ENTRIES_PER_PAGE"), 10, 32)
	variables := gin.H{
		"applicationName": os.Getenv("APPLICATION_NAME"),
		"entriesPerPage":  entriesPerPage,
		"data":            GetRequests(int(offset)),
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, variables)
}

func ShowRequest(c *gin.Context) {
	var request RecordByURI

	err := c.ShouldBindUri(&request)
	if err != nil {
		log.Println(err.Error())
	}

	requestDetails := GetDetailedRequest(request.UID)
	responseDetails := GetDetailedResponse(request.UID)

	variables := gin.H{
		"applicationName": os.Getenv("APPLICATION_NAME"),
		"data": gin.H{
			"request":  requestDetails,
			"response": responseDetails,
		},
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, variables)
}

func SearchRequest(c *gin.Context) {
	var request SearchRequestPayload
	err := c.ShouldBindBodyWith(&request, binding.JSON)

	if err != nil {
		log.Println(err.Error())
	}

	offsetQuery := c.DefaultQuery("offset", "0")
	offset, _ := strconv.ParseInt(offsetQuery, 10, 32)
	result := SearchRequests(request.Query, int(offset))

	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, result)
}
