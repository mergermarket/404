package main

import "net/http"
import (
    "os"
)

func healthcheck_handler(res http.ResponseWriter, req *http.Request) {
    res.WriteHeader(http.StatusOK)
}

func default_handler(res http.ResponseWriter, req *http.Request) {
    res.WriteHeader(http.StatusNotFound)
}

func getListenAddress() string {
	port := getEnv("PORT", "80")
	return ":" + port
}

func main() {
    
    http.HandleFunc("/internal/healthcheck", healthcheck_handler)
    http.HandleFunc("/", default_handler)
    http.ListenAndServe(getListenAddress(), nil)
}
