package main

import (
	"Gorm_Playground/models"
	"encoding/json"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5438
	user     = "postgres"
	password = "123456"
	dbname   = "test"
	sslmode  = "disable"
)

func main() {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v", host, user, password, dbname, port, sslmode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to database", db)

	err = db.AutoMigrate(&models.User{}, &models.Address{})

	if err != nil {
		panic(err)
	}

	// GetUserAndHisAddresses(db, 1)
	GetUserAndHisLatestAddr(db, 1)

}

func CreateUserAlongAddresses(db *gorm.DB) {
	user := models.User{
		Name: "John Doe",
		Addresses: []models.Address{
			{Street: "123 Main St", City: "Anytown", State: "CA", Zip: "12345"},
			{Street: "456 Oak St", City: "Othertown", State: "CA", Zip: "67890"},
		},
	}

	if err := db.Create(&user); err != nil {
		log.Fatalf("error creating user: %v", err)
	}
}

func GetUserAndHisAddresses(db *gorm.DB, userId int) {
	var user models.User

	if err := db.Preload("Addresses").First(&user, userId).Error; err != nil {
		log.Fatalf("error getting user: %v", err)
	}

	userJSON, err := json.MarshalIndent(user, "", "  ")

	if err != nil {
		panic(err)
	}

	fmt.Printf("user: %s", userJSON)
}

func GetUserAndHisLatestAddr(db *gorm.DB, userId int) {
	var user models.User

	if err := db.Preload("Addresses", func(db *gorm.DB) *gorm.DB {
		return db.Order("created_at desc").Limit(1)
	}).First(&user, userId).Error; err != nil {
		log.Fatalf("error getting user: %v", err)
	}

	userJSON, err := json.MarshalIndent(user, "", "  ")

	if err != nil {
		panic(err)
	}

	fmt.Printf("user: %s", userJSON)
}
