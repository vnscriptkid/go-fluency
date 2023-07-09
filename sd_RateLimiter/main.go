package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type RateLimitStrategy string

var RateLimitStrategy_fixedWindow RateLimitStrategy = "fixed-window"
var RateLimitStrategy_slidingWindow RateLimitStrategy = "sliding-window"

var redisClient *redis.Client

// define the rate limiting middleware
// https://redislabs.com/redis-best-practices/basic-rate-limiting/
// Fixed Window Counter: 10 requests / 5 seconds
// Pros: Simple, easy to implement (user_id: count)
// Cons: thundering herd problem, happens between 2 windows (in short period of time, huge amount of requests may be sent and still be accepted)
func rateLimiter_fixedWindow(c *gin.Context) {
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
		// set the initial count to 0
		err := redisClient.Set(key, 0, 5*time.Second).Err()

		if err != nil {
			c.AbortWithError(500, err)
			return
		}
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
	r.Use(NewRateLimiter(RateLimitStrategy_slidingWindow))

	// define a test route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	// start the server
	r.Run(":8080")
}

func NewRateLimiter(strategy RateLimitStrategy) func(c *gin.Context) {
	switch strategy {
	case "fixed-window":
		return rateLimiter_fixedWindow
	case "sliding-window":
		return rateLimiter_slidingWindow
	default:
		panic("invalid rate limiting strategy")
	}
}

// Pros: Solve the thundering herd problem
// Cons: footprint is bigger (user_id: timestamp)
func rateLimiter_slidingWindow(c *gin.Context) {
	// get the user ID from the x-user-id header
	userID := c.GetHeader("x-user-id")

	if userID == "" {
		c.AbortWithStatus(400) // 400 Bad Request
	}

	const Limit = 10
	const Window = time.Second * 5

	// set the key and limit based on the user ID
	key := fmt.Sprintf("rate_limiter:%s", userID)

	now := time.Now().UnixNano()
	windowStart := now - int64(Window)

	fmt.Printf("now: %v\n, windowStart: %v\n", now, windowStart)

	// Clean up old requests from the sliding window.
	_, err := redisClient.ZRemRangeByScore(key, "-inf", fmt.Sprintf("%v", windowStart)).Result()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	// Count the number of requests in the sliding window.
	count, err := redisClient.ZCount(key, fmt.Sprintf("%v", windowStart), "+inf").Result()
	c.Writer.Header().Set("X-RateLimit-Limit", fmt.Sprintf("%v/%v", count, Limit))
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	// Check if the rate limit has been reached.
	if count >= Limit {
		c.AbortWithStatus(429) // 429 Too Many Requests
		return
	}

	// Add the new request to the sliding window.
	_, err = redisClient.ZAdd(key, redis.Z{Score: float64(now), Member: now}).Result()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	// Automatically clean up old requests from the sliding window by setting an expiry time.
	_ = redisClient.Expire(key, Window)

	c.Next()
}
