package main

import (
	"encoding/json"
	"log"
)

func toJSON(t Transaction) []byte {
	b, err := json.Marshal(t)
	if err != nil {
		log.Fatalf("Unable to marshal transaction to JSON: %v\n", err)
	}
	return b
}
