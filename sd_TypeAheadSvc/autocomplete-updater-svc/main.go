package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
)

// Local directory to store search data
const localSearchDataDir = "./searches"

// S3 details
const (
	s3Bucket = "your-bucket"
	s3Region = "your-region"
)

// AWS Credentials
const (
	awsAccessKeyID     = "your-access-key-id"
	awsSecretAccessKey = "your-secret-access-key"
)

var SERVER_NAME = os.Getenv("SERVER_NAME")

// SearchData is the structure of the JSON data received from the /search endpoint
type SearchData struct {
	Keyword string `json:"keyword"`
}

func main() {
	go startServer()
	go job2()

	// Prevent the main function from exiting
	select {}
}

// startServer starts the HTTP server using Gin
func startServer() {
	if SERVER_NAME == "" {
		SERVER_NAME = "server-1"
	}

	r := gin.Default()
	r.POST("/search", searchHandler)
	r.Run() // listen and serve on 0.0.0.0:8080
}

// searchHandler handles the /search endpoint
func searchHandler(c *gin.Context) {
	var data SearchData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dir := getDir(time.Now())

	// Create the directory if it doesn't exist
	os.MkdirAll(filepath.Join(localSearchDataDir, dir), 0755)

	// Append the search keyword to the file in the local search data directory
	localFilePath := getFullFilePath(dir)

	file, _ := os.OpenFile(localFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	defer file.Close()

	_, _ = file.WriteString(data.Keyword + "\n")
}

func getDir(now time.Time) string {
	// Get the 5-minute interval
	interval := now.Minute() / 5 * 5

	// Format the current time to create the directory
	dir := now.Format("2006/01/02/15_") + fmt.Sprintf("%02d", interval)

	return dir
}

func getPrevDir(now time.Time) (string, string) {
	prevTime := now.Add(-5 * time.Minute)
	prevOfPrevTime := now.Add(-10 * time.Minute)

	// Get the 5-minute interval
	prev := prevTime.Minute() / 5 * 5
	prevOfPrev := prevOfPrevTime.Minute() / 5 * 5

	// Format the current time to create the directory
	prevDir := prevTime.Format("2006/01/02/15_") + fmt.Sprintf("%02d", prev)
	prevOfPrevDir := prevOfPrevTime.Format("2006/01/02/15_") + fmt.Sprintf("%02d", prevOfPrev)

	return prevDir, prevOfPrevDir
}

func job2() {
	// Initialize an AWS session with given credentials and region
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(s3Region),
		Credentials: credentials.NewStaticCredentials(awsAccessKeyID, awsSecretAccessKey, ""),
	}))

	for {
		// Job 2 runs every 1 minutes
		time.Sleep(1 * time.Minute)

		// Get prev and prev of prev dir
		prevDir, prevOfPrevDir := getPrevDir(time.Now())

		fmt.Println(sess)

		for _, dir := range []string{prevDir, prevOfPrevDir} {
			// Get the local file path
			localFilePath := getFullFilePath(dir)

			// Check if the file exists
			if !fileExists(localFilePath) {
				continue
			}

			// Upload the file to S3
			// Delete the file if the upload is successful
		}
	}
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

func getFullFilePath(dir string) string {
	return filepath.Join(localSearchDataDir, dir, fmt.Sprintf("%s.txt", SERVER_NAME))
}
