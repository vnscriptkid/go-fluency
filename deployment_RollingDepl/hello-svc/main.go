package main

import (
	"fmt"
	"net/http"
	"os"
)

var version = "v3.0.0" // application version

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/version", getVersion)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Listening on port %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func getVersion(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "App version:", version)
}
