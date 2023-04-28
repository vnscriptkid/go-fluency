package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"

	"github.com/streadway/amqp"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	deadQ    = "dead_letter_queue"
	deadQKey = "dead_letter_key"
	deadExc  = "dead_letter"
)

type NotificationType struct {
	ID       int    `gorm:"primaryKey"`
	TypeName string `gorm:"unique"`
}

type NotificationTemplate struct {
	ID                 int              `gorm:"primaryKey"`
	NotificationTypeID int              `gorm:"index"`
	NotificationType   NotificationType `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TemplateName       string           `gorm:"unique"`
	Title              string
	Content            string
}

func initDB() (*gorm.DB, error) {
	dsn := "user=user password=password dbname=meta_db host=localhost port=5433 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&NotificationType{}, &NotificationTemplate{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getQueueFromPriority(priority int) string {
	switch priority {
	case 1:
		return "priority_high"
	case 2:
		return "priority_medium"
	case 3:
		return "priority_low"
	default:
		return "priority_low"
	}
}

func getTemplate(db *gorm.DB, notificationType string, templateName string) (*NotificationTemplate, error) {
	var tmpl NotificationTemplate
	err := db.Joins("NotificationType").Where("type_name = ? AND template_name = ?", notificationType, templateName).Find(&tmpl).Error
	if err != nil {
		return nil, err
	}

	return &tmpl, nil
}

func main() {
	fmt.Println("Hello, world from control-svc")

	db, err := initDB()

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to DB: ", db)

	r := gin.Default()

	r.POST("/send-notification", func(c *gin.Context) {
		var jsonInput struct {
			// Can be phoneNumber, email, or pushToken
			Recipient    string                 `json:"recipient"`
			TemplateName string                 `json:"template_name"`
			Data         map[string]interface{} `json:"data"`
			Priority     int                    `json:"priority"`
			Channel      string                 `json:"channel"`
		}

		if err := c.BindJSON(&jsonInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		tmpl, err := getTemplate(db, jsonInput.Channel, jsonInput.TemplateName)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		t := template.Must(template.New("notification").
			Option("missingkey=error").
			Parse(tmpl.Content))

		renderedContent := &bytes.Buffer{}
		err = t.Execute(renderedContent, jsonInput.Data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		message := struct {
			Recipient string `json:"recipient"`
			Title     string `json:"title"`
			Content   string `json:"content"`
			Channel   string `json:"channel"`
		}{
			Recipient: jsonInput.Recipient,
			Title:     tmpl.Title,
			Content:   renderedContent.String(),
			Channel:   jsonInput.Channel,
		}

		messageJSON, err := json.Marshal(message)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		err = pushToRabbitMQ(getQueueFromPriority(jsonInput.Priority), messageJSON)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "success"})
	})

	seedData(db)

	r.Run(":8080")
}

func pushToRabbitMQ(queue string, message []byte) error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queue, // queue name
		true,
		false,
		false,
		false,
		amqp.Table{
			"x-dead-letter-exchange":    deadExc,
			"x-dead-letter-routing-key": deadQKey,
		},
	)

	if err != nil {
		return err
	}

	err = ch.Publish(
		// don't specify an exchange name
		// routes messages to the queue with a name that matches the routing key
		"",
		q.Name, // routing key
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	)
	if err != nil {
		fmt.Println("Failed to publish message", err)
		return err
	}

	return nil
}

func seedData(db *gorm.DB) error {
	// Seed NotificationType data
	notificationTypes := []NotificationType{
		{TypeName: "email"},
		{TypeName: "sms"},
		{TypeName: "push"},
	}

	for _, nType := range notificationTypes {
		err := db.FirstOrCreate(&nType, NotificationType{TypeName: nType.TypeName}).Error
		if err != nil {
			return err
		}
	}

	// Seed NotificationTemplate data
	notificationTemplates := []NotificationTemplate{
		{
			NotificationTypeID: 1,
			TemplateName:       "welcome_email",
			Title:              "Welcome to our platform!",
			Content: `Hi {{.UserName}},

Thank you for signing up for our service! We are excited to have you on board.

Please take a moment to familiarize yourself with our platform and let us know if you have any questions.

Best regards,
The {{.PlatformName}} Team`,
		},
		// Add more templates here for different notification types as needed
	}

	for _, tmpl := range notificationTemplates {
		err := db.FirstOrCreate(&tmpl, NotificationTemplate{TemplateName: tmpl.TemplateName}).Error
		if err != nil {
			return err
		}
	}

	return nil
}
