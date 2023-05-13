package main

import (
	"fmt"
	"main-svc/encoder"
	"main-svc/models"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type ShardsPool struct {
	dbShards   []*gorm.DB
	currentIdx int
	mu         sync.Mutex
}

type Server struct {
	pool *ShardsPool
}

func (p *ShardsPool) getShard() *gorm.DB {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.currentIdx = (p.currentIdx + 1) % len(p.dbShards)

	fmt.Println("Selected shard: ", p.currentIdx)

	return p.dbShards[p.currentIdx]
}

func (p *ShardsPool) applyAllShards(fn func(db *gorm.DB)) {
	for _, db := range p.dbShards {
		fn(db)
	}
}

func main() {
	pool := &ShardsPool{}
	server := &Server{pool: pool}

	// Define your connection strings for each shard
	connectionStrings := []string{
		"host=localhost port=5434 user=user password=password dbname=shard1 sslmode=disable",
		"host=localhost port=5435 user=user password=password dbname=shard2 sslmode=disable",
		// Add more as needed
	}

	// Initialize a pool of connections to different shards
	pool.dbShards = make([]*gorm.DB, len(connectionStrings))

	for i, connStr := range connectionStrings {
		db, err := gorm.Open("postgres", connStr)
		if err != nil {
			panic(fmt.Sprintf("failed to connect to shard %d: %v", i, err))
		}
		defer db.Close()
		pool.dbShards[i] = db
	}

	fmt.Println("Successfully connected to all shards")

	pool.applyAllShards(func(db *gorm.DB) {
		db.AutoMigrate(&models.Range{})
	})

	fmt.Println("Successfully migrated all shards")

	// seedData(pool)
	// fmt.Println("Successfully seeded all shards")

	r := gin.Default()

	r.GET("/getencodedval", server.getEncodedVal)

	r.Run() // listen and serve on 0.0.0.0:8080
}

func (s *Server) getEncodedVal(c *gin.Context) {
	url := c.Query("url")
	if url == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "url is required"})
		return
	}

	var selectedRange models.Range

	// Get random shard, pick random range
	shard := s.pool.getShard()

	shard.Order("random()").First(&selectedRange)

	fmt.Println("Current of selected range: ", selectedRange.Current)

	var updatedRange models.Range

	if err := shard.Model(&models.Range{}).Where("id = ?", selectedRange.ID).UpdateColumn("current", gorm.Expr("current + ?", 1)).Scan(&updatedRange).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	code := encoder.Encode(updatedRange.Current)

	// Save code to another db

	c.JSON(http.StatusOK, gin.H{"encoded_value": code})
}

func seedData(pool *ShardsPool) {
	ranges := []models.Range{
		{Min: 280475000000000, Max: 280475999999999, Current: 280475000000000},
		{Min: 280476000000000, Max: 280476999999999, Current: 280476000000000},
		{Min: 280477000000000, Max: 280477999999999, Current: 280477000000000},
		{Min: 280478000000000, Max: 280478999999999, Current: 280478000000000},
	}

	for _, item := range ranges {
		item.ID = uuid.New()

		pool.getShard().Create(&item)
	}
}
