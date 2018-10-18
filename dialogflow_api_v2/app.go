package main

import (
	"log"
	"net/http"

	// "github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	df "github.com/leboncoin/dialogflow-go-webhook"
)

func webhook(c *gin.Context) {

	var err error
	var dfr *df.Request

	if err = c.BindJSON(&dfr); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// spew.Dump(dfr)

	switch dfr.QueryResult.Action {
	case "search":
		log.Println("Search action detected")
		c.JSON(http.StatusOK, gin.H{})
	case "random":
		log.Println("Search action detected")
		c.JSON(http.StatusOK, gin.H{})
	default:
		log.Println("Unknown action")
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, gin.H{})
}

func main() {
	r := gin.Default()
	r.POST("/", webhook)
	if err := r.Run("127.0.0.1:8008"); err != nil {
		panic(err)
	}
}
