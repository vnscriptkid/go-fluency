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
	case "tokens-bucket":
		return rateLimiter_tokensBucket
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

func rateLimiter_tokensBucket(c *gin.Context) {
	const (
		rate     = 10  // 10 requests per second
		capacity = 600 // 10 tokens
	)

	var now = time.Now().Unix()
	var requested = 1

	userID := c.GetHeader("x-user-id")

	// Input:
	// tokensKey (tokens:userID, 123): the number of tokens left
	// timestampKey (timestamp:userID, 123): the timestamp of the last request
	// rate: the rate of the bucket in tokens/second
	// capacity: the capacity of the bucket in tokens
	// now: the current unix timestamp (seconds since epoch)
	// requested: the number of tokens requested for the current request
	// Output:
	// -1 if the request is rejected, otherwise the number of tokens left before the request is processed
	luaScript := `
	local key = KEYS[1]
	local tokensKey = "tokens:"..key
	local timestampKey = "timestamp:"..key

	local rate = tonumber(ARGV[1])
	local capacity = tonumber(ARGV[2])
	local now = tonumber(ARGV[3])
	local fill_time = capacity/rate
	local ttl = math.floor(fill_time*2)

	local tokens = tonumber(redis.call("get", tokensKey))
	local last_tokens = tokens

	if tokens == nil then
		tokens = capacity
		last_tokens = tokens
	else
		local last = redis.call("get", timestampKey)
		if last ~= nil then
			local elapsed = now - tonumber(last)
			local fill = elapsed * rate
			tokens = math.min(capacity, tokens + fill)
		end
	end

	local requested = tonumber(ARGV[4])
	if tokens >= requested then
		tokens = tokens - requested
		redis.call("setex", tokensKey, ttl, tokens)
		redis.call("setex", timestampKey, ttl, now)
		return last_tokens
	else
		return -1
	end
`

	sha, err := redisClient.ScriptLoad(luaScript).Result()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	res, err := redisClient.EvalSha(sha, []string{userID}, rate, capacity, now, requested).Result()

	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	if res == -1 {
		c.AbortWithStatus(429) // 429 Too Many Requests
		return
	}

	c.Next()
}
