package main

import (
	"log"
)

func home(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello from simple API."))
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)

	port := "8080"
	log.Printf("starting server on :%s", port)

	err := http.ListenAndServe(":"+port, mux)
	log.Fatal(err)
}
