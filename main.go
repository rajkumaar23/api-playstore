package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", getREADME)
	router.GET("/json", getPlaystoreJSON)
	router.Run("localhost:5003")
}
