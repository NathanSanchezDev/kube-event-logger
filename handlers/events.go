package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nathansanchezdev/kube-event-logger/models"
	"github.com/nathansanchezdev/kube-event-logger/pkg/db"
)

// GetEvents retrieves events from the database
func GetEvents(w http.ResponseWriter, r *http.Request) {
	if db.DB == nil {
		log.Println("Database connection is nil")
		http.Error(w, "Database connection is not initialized", http.StatusInternalServerError)
		return
	}

	rows, err := db.DB.Query(`
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

	var events []models.Event
	for rows.Next() {
		var e models.Event
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

	// Send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}
