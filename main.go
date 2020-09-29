// License: MIT
// Authors:
// 		- Josep Bigorra (averageflow)
package main

import (
	"log"

	"github.com/averageflow/goscope/goscope"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic("This project requires a .env file at the project root!")
	}

	router := gin.New()

	// Provide a route for the GoScope UI, which
	// can be customized i.e. to use a Auth middleware:
	//
	// ui := router.Group("/goscope").Use(gin.BasicAuth(gin.Accounts{
	//	"secret_user": "secret_password",
	// }))
	// goscope.Setup(router, ui)

	goscope.Setup(router, router.Group("/goscope"))
	_ = router.Run()
}
