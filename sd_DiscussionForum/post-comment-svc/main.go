package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"time"
)

type Post struct {
	ID        string    `json:"id" bson:"_id"`
	Title     string    `json:"title" bson:"title"`
	Content   string    `json:"content" bson:"content"`
	Author    string    `json:"author" bson:"author"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
	Tags      []string  `json:"tags" bson:"tags"`
}

type Comment struct {
	ID        string    `json:"id" bson:"_id"`
	Content   string    `json:"content" bson:"content"`
	Author    string    `json:"author" bson:"author"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
	PostID    string    `json:"post_id" bson:"post_id"`
}

var postsCollection *mongo.Collection
var commentsCollection *mongo.Collection

func main() {
	// Create a MongoDB client by establishing a connection to the MongoDB server
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27117"))
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the MongoDB server:
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	postsCollection = client.Database("mydb").Collection("posts")
	commentsCollection = client.Database("mydb").Collection("comments")

	defer client.Disconnect(ctx)

	router := gin.Default()

	router.POST("/posts", savePostHandler)
	router.POST("/comments", saveCommentHandler)

	router.Run("0.0.0.0:8080")
}

func savePostHandler(c *gin.Context) {
	// Parse the request body to extract the post data
	var post Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set the timestamp of the post
	post.Timestamp = time.Now()
	post.ID = uuid.New().String()

	// Insert the post into the database
	_, err := postsCollection.InsertOne(c, post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success response
	c.JSON(http.StatusCreated, gin.H{"message": "Post saved successfully"})
}

func saveCommentHandler(c *gin.Context) {
	// Parse the request body to extract the post data
	var comment Comment

	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post, err := getRandomPost(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if post.ID == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No posts found"})
		return
	}

	comment.Timestamp = time.Now()
	comment.ID = uuid.New().String()
	comment.PostID = post.ID

	// Insert the comment into the database
	_, err = commentsCollection.InsertOne(c, comment)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success response
	c.JSON(http.StatusCreated, gin.H{"message": "Comment saved successfully"})
}

func getRandomPost(ctx context.Context) (Post, error) {
	pipeline := []bson.M{
		{
			"$sample": bson.M{"size": 1},
		},
	}

	cursor, err := postsCollection.Aggregate(ctx, pipeline)
	if err != nil {
		return Post{}, err
	}

	defer cursor.Close(ctx)

	var result Post
	if cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return Post{}, err
		}
	}

	return result, nil
}
