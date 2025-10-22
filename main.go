package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Set up routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/health", healthHandler)

	// Start server
	port := "8080"
	fmt.Printf("Starting Zoho Email Scheduling App on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Zoho Email Scheduling App!\n")
	fmt.Fprintf(w, "This is a basic Go application.\n")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"status":"healthy"}`)
}
