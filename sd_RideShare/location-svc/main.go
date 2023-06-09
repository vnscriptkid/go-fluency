package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/Shopify/sarama"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var db *gorm.DB

type LocationUpdateRequest struct {
	DriverID int     `json:"driver_id"`
	Lat      float64 `json:"lat"`
	Lng      float64 `json:"lng"`
}

type Driver struct {
	gorm.Model
	CurrentLocation interface{} `gorm:"type:geometry"`
}

func main() {
	// Connect to the master pg db
	for {
		var err error

		dsn := "host=localhost user=user password=password dbname=locations port=5434 sslmode=disable"

		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			log.Println("Write DB connection failed. Retrying...", err)
			time.Sleep(1 * time.Second)
			continue
		}

		break
	}

	log.Println("Connected to DB")

	db.AutoMigrate(&Driver{})

	// Initialize Kafka consumer
	consumer, err := sarama.NewConsumer([]string{"localhost:29092"}, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	topic := "driver-locations"

	// Subscribe to the 'driver-locations' topic
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// Process messages
	for msg := range partitionConsumer.Messages() {
		var location LocationUpdateRequest
		err = json.Unmarshal(msg.Value, &location)

		if err != nil {
			log.Printf("Failed to unmarshal product event: %v\n", err)
			continue
		}

		// Update the driver's location in the database
		err = db.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&Driver{
			Model:           gorm.Model{ID: uint(location.DriverID)},
			CurrentLocation: gorm.Expr("ST_SetSRID(ST_Point(?, ?),4326)", location.Lng, location.Lat),
		}).Error

		if err != nil {
			log.Printf("Failed to update driver location: %v\n", err)
		}

		log.Printf("Updated driver %d location to %f, %f\n", location.DriverID, location.Lat, location.Lng)
	}
}
