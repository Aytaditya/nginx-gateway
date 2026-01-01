package main

import (
	"fmt"
	"net/http"
)

func main() {
	// here we will load config
	// here we will connect to database
	// here we will start http server
	router := http.NewServeMux()
	router.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to the Test Service"))
	})
	router.HandleFunc("GET /test/config", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Config Loaded Successfully"))
	})
	fmt.Println("Starting Test Service on :8081")
	err := http.ListenAndServe(":8081", router)
	if err != nil {
		panic(err)
	}
}
