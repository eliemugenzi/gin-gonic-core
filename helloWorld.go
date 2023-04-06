package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
func helloWorld(c *gin.Context) {
   c.JSON(http.StatusOK, gin.H {
    "message": "Hello World",
   })
}

func goodbyeWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H {
		"message": "Goodbye guys, it was fun...",
	})
}