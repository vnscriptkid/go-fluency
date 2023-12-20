package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// // Define a route for the HTML file
	r.GET("/", func(c *gin.Context) {
		c.File("./index.html")
	})

	r.StaticFS("/videos", http.Dir("./output"))

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
