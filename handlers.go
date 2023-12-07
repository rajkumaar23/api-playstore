package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func getREADME(c *gin.Context) {
	readme, err := os.ReadFile("README.md")
	if err != nil {
		c.Data(http.StatusInternalServerError, "text/html; charset=utf-8", []byte("an internal error occurred"))
	}

	c.Data(http.StatusOK, "text/html; charset=utf-8", mdToHTML(readme))
}

func getPlaystoreJSON(c *gin.Context) {
	packageID := c.Query("id")
	if packageID == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "package id is mandatory"})
		return
	}

	resBody, statusCode := fetchHTML(packageID)
	if statusCode == http.StatusNotFound {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "package id is invalid"})
		return
	} else if statusCode != 200 {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"message": "an unexpected error occurred"})
		return
	}

	parsedPlaystoreData, err := parsePlaystoreData(packageID, resBody)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "an unexpected error occurred", "error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, *parsedPlaystoreData)
}
