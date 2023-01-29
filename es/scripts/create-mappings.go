package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
)

type mapping struct {
	Mappings map[string]interface{} `json:"mappings"`
}

type EsConfig struct {
	client *elasticsearch.Client
}

func (c *EsConfig) createIdxFromFile(filePath string, idxName string) error {
	file, err := ioutil.ReadFile(filePath)

	if err != nil {
		return errors.New("error reading mapping file")
	}

	var m mapping
	err = json.Unmarshal(file, &m)

	if err != nil {
		return errors.New("error unmarshal file")
	}

	mappingJson, err := json.Marshal(m.Mappings)

	if err != nil {
		return errors.New("error marshalling mapping to json")
	}

	_, err = c.client.Indices.Create(
		idxName,
		c.client.Indices.Create.WithBody(strings.NewReader(string(mappingJson))),
	)

	if err != nil {
		return fmt.Errorf("error creating index: %s", err)
	}

	return nil
}

func (c *EsConfig) searchFromIdx(idxName string) (any, error) {
	searchResult, err := c.client.Search(
		c.client.Search.WithIndex(idxName),
		c.client.Search.WithBody(strings.NewReader(`{
			"query": {
				"match_all": {}
			}
		}`)),
		c.client.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		return nil, fmt.Errorf("error retrieving documents from index: %s", err)
	}

	return searchResult, nil
}

func main() {
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}
	es, err := elasticsearch.NewClient(cfg)

	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	esConfig := EsConfig{
		client: es,
	}

	err = esConfig.createIdxFromFile("./shakes-mapping.json", "shakespeare")

	if err != nil {
		fmt.Printf("err: %s", err)
		return
	}

	fmt.Println("Success")

	result, err := esConfig.searchFromIdx("shakespeare")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Result from shakespeare: %v", result)
}
