package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Event represents a Kubernetes event log
type Event struct {
	Type      string `json:"type"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

// Predefined mock events
var events = []Event{
	{Type: "INFO", Message: "Pod kube-api restarted", Timestamp: time.Now().Format(time.RFC3339)},
	{Type: "WARNING", Message: "Node memory pressure detected", Timestamp: time.Now().Format(time.RFC3339)},
	{Type: "ERROR", Message: "Failed to pull image: registry.example.com/nginx", Timestamp: time.Now().Format(time.RFC3339)},
}

// Handler for GET /events
func getEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

func main() {
	http.HandleFunc("/events", getEvents)

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
