package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Product struct {
	ID          uint    `gorm:"primary_key" json:"id"`
	Name        string  `gorm:"size:255" json:"name"`
	Description string  `gorm:"size:255" json:"description"`
	Price       float64 `json:"price"`
}

var db *gorm.DB
var dbRead *gorm.DB
var err error
var dbReadIdx int

type Config struct {
	DbHost      string
	DbPort      string
	DbUser      string
	DbName      string
	DbPassword  string
	AppHostPath string
	AppHostPort string
	DbReadHost  string
}

func GetDbRead() *gorm.DB {
	DBs := []*gorm.DB{
		db,
		dbRead,
	}

	dbReadIdx = (dbReadIdx + 1) % len(DBs)

	log.Printf("Using DB %d\n", dbReadIdx)

	return DBs[dbReadIdx]
}

func LoadConfig() *Config {
	cfg := &Config{
		DbHost:      os.Getenv("DB_HOST"),
		DbReadHost:  os.Getenv("DB_READ_HOST"),
		DbPort:      os.Getenv("DB_PORT"),
		DbUser:      os.Getenv("DB_USER"),
		DbName:      os.Getenv("DB_NAME"),
		DbPassword:  os.Getenv("DB_PASSWORD"),
		AppHostPath: os.Getenv("APP_HOST_PATH"),
		AppHostPort: os.Getenv("APP_HOST_PORT"),
	}

	if cfg.DbHost == "" {
		cfg.DbHost = "0.0.0.0" // Do not put localhost here
	}

	if cfg.DbPort == "" {
		cfg.DbPort = "5433"
	}

	if cfg.DbUser == "" {
		cfg.DbUser = "postgres"
	}

	if cfg.DbName == "" {
		cfg.DbName = "catalog"
	}

	if cfg.DbPassword == "" {
		cfg.DbPassword = "password"
	}

	if cfg.AppHostPath == "" {
		cfg.AppHostPath = "localhost"
	}

	if cfg.AppHostPort == "" {
		cfg.AppHostPort = "8080"
	}

	fmt.Printf("ENV: %+v\n", cfg)

	return cfg
}

func main() {
	cfg := LoadConfig()

	// Connect to the master pg db
	for {
		db, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbName, cfg.DbPassword))

		if err != nil {
			log.Println("Write DB connection failed. Retrying...")
			time.Sleep(1 * time.Second)
			continue
		}

		break
	}

	for {
		dbRead, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", cfg.DbReadHost, cfg.DbPort, cfg.DbUser, cfg.DbName, cfg.DbPassword))

		if err != nil {
			log.Println("Read DB connection failed. Retrying...")
			time.Sleep(1 * time.Second)
			continue
		}

		break
	}

	defer db.Close()
	db.AutoMigrate(&Product{})

	r := gin.Default()

	// Register the middleware
	r.Use(AppInstanceMiddleware())

	r.GET("health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.GET("/products", GetProducts)
	r.GET("/products/:id", GetProduct)
	r.POST("/products", CreateProduct)
	r.PUT("/products/:id", UpdateProduct)
	r.DELETE("/products/:id", DeleteProduct)

	r.Run(fmt.Sprintf("%s:%s", cfg.AppHostPath, cfg.AppHostPort))
}

func GetProducts(c *gin.Context) {
	var products []Product
	if err := GetDbRead().Find(&products).Error; err != nil {
		c.AbortWithStatus(404)
	} else {
		c.JSON(http.StatusOK, products)
	}
}

func GetProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var product Product
	if err := GetDbRead().Where("id = ?", id).First(&product).Error; err != nil {
		c.AbortWithStatus(404)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func CreateProduct(c *gin.Context) {
	var product Product
	c.BindJSON(&product)

	db.Create(&product)
	c.JSON(http.StatusCreated, product)
}

func UpdateProduct(c *gin.Context) {
	var product Product
	id := c.Params.ByName("id")

	if err := db.Where("id = ?", id).First(&product).Error; err != nil {
		c.AbortWithStatus(404)
	}
	c.BindJSON(&product)

	db.Save(&product)
	c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var product Product
	d := db.Where("id = ?", id).Delete(&product)
	fmt.Println(d)
	c.JSON(http.StatusOK, gin.H{"id #" + id: "deleted"})
}

func AppInstanceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		appInstance := os.Getenv("APP_INSTANCE")
		c.Header("x-app-instance", appInstance)
		c.Next()
	}
}
