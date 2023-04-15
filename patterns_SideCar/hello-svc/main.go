package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func rootHanlder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func crossHanlder(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://greeting-service:8081")

	if err != nil {
		http.Error(w, "Error calling greeting service", http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading greeting service response", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Hello, World! %s", string(body))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHanlder)
	mux.HandleFunc("/cross", crossHanlder)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
