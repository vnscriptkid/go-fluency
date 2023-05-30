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

	// // Serve the HLS playlist
	// r.GET("/playlist.m3u8", func(c *gin.Context) {
	// 	c.File("../transcoding-svc/output/playlist.m3u8")
	// })

	// // Serve the video segments
	// r.GET("/segment/:bitrate/:filename", func(c *gin.Context) {
	// 	bitrate := c.Param("bitrate")
	// 	filename := c.Param("filename")
	// 	c.File("../transcoding-svc/output/output_" + bitrate + "/" + filename)
	// })

	r.StaticFS("/videos", http.Dir("./output"))

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
