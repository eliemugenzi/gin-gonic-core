package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	/*
	  How to write a log file
	*/

	// Disabling Console color, we don't need console color when writing the logs to the file
	gin.DisableConsoleColor()

	// Logging to the file `gin.log`
	f, _ := os.Create("gin.log")

	/*
	  Setting a default writer. 
	  If you want to only write logs into the file, not the console, use the following statement:
	  gin.DefaultWriter = io.MultiWriter(f)

	*/
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := gin.Default()

	// Groupping routes
	v1 := router.Group("/v1")
	{
		v1.GET("/hello", helloWorld)
		v1.GET("/bye", goodbyeWorld)
	}

	router.GET("/ping", func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H {
			"message" : "pong",
		})
	})

	portString := fmt.Sprintf(":%v", os.Getenv("PORT"))
	listener, _ := net.Listen("tcp", portString)
	_, port, _ := net.SplitHostPort(listener.Addr().String())

	fmt.Println("Listening on port ",port)

	http.Serve(listener, router)
}