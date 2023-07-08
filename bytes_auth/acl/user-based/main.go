package main

import (
	"fmt"
)

type User struct {
	ID   string
	Name string
}

type Resource struct {
	ID   string
	Name string
}

type Permission struct {
	Read  bool
	Write bool
	Exec  bool
}

type ACL struct {
	User       User
	Resource   Resource
	Permission Permission
}

func main() {
	// Define users
	user1 := User{ID: "1", Name: "Alice"}
	user2 := User{ID: "2", Name: "Bob"}

	// Define resources
	res1 := Resource{ID: "1", Name: "Document1"}
	res2 := Resource{ID: "2", Name: "Document2"}

	// Define ACL
	acl1 := ACL{
		User:       user1,
		Resource:   res1,
		Permission: Permission{Read: true, Write: true, Exec: false},
	}
	acl2 := ACL{
		User:       user1,
		Resource:   res2,
		Permission: Permission{Read: false, Write: false, Exec: true},
	}
	acl3 := ACL{
		User:       user2,
		Resource:   res1,
		Permission: Permission{Read: true, Write: false, Exec: false},
	}

	// Add ACLs to a list
	aclList := []ACL{acl1, acl2, acl3}

	// Check permission
	checkUserPermission(aclList, "1", "1")

	// Check if user can perform an action
	fmt.Println(canPerformAction(aclList, "1", "1", "read"))
	fmt.Println(canPerformAction(aclList, "1", "1", "exec"))
}

func checkUserPermission(aclList []ACL, userID string, resourceID string) {
	for _, acl := range aclList {
		if acl.User.ID == userID && acl.Resource.ID == resourceID {
			fmt.Printf("User %s has the following permissions on Resource %s: %+v\n", acl.User.Name, acl.Resource.Name, acl.Permission)
		}
	}
}

func canPerformAction(aclList []ACL, userID string, resourceID string, action string) bool {
	for _, acl := range aclList {
		if acl.User.ID == userID && acl.Resource.ID == resourceID {
			switch action {
			case "read":
				return acl.Permission.Read
			case "write":
				return acl.Permission.Write
			case "exec":
				return acl.Permission.Exec
			}
		}
	}
	return false
}
