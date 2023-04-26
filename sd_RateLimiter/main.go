package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

var redisClient *redis.Client

// define the rate limiting middleware
func rateLimiter(c *gin.Context) {
	// get the user ID from the x-user-id header
	userID := c.GetHeader("x-user-id")

	if userID == "" {
		c.AbortWithStatus(400) // 400 Bad Request
	}

	// set the key and limit based on the user ID
	key := fmt.Sprintf("rate_limiter:%s", userID)
	limit := 10

	// get the current count for the key
	count, err := redisClient.Get(key).Int()
	if err != nil && err != redis.Nil {
		c.AbortWithError(500, err)
		return
	}

	c.Writer.Header().Set("X-RateLimit-Limit", fmt.Sprintf("%v/%v", count, limit))

	if count >= limit {
		c.AbortWithStatus(429) // 429 Too Many Requests
		return
	}

	if count == 0 {
		// set the initial count to 1
		err := redisClient.Set(key, 1, 5*time.Second).Err()
		if err != nil {
			c.AbortWithError(500, err)
			return
		}

		c.Next()
		return
	}

	// increment the count
	err = redisClient.Incr(key).Err()

	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.Next()
}

func main() {
	// create the Gin router
	r := gin.Default()

	// connect to Redis
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6375",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := redisClient.Ping().Result()

	if err != nil {
		log.Fatal(err)
	}

	// use the rate limiting middleware on all routes
	r.Use(rateLimiter)

	// define a test route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	// start the server
	r.Run(":8080")
}
