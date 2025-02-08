package main

import (
	"log"
	"net/http"
	"os"

	"github.com/nathansanchezdev/kube-event-logger/handlers"
	"github.com/nathansanchezdev/kube-event-logger/pkg/db"
)

func main() {
	// Initialize Database and handle errors properly
	if err := db.InitDB(); err != nil {
		log.Fatalf("❌ Database initialization failed: %v", err)
	}
	defer db.CloseDB()

	// Define API routes
	http.HandleFunc("/events", handlers.GetEvents)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server is running on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
