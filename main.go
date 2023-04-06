package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := gin.Default()

	r.GET("/ping", func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H {
			"message" : "pong",
		})
	})

	portString := fmt.Sprintf(":%v", os.Getenv("PORT"))
	ln, _ := net.Listen("tcp", portString)
	_, port, _ := net.SplitHostPort(ln.Addr().String())

	fmt.Println("Listening on port ",port)

	http.Serve(ln, r)
}