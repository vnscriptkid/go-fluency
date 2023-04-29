package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func addUserLocation(ctx context.Context, rdb *redis.Client, userID string, lat float64, lon float64) error {
	_, err := rdb.GeoAdd(ctx, "user_locations", &redis.GeoLocation{
		Longitude: lon,
		Latitude:  lat,
		Name:      userID,
	}).Result()
	return err
}

func main() {
	// Set up the Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Set up the Gin router
	router := gin.Default()

	// Endpoint for storing a user's location
	router.POST("/store-location/:userID/:lat/:lon", func(c *gin.Context) {
		userID := c.Param("userID")
		lat, err := strconv.ParseFloat(c.Param("lat"), 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid latitude"})
			return
		}
		lon, err := strconv.ParseFloat(c.Param("lon"), 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid longitude"})
			return
		}

		err = addUserLocation(c, rdb, userID, lat, lon)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error storing user location"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User location stored"})
	})

	// Endpoint for retrieving nearby users
	router.GET("/nearby-users/:userID", func(c *gin.Context) {
		userID := c.Param("userID")

		lon, lat, err := getUserLocation(c, rdb, userID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error retrieving user location"})
			return
		}

		nearbyUsers, err := getNearbyUsers(c, rdb, userID, lat, lon, 1.0, 10.0, 1.0)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving nearby users"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"nearbyUsers": nearbyUsers})
	})

	// Start the Gin server
	router.Run(":8080")
}

func getUserLocation(ctx context.Context, rdb *redis.Client, userID string) (float64, float64, error) {
	location, err := rdb.GeoPos(ctx, "user_locations", userID).Result()
	if err != nil {
		return 0, 0, err
	}
	if len(location) == 0 {
		return 0, 0, fmt.Errorf("user location not found")
	}
	return location[0].Longitude, location[0].Latitude, nil
}

func getNearbyUsers(ctx context.Context, rdb *redis.Client, userID string, lat float64, lon float64, radius float64, maxRadius float64, radiusIncrement float64) ([]redis.GeoLocation, error) {
	var nearbyUsers []redis.GeoLocation
	var err error
	foundNewUsers := false

	for !foundNewUsers && radius <= maxRadius {
		nearbyUsers, err = rdb.GeoRadius(ctx, "user_locations", lon, lat, &redis.GeoRadiusQuery{
			Radius:    radius,
			Unit:      "km",
			WithCoord: true,
			WithDist:  true,
			Sort:      "ASC",
		}).Result()

		if err != nil {
			return nil, err
		}

		newUsers := make([]redis.GeoLocation, 0, len(nearbyUsers))
		for _, user := range nearbyUsers {
			isInHistory, err := isUserInHistory(ctx, rdb, userID, user.Name)
			if err != nil {
				return nil, err
			}
			if !isInHistory && user.Name != userID { // Exclude users that are in history or the user themself
				newUsers = append(newUsers, user)
			}
		}

		if len(newUsers) > 0 {
			foundNewUsers = true
			nearbyUsers = newUsers
		} else {
			radius += radiusIncrement
		}
	}

	return nearbyUsers, nil
}

// Define a function to check if a user is in another user's history:
func isUserInHistory(ctx context.Context, rdb *redis.Client, userID string, seenUserID string) (bool, error) {
	return rdb.SIsMember(ctx, "user_history:"+userID, seenUserID).Result()
}
