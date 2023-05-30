package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/gin-gonic/gin"
)

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

var es *elasticsearch.Client

func main() {

	// Initialize Elasticsearch client
	_es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	})

	if err != nil {
		log.Fatal(err)
	}

	es = _es

	// Create the index and bulk index the user data
	// Sample user data
	users := []User{
		{ID: 1, Username: "ronaldo", Name: "Cristiano Ronaldo", Email: "ronaldo@example.com"},
		{ID: 2, Username: "messi", Name: "Lionel Messi", Email: "messi@example.com"},
		{ID: 3, Username: "neymar", Name: "Neymar Jr.", Email: "neymar@example.com"},
		{ID: 4, Username: "mbappe", Name: "Kylian Mbappe", Email: "mbappe@example.com"},
		{ID: 5, Username: "salah", Name: "Mohamed Salah", Email: "salah@example.com"},
		{ID: 6, Username: "modric", Name: "Luka Modric", Email: "modric@example.com"},
		{ID: 7, Username: "kane", Name: "Harry Kane", Email: "kane@example.com"},
		{ID: 8, Username: "lewandowski", Name: "Robert Lewandowski", Email: "lewandowski@example.com"},
		{ID: 9, Username: "debruyne", Name: "Kevin De Bruyne", Email: "debruyne@example.com"},
		{ID: 10, Username: "pogba", Name: "Paul Pogba", Email: "pogba@example.com"},

		// Similar usernames and names for testing edge n-gram
		{ID: 11, Username: "ron", Name: "Ronaldinho", Email: "ron@example.com"},
		{ID: 12, Username: "mes", Name: "Mesut Ozil", Email: "mes@example.com"},
		{ID: 13, Username: "ney", Name: "Neymar Santos", Email: "ney@example.com"},
		{ID: 14, Username: "mba", Name: "Mbappe Lottin", Email: "mba@example.com"},
		{ID: 15, Username: "sal", Name: "Salah Mane", Email: "sal@example.com"},
		{ID: 16, Username: "mod", Name: "Modric Luka", Email: "mod@example.com"},
		{ID: 17, Username: "kan", Name: "Kane Harry", Email: "kan@example.com"},
		{ID: 18, Username: "lew", Name: "Lewandowski Robert", Email: "lew@example.com"},
		{ID: 19, Username: "deb", Name: "De Bruyne Kevin", Email: "deb@example.com"},
		{ID: 20, Username: "pog", Name: "Pogba Paul", Email: "pog@example.com"},
	}

	// Create the index and bulk index the user data
	skipSeeder := false
	if err := createUserMapping(users); err != nil {
		// log.Fatal(err)
		if strings.Contains(err.Error(), "resource_already_exists_exception") {
			log.Println("Index already exists")
			skipSeeder = true
		} else {
			log.Fatal(err)
		}
	}

	if !skipSeeder {
		if err := bulkIndexUser(users); err != nil {
			log.Fatal(err)
		}
	}

	// Serve API
	r := gin.Default()

	r.GET("/search", searchUsers)

	// Define a route for the HTML file
	r.GET("/", func(c *gin.Context) {
		c.File("./public/index.html")
	})

	if err := r.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}

func searchUsers(c *gin.Context) {
	query := c.Query("query")

	// Retrieve the user from Elasticsearch
	users, err := searchUserData(query)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

// Index user data into Elasticsearch with edge n-gram
func createUserMapping(users []User) error {
	// Define the mapping for the index
	mapping := `{
		"mappings": {
			"properties": {
				"username": {
					"type": "text",
					"analyzer": "autocomplete",
					"search_analyzer": "standard"
				},
				"name": {
					"type": "text",
					"analyzer": "autocomplete",
					"search_analyzer": "standard"
				}
			}
		},
		"settings": {
			"analysis": {
				"filter": {
					"autocomplete_filter": {
						"type": "edge_ngram",
						"min_gram": 1,
						"max_gram": 20
					}
				},
				"analyzer": {
					"autocomplete": {
						"type": "custom",
						"tokenizer": "standard",
						"filter": [
							"lowercase",
							"autocomplete_filter"
						]
					}
				}
			}
		}
	}`

	// Create the index with the mapping
	createIndexReq := esapi.IndicesCreateRequest{
		Index: "users",
		Body:  strings.NewReader(mapping),
	}

	createIndexRes, err := createIndexReq.Do(context.Background(), es)

	if err != nil {
		return fmt.Errorf("error creating index: %s", err)
	}

	defer createIndexRes.Body.Close()

	if createIndexRes.IsError() {
		return fmt.Errorf("index creation failed: %s", createIndexRes.String())
	}

	return nil
}

func bulkIndexUser(users []User) error {
	// Bulk index request body
	var buf strings.Builder

	for _, user := range users {
		indexReq := fmt.Sprintf(`{ "index" : { "_index" : "users" } }%s`, "\n")
		data, err := json.Marshal(user)
		if err != nil {
			return fmt.Errorf("error encoding user data: %s", err)
		}
		buf.WriteString(indexReq)
		buf.Write(data)
		buf.WriteString("\n")
	}

	// Set up the bulk index request
	bulkReq := esapi.BulkRequest{
		Body: strings.NewReader(buf.String()),
	}

	// Perform the bulk index request
	bulkRes, err := bulkReq.Do(context.Background(), es)
	if err != nil {
		return fmt.Errorf("error performing bulk index request: %s", err)
	}
	defer bulkRes.Body.Close()

	// Check the response status
	if bulkRes.IsError() {
		return fmt.Errorf("bulk index request failed: %s", bulkRes.String())
	}

	// Parse the response
	var result map[string]interface{}
	if err := json.NewDecoder(bulkRes.Body).Decode(&result); err != nil {
		return fmt.Errorf("error parsing bulk index response: %s", err)
	}

	// Check for any errors in the response
	if result["errors"] != nil && result["errors"].(bool) {
		return fmt.Errorf("bulk index request contains errors")
	}

	return nil
}

// Search user data in Elasticsearch
func searchUserData(query string) ([]User, error) {
	// Define the search request
	searchReq := esapi.SearchRequest{
		Index: []string{"users"},
		Body: strings.NewReader(fmt.Sprintf(`{
			"query": {
				"multi_match": {
					"query": "%s",
					"type": "cross_fields",
					"fields": ["username", "name"],
					"operator": "or"
				}
			}
		}`, query)),
	}

	// Perform the search request
	searchRes, err := searchReq.Do(context.Background(), es)
	if err != nil {
		return nil, fmt.Errorf("error performing search request: %s", err)
	}
	defer searchRes.Body.Close()

	// Check the response status
	if searchRes.IsError() {
		return nil, fmt.Errorf("search request failed: %s", searchRes.String())
	}

	// Parse the response
	var response map[string]interface{}
	if err := json.NewDecoder(searchRes.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("error parsing search response: %s", err)
	}

	// Extract the search results
	hits := response["hits"].(map[string]interface{})["hits"].([]interface{})

	// Prepare the user data
	var users []User
	for _, hit := range hits {
		source := hit.(map[string]interface{})["_source"]
		userData, err := json.Marshal(source)
		if err != nil {
			return nil, fmt.Errorf("error encoding user data: %s", err)
		}
		var user User
		if err := json.Unmarshal(userData, &user); err != nil {
			return nil, fmt.Errorf("error decoding user data: %s", err)
		}
		users = append(users, user)
	}

	return users, nil
}
