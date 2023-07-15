package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

var redisClient *redis.Client

func main() {
	// Create the Gin router
	r := gin.Default()

	// Connect to Redis
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6376",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := redisClient.Ping().Result()

	if err != nil {
		log.Fatal(err)
	}

	// Parse the URLs of the backend services.
	service1URL, _ := url.Parse("http://service1-url")
	service2URL, _ := url.Parse("http://service2-url")

	// Create reverse proxies for the backend services.
	proxy1 := httputil.NewSingleHostReverseProxy(service1URL)
	proxy2 := httputil.NewSingleHostReverseProxy(service2URL)

	// Use the checkBlacklist middleware.
	r.Use(checkBlacklist())

	// Route requests to /service1/* to the first service.
	r.Any("/service1/*any", func(c *gin.Context) {
		proxy1.ServeHTTP(c.Writer, c.Request)
	})

	// Route requests to /service2/* to the second service.
	r.Any("/service2/*any", func(c *gin.Context) {
		proxy2.ServeHTTP(c.Writer, c.Request)
	})

	r.Run(":8080") // listen and serve on :8080
}

func checkBlacklist() gin.HandlerFunc {
	return func(c *gin.Context) {
		// For simplicity, assume user ID comes as a header.
		userID := c.GetHeader("User-ID")
		if isUserBlacklisted(userID) {
			// If the user is blacklisted, return an error response and abort the request.
			c.JSON(http.StatusForbidden, gin.H{"error": "User is blacklisted."})
			c.Abort()
			return
		}

		// If the user is not blacklisted, continue processing the request.
		c.Next()
	}
}

// Sidecar/GlobalCache that can be plugged into any services
// Pros: Reduce network hops, reduce latency
// Cons: Support for multiple languages in case of polyglot services
func isUserBlacklisted(userID string) bool {
	// Check if the user ID is in the Redis set of blacklisted users.
	isBlacklisted, err := redisClient.SIsMember("blacklistedUsers", userID).Result()

	if err != nil {
		// Handle error in a better way in a real-world scenario.
		panic(err)
	}

	return isBlacklisted
}
