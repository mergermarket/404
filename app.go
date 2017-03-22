package main

import "net/http"

func healthcheck_handler(res http.ResponseWriter, req *http.Request) {
    res.WriteHeader(http.StatusOK)
}

func default_handler(res http.ResponseWriter, req *http.Request) {
    res.WriteHeader(http.StatusNotFound)
}

func main() {
    http.HandleFunc("/internal/healthcheck", healthcheck_handler)
    http.HandleFunc("/", default_handler)
    http.ListenAndServe(":80", nil)
}
