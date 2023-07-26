package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Client struct {
	Appurl string `json:"appurl,omitempty"`
	IP       string `json:"ip,omitempty"`
}


var clients = []Client{}

func getappurl(c *gin.Context) {
	ID := c.Param("id")
	for _, i := range clients {
		if i.Appurl == ID {
			c.IndentedJSON(http.StatusOK, i)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "appurl not found!"})
}


func main() {
	router := gin.Default()

	router.GET("/client/:id", getappurl)

	router.Run(":8080")
}