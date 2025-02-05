package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

// Event represents a Kubernetes event log
type Event struct {
	ID          int64     `json:"id"`
	Type        string    `json:"type"`
	Message     string    `json:"message"`
	Timestamp   time.Time `json:"timestamp"`
	Severity    string    `json:"severity"`
	Source      string    `json:"source"`
	ClusterName string    `json:"clusterName"`
	Namespace   string    `json:"namespace"`
}

// Database connection
var db *sql.DB

func init() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:postgres@localhost:5432/eventdb?sslmode=disable"
	}

	// Initialize database connection
	var err error
	db, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		log.Fatal("Error pinging database: ", err)
	}

	log.Println("Successfully connected to database")
}

func getEvents(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`
		SELECT id, type, message, timestamp, severity, source, cluster_name, namespace 
		FROM events 
		ORDER BY timestamp DESC 
		LIMIT 100
	`)
	if err != nil {
		log.Printf("Error querying events: %v", err)
		http.Error(w, "Error retrieving events", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var e Event
		err := rows.Scan(
			&e.ID,
			&e.Type,
			&e.Message,
			&e.Timestamp,
			&e.Severity,
			&e.Source,
			&e.ClusterName,
			&e.Namespace,
		)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}
		events = append(events, e)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

func main() {
	http.HandleFunc("/events", getEvents)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server is running on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
