package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", getREADME)
	router.GET("/json", getAllData)
	router.GET("/:key", getDataByKey)
	router.Run("localhost:5003")
}
