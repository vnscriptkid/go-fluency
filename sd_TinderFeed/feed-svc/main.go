package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FeedItem struct {
	UserID          int       `bson:"user_id"`
	CandidateUserID int       `bson:"candidate_user_id"`
	CreatedAt       time.Time `bson:"created_at"`
	SwipeRight      bool      `bson:"swipe_right"`
}

func main() {
	// Set up MongoDB connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27117"))
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatalf("Error disconnecting from MongoDB: %v", err)
		}
	}()

	// Create the feed collection
	feedsCollection := client.Database("feed_db").Collection("feeds")

	fmt.Println("Feed collection: ", feedsCollection)

	SeedsDB(feedsCollection)
}

func SeedsDB(feedsCollection *mongo.Collection) {
	userIDs := []int{4, 5, 6, 7}

	for _, userID := range userIDs {
		for i := 0; i < 100; i++ {
			feedDoc := bson.M{
				"user_id":           userID,
				"candidate_user_id": userID*1000 + i,
				"created_at":        time.Now(),
			}

			_, err := feedsCollection.InsertOne(context.Background(), feedDoc)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	fmt.Println("Data seeding completed.")
}
