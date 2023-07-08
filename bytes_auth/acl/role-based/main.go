package main

import (
	"fmt"
)

type RolePermissions map[string]ResourcePermissions
type ResourcePermissions map[string]CrudPermissions
type CrudPermissions map[string]bool

var rolePermissions = RolePermissions{
	"Admin": ResourcePermissions{
		"Post": CrudPermissions{
			"Read":   true,
			"Write":  true,
			"Update": true,
			"Delete": true,
		},
		"Comment": CrudPermissions{
			"Read":   true,
			"Write":  true,
			"Update": true,
			"Delete": true,
		},
	},
	"Editor": ResourcePermissions{
		"Post": CrudPermissions{
			"Read":   true,
			"Write":  true,
			"Update": true,
			"Delete": false,
		},
		"Comment": CrudPermissions{
			"Read":   true,
			"Write":  false,
			"Update": false,
			"Delete": false,
		},
	},
	"User": ResourcePermissions{
		"Post": CrudPermissions{
			"Read":   true,
			"Write":  false,
			"Update": false,
			"Delete": false,
		},
		"Comment": CrudPermissions{
			"Read":   true,
			"Write":  true,
			"Update": false,
			"Delete": false,
		},
	},
}

type User struct {
	Name  string
	Roles []string
}

func CanPerformAction(user User, resource string, action string) bool {
	for _, role := range user.Roles {
		resourcePermissions, ok := rolePermissions[role][resource]
		if ok && resourcePermissions[action] {
			return true
		}
	}

	return false
}

func main() {
	// 	Admin can Read, Write, Update and Delete both Post and Comment.
	// Editor can Read, Write, and Update a Post but can only Read a Comment.
	// User can Read a Post and a Comment but can Write a Comment.
	alice := User{Name: "Alice", Roles: []string{"Admin"}}
	bob := User{Name: "Bob", Roles: []string{"Editor", "User"}}
	charlie := User{Name: "Charlie", Roles: []string{"User"}}

	fmt.Println(CanPerformAction(alice, "Post", "Delete"))     // Output: true
	fmt.Println(CanPerformAction(bob, "Post", "Delete"))       // Output: false
	fmt.Println(CanPerformAction(charlie, "Comment", "Write")) // Output: true
	fmt.Println(CanPerformAction(charlie, "Post", "Write"))    // Output: false
	fmt.Println(CanPerformAction(bob, "Comment", "Write"))     // Output: true
}
