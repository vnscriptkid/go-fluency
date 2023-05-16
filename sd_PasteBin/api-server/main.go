package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Paste struct {
	gorm.Model
	UserID    string
	FileID    string `gorm:"unique"`
	Content   string
	ExpiresAt time.Time
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to the database
	dsn := "user=user password=password dbname=metadata host=localhost port=5435 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	// AutoMigrate the Paste struct
	db.AutoMigrate(&Paste{})

	// Set up AWS session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_ACCESS_KEY_ID"),
			os.Getenv("AWS_SECRET_ACCESS_KEY"),
			"",
		),
	})
	if err != nil {
		log.Fatal(err)
	}
	s3Client := s3.New(sess)

	r := gin.Default()

	// Serve the HTML form
	r.GET("/", func(c *gin.Context) {
		c.File("public/index.html")
	})

	// Handle the form submission
	r.POST("/submit", func(c *gin.Context) {
		userID := c.PostForm("user_id")
		content := c.PostForm("content")
		fileID := uuid.New().String()

		tx := db.Begin()

		// Store metadata in Postgres
		paste := Paste{UserID: userID, FileID: fileID, Content: content, ExpiresAt: time.Now().Add(time.Hour * 24 * 7) /*in 7 days*/}

		err := tx.Create(&paste).Error

		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create paste"})
			return
		}

		// Upload file to S3
		_, err = s3Client.PutObject(&s3.PutObjectInput{
			Body:   aws.ReadSeekCloser(strings.NewReader(content)),
			Bucket: aws.String(os.Getenv("AWS_BUCKET_NAME")),
			Key:    aws.String(fmt.Sprintf("%s/%s", userID, fileID)),
		})

		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to upload file", "details": err.Error()})
			return
		}

		tx.Commit()

		c.Redirect(http.StatusMovedPermanently, "/")
	})

	r.GET("/get/:fileID", func(c *gin.Context) {
		fileID := c.Param("fileID")

		// Query the database to get the user_id corresponding to the file_id
		var paste Paste
		if err := db.Where("file_id = ?", fileID).First(&paste).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve file metadata from database"})
			return
		}

		userID := paste.UserID

		// Get the object from S3
		result, err := s3Client.GetObject(&s3.GetObjectInput{
			Bucket: aws.String(os.Getenv("AWS_BUCKET_NAME")),
			Key:    aws.String(fmt.Sprintf("%s/%s", userID, fileID)),
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve file from S3"})
			return
		}

		defer result.Body.Close()

		go func() {
			// Extract client data for analytics
			clientData := map[string]string{
				"userAgent": c.GetHeader("User-Agent"),
				"ip":        c.ClientIP(),
				"fileID":    fileID,
				"referrer":  c.GetHeader("Referer"),
				"dateTime":  time.Now().Format(time.RFC3339),
			}

			// Convert client data to JSON for sending to Kafka
			clientDataJSON, err := json.Marshal(clientData)

			if err != nil {
				log.Printf("Failed to marshal client data to JSON: %v", err)
				// we don't return an error to the client here because the file has already been sent,
				// and the failure to send data to Kafka doesn't impact the client
			} else {
				// Send client data to Kafka
				// msg := &kafka.ProducerMessage{
				// 	Topic: "client-data",
				// 	Value: kafka.StringEncoder(clientDataJSON),
				// }

				// _, _, err = kafkaProducer.SendMessage(msg)
				// if err != nil {
				// 	log.Printf("Failed to send message to Kafka: %v", err)
				// }
				log.Printf("Client data: %s", clientDataJSON)
			}
		}()

		// Set the proper content type so that the browser can handle response correctly
		c.Header("Content-Type", "application/octet-stream")

		// Forward the file to the client
		_, err = io.Copy(c.Writer, result.Body)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write file to response"})
			return
		}

		c.Status(http.StatusOK)
	})

	go startDeletionJob(db, s3Client)

	r.Run() // listen and serve on 0.0.0.0:8080
}

func startDeletionJob(db *gorm.DB, s3svc *s3.S3) {
	ticker := time.NewTicker(1 * time.Hour) // Adjust the interval to suit your needs

	go func() {
		for range ticker.C {
			var pastes []Paste
			db.Where("expires_at < ?", time.Now()).Find(&pastes)

			for _, paste := range pastes {
				// Delete the file from S3
				_, err := s3svc.DeleteObject(&s3.DeleteObjectInput{
					Bucket: aws.String(os.Getenv("AWS_BUCKET_NAME")),
					Key:    aws.String(fmt.Sprintf("%s/%s", paste.UserID, paste.FileID)),
				})

				if err != nil {
					log.Printf("Failed to delete file from S3: %v", err)
					continue
				}

				// Delete the record from the database
				db.Delete(&paste)
			}
		}
	}()
}
