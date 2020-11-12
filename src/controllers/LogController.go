package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/averageflow/goscope/src/repository"
	"github.com/averageflow/goscope/src/types"

	"github.com/averageflow/goscope/src/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func LogList(c *gin.Context) {
	offsetQuery := c.DefaultQuery("offset", "0")
	offset, _ := strconv.ParseInt(offsetQuery, 10, 32)

	variables := gin.H{
		"applicationName": utils.Config.ApplicationName,
		"entriesPerPage":  utils.Config.GoScopeEntriesPerPage,
		"data":            repository.FetchLogs(int(offset)),
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, variables)
}

func ShowLog(c *gin.Context) {
	var request types.RecordByURI

	err := c.ShouldBindUri(&request)
	if err != nil {
		log.Println(err.Error())
	}

	logDetails := repository.FetchDetailedLog(request.UID)

	variables := gin.H{
		"applicationName": utils.Config.ApplicationName,
		"data": gin.H{
			"logDetails": logDetails,
		},
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, variables)
}

func SearchLog(c *gin.Context) {
	var request types.SearchRequestPayload

	err := c.ShouldBindBodyWith(&request, binding.JSON)
	if err != nil {
		log.Println(err.Error())
	}

	offsetQuery := c.DefaultQuery("offset", "0")
	offset, _ := strconv.ParseInt(offsetQuery, 10, 32)
	result := repository.FetchSearchLogs(request.Query, int(offset))

	variables := gin.H{
		"applicationName": utils.Config.ApplicationName,
		"entriesPerPage":  utils.Config.GoScopeEntriesPerPage,
		"data":            result,
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, variables)
}

func SearchLogOptions(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.JSON(http.StatusOK, nil)
}
