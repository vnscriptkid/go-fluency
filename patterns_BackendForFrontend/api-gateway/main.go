package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func routeRequest(targetURL string, w http.ResponseWriter, r *http.Request) {
	target, _ := url.Parse(targetURL)

	proxy := httputil.NewSingleHostReverseProxy(target)

	proxy.ServeHTTP(w, r)
}

func main() {
	http.HandleFunc("/user-profile", func(w http.ResponseWriter, r *http.Request) {
		clientType := r.Header.Get("X-Client-Type")

		switch clientType {
		case "web":
			routeRequest("http://web-bff-service", w, r)
		case "mobile":
			routeRequest("http://mobile-bff-service", w, r)
		default:
			http.Error(w, "Invalid client type", http.StatusBadRequest)
		}
	})

	log.Fatal(http.ListenAndServe(":8085", nil))
}
