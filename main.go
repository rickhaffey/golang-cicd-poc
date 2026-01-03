package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from simple API."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)

	port := "8080"
	log.Printf("starting server on :%s", port)

	err := http.ListenAndServe(":"+port, mux)
	log.Fatal(err)
}
