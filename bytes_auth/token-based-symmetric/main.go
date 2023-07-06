package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("your_secret_key")

// Create a struct to read the username and password from the request body
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// Create a struct that will be encoded to a JWT
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Signin(c *gin.Context) {
	var creds Credentials

	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Simulate checking the credentials against a database
	if creds.Username != "user" || creds.Password != "pass" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "username or password incorrect"})
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// HMAC using SHA-256
	// - HMAC: uses a cryptographic hash function in combination with a secret key
	// - SHA-256 is a specific type of hash function that produces a 256-bit (32-byte) hash value
	// - symmetric algorithm: the same key is used to sign and verify the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256 /*jwt header*/, claims /*jwt payload*/)
	tokenString, err := token.SignedString(jwtKey)

	/*
		HMACSHA256(
			base64UrlEncode(header) + "." +
			base64UrlEncode(payload),
			"your-256-bit-secret"
		)
	*/

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.SetCookie("token", tokenString, int(expirationTime.Sub(time.Now()).Seconds()), "/", "", false, true)
}

func ValidateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No Authorization header provided"})
			c.Abort()
			return
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Malformed token"})
			c.Abort()
			return
		}

		tokenStr := bearerToken[1]

		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func Welcome(c *gin.Context) {
	cookie, _ := c.Cookie("token")

	claims := &Claims{}

	token, _ := jwt.ParseWithClaims(cookie, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("Welcome, %s!", claims.Username))
}

func main() {
	router := gin.Default()

	router.POST("/signin", Signin)

	auth := router.Group("/")
	auth.Use(ValidateMiddleware())
	{
		auth.GET("/welcome", Welcome)
	}

	router.Run(":8000")
}
