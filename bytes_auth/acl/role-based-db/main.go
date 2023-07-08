package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Roles    []Role `gorm:"many2many:user_roles;"`
}

type Role struct {
	gorm.Model
	Name        string
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}

type Resource struct {
	gorm.Model
	Name string
}

type Action struct {
	gorm.Model
	Name string
}

type Permission struct {
	gorm.Model
	ActionID   uint
	ResourceID uint
	Action     Action
	Resource   Resource
}

func CanPerformAction(db *gorm.DB, user User, resource string, action string) bool {
	// fetch the user,
	// - the roles associated with that user,
	// - the permissions associated with each of those role
	db.Preload("Roles.Permissions.Action").Preload("Roles.Permissions.Resource").Find(&user, user.ID)

	for _, role := range user.Roles {
		for _, permission := range role.Permissions {
			if permission.Resource.Name == resource && permission.Action.Name == action {
				return true
			}
		}
	}

	return false
}

func main() {
	r := gin.Default()

	var (
		host     = "localhost"
		port     = 5435
		user     = "user"
		password = "password"
		dbname   = "rbac"
		sslmode  = "disable"
	)

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v", host, user, password, dbname, port, sslmode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&User{}, &Role{}, &Resource{}, &Action{}, &Permission{})

	if err != nil {
		panic("failed to migrate database")
	}

	r.GET("/resource", func(c *gin.Context) {
		username := c.Query("username") // retrieve the username from the query parameters
		var user User
		db.Where("username = ?", username).First(&user)

		fmt.Println(user)

		if CanPerformAction(db, user, "Resource", "Read") {
			c.JSON(200, gin.H{
				"message": "Access granted",
			})
		} else {
			c.JSON(403, gin.H{
				"message": "Access denied",
			})
		}
	})

	r.POST("/assign-role", func(c *gin.Context) {
		// Define a new role
		editor := Role{Name: "Editor"}
		db.Save(&editor)

		// Define new resources
		resource := Resource{Name: "Resource"}
		db.Save(&resource)

		// Define new actions
		actionRead := Action{Name: "Read"}
		db.Save(&actionRead)

		// Define a new permission
		readResourcePermission := Permission{ResourceID: resource.ID, ActionID: actionRead.ID}
		db.Save(&readResourcePermission)

		// Assign the permission to the role
		editor.Permissions = append(editor.Permissions, readResourcePermission)
		db.Save(&editor)

		// Create a user and assign them the role
		user := User{Username: "john"}
		user.Roles = append(user.Roles, editor)
		db.Save(&user)

		c.JSON(200, gin.H{
			"message": "Role assigned",
		})
	})

	r.Run()
}
