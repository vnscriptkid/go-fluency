package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Greetings from the greeting service!")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
