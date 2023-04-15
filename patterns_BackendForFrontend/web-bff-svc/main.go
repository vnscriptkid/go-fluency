package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type UserProfile struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	LastPost    Post   `json:"last_post"`
	LastComment string `json:"last_comment"`
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Post struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func fetchUser() (User, error) {
	resp, err := http.Get("http://user-service")
	if err != nil {
		return User{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return User{}, err
	}

	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func fetchPost() (Post, error) {
	resp, err := http.Get("http://post-service")
	if err != nil {
		return Post{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Post{}, err
	}

	var post Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		return Post{}, err
	}

	return post, nil
}

func main() {
	http.HandleFunc("/user-profile", func(w http.ResponseWriter, r *http.Request) {
		user, err := fetchUser()
		if err != nil {
			http.Error(w, "Error fetching user data", http.StatusInternalServerError)
			return
		}

		post, err := fetchPost()
		if err != nil {
			http.Error(w, "Error fetching post data", http.StatusInternalServerError)
			return
		}

		profile := UserProfile{
			Name:        user.Name,
			Email:       user.Email,
			LastPost:    Post{Title: post.Title},
			LastComment: "This is the latest comment.",
		}

		jsonData, err := json.Marshal(profile)

		if err != nil {
			http.Error(w, "Error creating user profile JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
