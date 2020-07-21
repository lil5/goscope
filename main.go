// License: MIT
// Authors:
// 		- Josep Bigorra (averageflow)
package main

import (
	"github.com/averageflow/goscope/goscope"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic("This project requires a .env file at the project root!")
	}
	router := gin.New()
	goscope.Setup(router)
	_ = router.Run()
}
